# 📦 compress

`compress` is a generic utility package for working with slices in Go using a fluent, chainable API. Inspired by functional programming concepts like `map`, `filter`, `reduce`, `find`, and more.

---

## ✨ Features

- ✅ `Filter`
- ✅ `Map`
- ✅ `Reduce`
- ✅ `Find`
- ✅ `Some`
- ✅ `Every`
- ✅ `Head`, `Tail`, `Pop`, `Shift`
- ✅ `Entries`, `At`, `Slice`, `Result`

---

## 📦 Installation

```bash
go get github.com/ronanzindev/compress
```

## Example: Todo List with compress

```golang
package main

import (
	"fmt"

	"github.com/ronanzindev/compress"
)

type Task struct {
	Title     string
	Completed bool
}

type TodoList struct {
	tasks []Task
}

// Implements ICompress interface
func (t *TodoList) Compress() *compress.Compress[Task] {
	return compress.New(t.tasks)
}

func main() {
	todo := &TodoList{
		tasks: []Task{
			{"Buy milk", false},
			{"Clean room", true},
			{"Go to gym", false},
			{"Read book", true},
		},
	}

	// Filter completed tasks
	completedTasks := todo.Compress().
		Filter(func(task Task) bool {
			return task.Completed
		}).Result()

	fmt.Println("✅ Completed Tasks:")
	for _, task := range completedTasks {
		fmt.Println("-", task.Title)
	}

	// Mark all tasks as completed
	allDone := todo.Compress().
		Map(func(task Task) Task {
			task.Completed = true
			return task
		}).Result()

	fmt.Println("\n📌 All Tasks Marked as Done:")
	for _, task := range allDone {
		fmt.Printf("- %s: %v\n", task.Title, task.Completed)
	}

	// Check if all tasks are completed
	areAllDone := todo.Compress().
		Every(func(t Task) bool { return t.Completed })

	fmt.Printf("\n🧪 Are all tasks complete? %v\n", areAllDone)
}
```

## Example: Chaining Operations

```golang
incompleteTitles := todo.Compress().Filter(func(task Task) bool {
		return !task.Completed
	}).Map(func(task Task) Task {
		task.Title = "[TODO] " + task.Title
		return task
	}).Map(func(task Task) Task {
		fmt.Println("Processing:", task.Title)
		return task
	}).Result()

    fmt.Println("\n📝 Incomplete Tasks:")
    for _, task := range incompleteTitles {
	    fmt.Println("-", task.Title)
    }
```