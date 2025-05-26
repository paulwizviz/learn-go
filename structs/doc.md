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
* method argument and return types

### Method composition

In the context of struct composition:

Here are [working examples](./methods_test.go)

## References

* [Go Class: 19 Composition](https://www.youtube.com/watch?v=0X6AcnwocbM)
