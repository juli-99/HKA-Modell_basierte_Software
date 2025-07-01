package stack

/* Using generics for the stack makes sense
 * if we want it to be type-safe and consistent.
 * If we needed a stack that could store different types of values,
 * we could have used interface{} instead.
 * However, using generics allows the compiler to enforce that all elements
 * in the stack are of the same type,
 * avoiding the need for type assertions and reducing the risk of runtime errors.
 */

// generic stack structure
type Stack[T any] struct {
	items []T
}

// create a new Stack
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// add item to the top of stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// remove and return from top of the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var default_val T
		return default_val, false // return default value and false if stack is empty
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// return from top of the stack
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false // return default value and false if stack is empty
	}
	return s.items[len(s.items)-1], true
}

// checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
