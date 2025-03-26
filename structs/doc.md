# Struct types

A a struct type is a collection of primitive, struct types and function types as fields. A struct may also include [member methods](https://go101.org/article/method.html).
 
## Equality

[Example of equality](./equality_test.go)

## Embedded struct

You can embed a struct within a struct. Here is a [working example](./embedded_test.go)

## Member method

A member method of a struct is a syntatic sugar for a normal function.

Here is a [working example](./methods_test.go)