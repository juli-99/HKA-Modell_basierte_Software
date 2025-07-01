package queue

/* Using generics for the queue makes sense
 * if we want it to be type-safe and consistent.
 * If we needed a queue that could store different types of values,
 * we could have used interface{} instead.
 * However, using generics allows the compiler to enforce that all elements
 * in the queue are of the same type,
 * avoiding the need for type assertions and reducing the risk of runtime errors.
 */

// generic queue structure
type Queue[T any] struct {
	items []T
}

// create a new queue
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// add item to the top of queue
func (q *Queue[T]) Add(item T) {
	q.items = append(q.items, item)
}

// remove and return from top of the stack
func (q *Queue[T]) Next() (T, bool) {
	if len(q.items) == 0 {
		var default_val T
		return default_val, false // return default value and false if stack is empty
	}
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return item, true
}

// return from top of the queue
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false // return default value and false if queue is empty
	}
	return q.items[0], true
}

// checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
