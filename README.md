# Generics in Go

Generics in Go allow developers to write functions and data structures that can operate on any type without sacrificing type safety.

A common motivating example is a [**stack**](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)).

## How it used to be done

Before Go 1.18, developers either implemented the same functionality for multiple data types or used `interface{}` for generic behavior.

### Type specific

The simplest way to create a stack that can use different types is to implement it for each type.

```go
// int stack structure
type IntStack struct {
    items []int
}

// create a new IntStack
func NewIntStack() *IntStack {
    return &IntStack{}
}

// add item to the top of stack
func (s *IntStack) Push(item int) {
    s.items = append(s.items, item)
}

// string stack structure
type StringStack struct {
    items []string
}

// create a new StringStack
func NewStringStack() *StringStack {
    return &StringStack{}
}

// add item to the top of stack
func (s *StringStack) Push(item string) {
    s.items = append(s.items, item)
}
```

The problem is that it does not work for any type but only for the implemented ones.
And it creates a lot of code duplication.

### `interface{}`

Using `interface{}` the code duplication can be minimized and it accepts any possible type.

```go
// stack structure
type Stack struct {
    items []interface{}
}

// create a new Stack
func NewStack() *Stack {
    return &Stack{}
}

// add item to the top of stack
func (s *Stack) Push(item interface{}) {
    s.items = append(s.items, item)
}

// add elements to the stack
stack := NewStack()
stack.Push(123)
stack.Push("asdf")
stack.Push(NewStack())
```

A problem is that the stack allows all types of data and is not type consistant.
That can lead to the following problems:

- **Type assertions** are needed when popping values (`item.(int)`).
- There's **no compile-time check**, a `string` might get into a stack of `int`.
- It's easy to introduce runtime bugs.

## Syntax and How to Use

With generics, code becomes:

- **More reusable** (single definition works for any type),
- **More readable** (type is preserved),
- **Safer** (type-checked at compile time).

Go’s generics syntax introduces **type parameters** using square brackets `[]`.
You define type parameters alongside types and functions to make them generic.

### Structs and functions

#### Example: Generic Stack

You can define a struct that works for any type `T` by parameterizing it.

```go
type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{}
}
```

This defines a stack where all elements are of type `T`. You can then use it like:

```go
intStack := NewStack[int]()
strStack := NewStack[string]()
```

Generic functions allow you to define reusable logic for any type.

```go
// add item to the top of stack
func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}
```

The function works for any `T`, preserving its type, so that no type assertions are required.

#### Example: 2-Tuple

Generics allow the user to use multiple types in structs or functions.

```go
type Tuple[A, B any] struct {
    First  A
    Second B
}

func NewTuple[A, B any](a A, b B) Tuple[A, B] {
    return Tuple[A, B]{First: a, Second: b}
}
```

To simplify the code, Go is able to infer the type arguments from the types of the function arguments.

```go
pair1 := NewTuple[string, int]("ModSoft", 2025) // explicity declared
pair2 := NewTuple("ModSoft", 2025)              // infered from arguments
```

### With Type Requirements

You can constrain type parameters using interfaces (also known as "type sets").
This limits what operations are valid for a generic type.

```go
type Ordered interface {
    ~int | ~float64 | ~string
}

func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

Now, `Max` works only for types that support the `>` operator.

## When to Use

Generics should be used when they improve code clarity, reduce duplication, and ensure type safety. The most effective use cases typically follow one simple guideline:

> If you find yourself **writing the same code multiple times** for different types, consider using a type parameter.

### Functions working with built-in containers (slices, maps, channels)

When writing functions that operate on slices, maps, or channels without making assumptions about the types of the elements, generics can make the code reusable and type-safe.

For example, instead of duplicating logic for `[]int` and `[]string`, following code can be writen:

```go
func MapKeys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
```

This function works for any map type, with compile-time checks and no reflection.

### General-purpose data structure

A `Stack[T]` is the perfect example for a generic structure:

- It does not care about methods on `T`.
- All operations (push, pop, peek) are type-agnostic.
- Using `Stack[int]` vs. `Stack[string]` ensures type consistency/safty at compile time.

### Methods that are identical for all types

When implementing a method like `Swap`, `Length`, or `Less` that behaves the same way for all slice types, generics are ideal.

```go
type SliceFn[T any] struct {
    s    []T
    less func(T, T) bool
}

func (s SliceFn[T]) Swap(i, j int) {
    s.s[i], s.s[j] = s.s[j], s.s[i]
}
```

This is especially useful when implementing interfaces like `sort.Interface` with different types.

### Ensuring consistent type usage between inputs and outputs

Another case is when a function needs to accept and return the same type—without knowing which one ahead of time. This allows Go to infer types automatically and reduce the need for boilerplate.

### Do not use when

- Interfaces are enough to provide needed flexibility.
- Method implementations differ based on type.

## Example

In the example code (`main.go`) generics are the only way to implement the wanted functionality in an efficient and type-safe way.
The code features comments explaining why.

## Summary

Generics in Go allow you to write clean, reusable, and type-safe code.

When writing data structures and functions that repeat across types, it is best to use generics or interfaces.
Use generics when the logic is the same, and use interfaces when the behavior (methods) is the focus.

For more information, see these official resources:

- [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
- [Blog: When To Use Generics](https://go.dev/blog/when-generics)
- [Generics in the Go spec](https://go.dev/ref/spec#Type_parameter_declarations)
