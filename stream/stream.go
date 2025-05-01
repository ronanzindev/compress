package stream

import "sync"

type Stream[T any] struct {
	data <-chan T
}

func NewStream[T any](data []T) *Stream[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, item := range data {
			ch <- item
		}
	}()
	return &Stream[T]{data: ch}
}

func (s *Stream[T]) Filter(predicate func(T) bool) *Stream[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for item := range s.data {
			if predicate(item) {
				ch <- item
			}
		}
	}()
	return &Stream[T]{data: ch}
}

func (s *Stream[T]) Map(predicate func(T) T) *Stream[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for item := range s.data {
			result := predicate(item)
			ch <- result
		}
	}()
	return &Stream[T]{data: ch}
}

func (s *Stream[T]) FlatMap(transform func(T) []T) *Stream[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for item := range s.data {
			for _, transformed := range transform(item) {
				ch <- transformed
			}
		}
	}()
	return &Stream[T]{data: ch}
}

func (s *Stream[T]) Reduce(inital T, reduce func(T, T) T) T {
	result := inital

	for item := range s.data {
		result = reduce(inital, item)
	}
	return result
}

func (s *Stream[T]) Limit(n int) *Stream[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		count := 0
		for item := range s.data {
			if count >= n {
				break
			}
			ch <- item
			count++
		}
	}()
	return &Stream[T]{data: ch}
}

func (s *Stream[T]) Parallel(workers int) *Stream[T] {
	ch := make(chan T)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for item := range s.data {
				ch <- item
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return &Stream[T]{data: ch}
}

func (s *Stream[T]) Collect() []T {
	result := make([]T, 0)
	for item := range s.data {
		result = append(result, item)
	}
	return result
}
