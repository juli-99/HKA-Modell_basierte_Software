package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/juli-99/hka-modell_basierte_software/stack"
)

const NUM_WORKERS int = 3
const NUM_INTS int = 20

/* Using generics instead of interfaces is necessary in this case
 * because the validation function can operate on arbitrary types,
 * and we don't know in advance what operations it will perform.
 * By using generics, we ensure that the item and the validation function
 * both use the same concrete type, enabling full compile-time type safety
 * without requiring a common interface. This flexibility would not be possible
 * with interfaces alone, since interfaces require predefined method sets.
 */
func worker[T any](
	id int,
	in <-chan T,
	out chan<- bool,
	validate func(T) bool,
) {
	for item := range in {
		valid := validate(item)
		fmt.Printf("worker %d: item: %v result: %t\n", id, item, valid)
		out <- valid
	}
	fmt.Printf("worker %d: Finished!\n", id)
}

func main() {
	// Create a stack for integers
	stack_int := stack.New[int]()
	for i := 0; i <= NUM_INTS; i++ {
		stack_int.Push(5 + i*7)
	}

	// Validation function: even numbers are valid
	validate_int := func(n int) bool {
		return n%2 == 0
	}

	ch_in_int := make(chan int)
	ch_out_int := make(chan bool)
	var num_valid_int int

	// Start workers
	for i := 1; i <= NUM_WORKERS; i++ {
		go worker(i, ch_in_int, ch_out_int, validate_int)
	}

	for !stack_int.IsEmpty() {
		item, _ := stack_int.Pop()
		ch_in_int <- item
		if <-ch_out_int {
			num_valid_int++
		}
	}
	close(ch_in_int)
	fmt.Printf("Number of valid items: %d\n", num_valid_int)

	time.Sleep(2 * time.Second) // sleep for 10 seconds
	/////////////

	// Create a stack for strings
	stack_str := stack.New[string]()
	stack_str.Push("Hello World")
	stack_str.Push("Generics")
	stack_str.Push("World Wide Web")

	// Validation function: string contains World
	validate_str := func(s string) bool {
		return strings.Contains(s, "World")
	}

	ch_in_str := make(chan string)
	ch_out_str := make(chan bool)
	var num_valid_str int

	// Start workers
	for i := 1; i <= NUM_WORKERS; i++ {
		go worker(i+10, ch_in_str, ch_out_str, validate_str)
	}

	for !stack_str.IsEmpty() {
		item, _ := stack_str.Pop()
		ch_in_str <- item
		if <-ch_out_str {
			num_valid_str++
		}
	}
	close(ch_in_str)
	fmt.Printf("Number of valid items: %d\n", num_valid_str)
}
