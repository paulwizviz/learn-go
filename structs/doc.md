# Struct

A **struct** is a composite type that groups together variables (called fields) under a single name. You can also assign a method against a struct.

## Comparison

Structs are comparable (i.e. can be used with == and !=) **if and only if all their fields are comparable**.

Structs **cannot be compared** if they contain **non-comparable** fields, like:

* slices (`[]T`)
* maps (`map[K]V`)
* functions
* channels (some edge cases)
* structs containing any of the above

How to compare uncomparable fields?

Two options:

* Manually compare
* Use DeepEqual from `reflect.DeepEqual`

> NOTE: `reflect.DeepEqual` is convenient but slower, and has some subtleties (e.g., it considers nil and empty slices not equal).

Here are some [working examples](./compare_test.go)

## Composition

Struct composition involves embedding a struct in another struct.

When you embed a struct within another struct, the fields of the embedded struct is "promoted" but not "inherited".

Here are [working examples](./composition_test.go).

## A Method

A **method** is a function that is associated with a specific type—either a struct or a named type. The key components of a method are:

* reciever typea
* method name
* method with zero or more arguments and return types

### Receiver types

A receiver type can be a pointer or a value type.

```go
type Data struct{
    n string
}

// In this case, d is a value type
func (d Data) DoSomethig(){
    d.n = "Hello" // This will not update the field in Data struct
}

// In this case, d is a pointer type
func (d *Data) DoSomethingElse(){
    d.n = "Ola" // This will update the field in Data struct
}
```

>Notes:
>
> * Use pointer receiver to mutate the associated type
> * Use value receiver if you do not plan to mutate the associated type

### Method expression

A method expression in Go is a way to refer to a method as a standalone function — one where the receiver is passed explicitly as the first argument.

```go
type Real struct{}

func (r Real) DoSomething() {
}

func (r *Real) DoAnother() {
}

// Calling method as a standalone function where
// f and f1 are function types base on methods
f := (Real).DoSomething
f1 := (*Real).DoAnother

// calling f and f1
r := Real{}
f(r)
f1(&r)

f2 := r.DoSomething
f3 := r.DoAnother
f2() // this is the equivalent of f(r)
f3() // this is the equivalent of f(&r)

```

### Method composition

In the context of struct composition:

Here are [working examples](./methods_test.go)

## References

* [Go Class: 19 Composition](https://www.youtube.com/watch?v=0X6AcnwocbM)
* [Methods in Golang | Intermediate level](https://www.youtube.com/watch?v=bAWI0NVZlkc)
