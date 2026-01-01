# Custom Types

Having explored Go's basic types, our journey now takes us into the realm of creating our own: **custom types**. In this part of our journal, we document our experiences with `interfaces` and `structs`, which we found to be the fundamental building blocks for structuring complex programs in Go. All our examples have been validated against the Go version specified in [go.mod](./go.mod).

* [Interface](#interface)
* [Struct](#struct)

## Generics

As our journey into Go's type system deepened, we encountered a feature that promised to elevate our code's flexibility and reusability: **generics**. This powerful addition, introduced in Go 1.18, allowed us to write code that could work with any of a set of types, without sacrificing type safety. We found this to be a significant evolution in how we could structure our programs, enabling us to create more abstract and versatile functions and data structures. Here is [Working Example](./generics_test.go)

## Interface

Our first exploration into custom types was the **interface**. We came to understand it not as a type that holds data, but as one that defines behaviour. It appeared to be a contract, specifying a set of method signatures. Our key observation was that any type that implements these methods automatically and implicitly *satisfies* the interface, a concept that struck us as both powerful and elegant.

### Single-Method Interfaces

A recurring pattern we noticed during our investigation was the prevalence of interfaces with just a single method. We concluded that this is an idiomatic practice in Go, favouring small, focused contracts. This approach, we found, leads to more decoupled and composable code.

During our survey of the standard library, we encountered numerous examples of this principle. In the `fmt` package, for instance, we found the `Stringer` interface, which we found essential for creating custom string representations of our types:

```go
type Stringer interface {
    String() string
}
```

We realised that by implementing just one method, a type could participate in a wide range of standard library functionalities.

### Implementing Interfaces

A significant discovery for us was that *any* type in Go can satisfy an interface. We experimented by attaching methods to various types, including alias types based on primitives like `int`, as well as `struct` and even `function` types. This confirmed for us that interfaces are a truly general-purpose mechanism for abstracting behaviour. Our findings are documented in this [working example](./implement_test.go).

### Grouping with Interfaces

We also observed that interfaces provide a powerful way to group different types that share a common behaviour. By defining an interface, we could treat a collection of disparate types in a uniform manner, focusing only on the methods we needed. This allowed us to write more generic and flexible code. We explored this concept of aggregation in our [working example](./grouping_test.go).

### Embedding Interfaces

Our journey led us to experiment with embedding interfaces within structs. Our initial thought was that this would be a straightforward way to delegate behaviour. However, we soon learned that this is generally not considered idiomatic in Go. The prevailing wisdom we encountered suggests that interfaces are best defined by the consumer of a type, not embedded within the producer. While we noted some niche use cases, such as in frameworks or mocks, we concluded that favouring composition with named fields over interface embedding leads to clearer and more maintainable code. Our experiments are documented in this [working example](./embedded_test.go).

## Struct

Our next focus was the **struct**, which we found to be Go's primary tool for creating composite data types. We observed that it allows grouping together different fields, each with its own type, under a single, new type name. This appeared to be the fundamental way to model custom data structures.

### Comparing Structs

An interesting aspect we investigated was how structs are compared. Our observation was that a struct is comparable with `==` or `!=` only if all of its fields are themselves comparable. We found that fields like slices, maps, and functions render a struct non-comparable. For these cases, we learned that a manual, field-by-field comparison or the `reflect.DeepEqual` function were the necessary alternatives. Our notes on this are in these [working examples](./compare_test.go).

### Struct Composition

We then explored struct **composition**, which involves embedding one struct within another. Our key finding here was that the fields and methods of the embedded struct are "promoted" to the containing struct, making them directly accessible. We concluded that this is Go's powerful alternative to classical inheritance, favouring a "has-a" relationship over an "is-a" one. We documented our exploration of this concept in these [working examples](./composition_test.go).

### Methods

Attaching behaviour to our custom `struct` types led us to **methods**. We understood a method to be a function associated with a specific type, which is known as the receiver.

#### Receiver Types

Our investigation revealed two types of receivers: value receivers and pointer receivers. We concluded that the choice between them is critical: a pointer receiver (`func (d *Data)`) is necessary to mutate the receiver's data, whereas a value receiver (`func (d Data)`) operates on a copy and cannot change the original value.

#### Method Expressions

An interesting discovery was the **method expression**. This allowed us to treat a method as a standalone function where the receiver is passed as an explicit first argument. We found this useful for scenarios requiring function-like references to methods.

Our complete findings on methods, including composition, are documented in these [working examples](./methods_test.go).

## Further Reading

During our exploration, we found the following external resources particularly insightful:

* [Go Class: 19 Composition](https://www.youtube.com/watch?v=0X6AcnwocbM)
* [Methods in Golang | Intermediate level](https://www.youtube.com/watch?v=bAWI0NVZlkc)
* [What Makes Golang Go: The Power of Go Interfaces — Ricardo Gerardi](https://www.youtube.com/watch?v=TRoRluGIixs)