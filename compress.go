package compress

// ICompress is an interface that requires a type to return a pointer to Compress[T].
type ICompress[T any] interface {
	Compress() *Compress[T]
}

// Compress provides functional-style operations (map, filter, etc.) on a generic slice.
type Compress[T any] struct {
	data []T
}

// New creates a new Compress instance from a slice of T.
func New[T any](data []T) *Compress[T] {
	if len(data) == 0 {
		data = make([]T, 0)
	}
	return &Compress[T]{data}
}

// filterFunc defines a function that returns true if an element should be kept.
type filterFunc[T any] = func(elem T) bool

// mapFunc defines a function that transforms an element.
type mapFunc[T any] = func(elem T) T

// Filter keeps only the elements for which the provided function returns true.
// If the receiver or its data is nil, it returns nil.
func (c *Compress[T]) Filter(f filterFunc[T]) *Compress[T] {
	if len(c.data) == 0 {
		return c
	}
	filteredData := make([]T, 0, len(c.data))
	for _, elem := range c.data {
		if f(elem) {
			filteredData = append(filteredData, elem)
		}
	}
	c.data = filteredData
	return c
}

// Map applies the provided function to each element in the slice, modifying it in place.
// If the receiver or its data is nil, it returns nil.
func (c *Compress[T]) Map(f mapFunc[T]) *Compress[T] {
	if len(c.data) == 0 {
		return c
	}
	for i, elem := range c.data {
		c.data[i] = f(elem)
	}
	return c
}

// At returns the element at the specified index.
// If the index is out of bounds, it clamps to [0, len-1].
// If the slice is nil or empty, returns the zero value of T.
func (c *Compress[T]) At(index int) T {
	var value T
	if len(c.data) == 0 {
		return value
	}
	if index >= len(c.data) {
		index = len(c.data) - 1
	}
	if index < 0 {
		index = 0
	}
	return c.data[index]
}

// Head returns the first element of the slice.
// It returns the zero value of T if the slice is nil or empty.
func (c *Compress[T]) Head() T {
	var value T
	if len(c.data) == 0 {
		return value
	}
	return c.data[0]
}

func (c *Compress[T]) Tail() T {
	var value T
	if len(c.data) == 0 {
		return value
	}
	return c.data[len(c.data)-1]
}

// Pop removes and returns the last element of the slice.
// If the slice is nil or empty, it returns the zero value of T.
func (c *Compress[T]) Pop() T {
	var value T
	if len(c.data) == 0 {
		return value
	}
	lastIndex := len(c.data) - 1
	value = c.data[lastIndex]
	c.data = c.data[:lastIndex]
	return value
}

// Shift removes the first element from the slices and returns it.
// If the slice is nil or empty, it returns the zero value of T.
func (c *Compress[T]) Shift() T {
	var value T
	if c == nil || len(c.data) == 0 {
		return value
	}
	value = c.data[0]
	c.data = c.data[1:]
	return value
}

// Range returns a new Compress with elements from index start to end (exclusive).
// If indices are out of bounds, they are clamped to valid ranges.
// If start >= end, an empty slice is returned.
func (c *Compress[T]) Range(start, end int) *Compress[T] {
	if len(c.data) == 0 {
		return c
	}
	if start < 0 {
		start = 0
	}
	if end > len(c.data) {
		end = len(c.data)
	}
	if start > end {
		start = end
	}
	c.data = c.data[start:end]
	return c
}

type everyFunc[T any] = func(elem T) bool

// Every checks if all elements in the slice satisfy the given predicate function.
// It returns false if the slice is nil or empty.
func (c *Compress[T]) Every(f everyFunc[T]) bool {
	if len(c.data) == 0 {
		return false
	}
	for _, elem := range c.data {
		if !f(elem) {
			return false
		}
	}
	return true
}

// Entries returns a slice of [index, value] pairs from the internal data slice.
// Each pair is represented as [2]any, where the first is the index (int) and second is the value (T).
func (c *Compress[T]) Entries() [][2]any {
	result := make([][2]any, len(c.data))
	if len(c.data) == 0 {
		return result
	}
	for i, elem := range c.data {
		data := [2]any{i, elem}
		result[i] = data
	}
	return result
}

type findFunc[T any] = func(elem T) bool

// Find returns the first element in the slice that satisfies the predicate function.
// If no element matches or the slice is nil/empty, it returns the zero value of T.
func (c *Compress[T]) Find(f findFunc[T]) T {
	var value T
	if len(c.data) == 0 {
		return value
	}
	for _, e := range c.data {
		if f(e) {
			return e
		}
	}
	return value
}

// Returns the slice modified
func (c *Compress[T]) Result() []T {
	return c.data
}
