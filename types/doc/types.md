# User Defined Data Types

This secton discuss aspects of Go you use to specify and implement custom data types.

## Struct types

A a struct type is a collection of primitive, struct types and function types as fields. A struct may also include [member methods](https://go101.org/article/method.html).
 
<u>Embedded struct</u>

You can embed a struct within a struct. Here is a [working example](../examples/struct/ex1/main.go)

<u>Member method</u>

A member method of a struct is a syntatic sugar for a normal function.

Here is a [working example](../examples/struct/ex2/main.go)


## Interfaces

A Golang interface is a custon type from a collection of method signatures. It specifies the behavoiour but does not implement a type.

<u>Useful references</u>

* [What Makes Golang Go: The Power of Go Interfaces — Ricardo Gerardi](https://www.youtube.com/watch?v=TRoRluGIixs)
* [Ardan Labs - interfaces 101](https://www.youtube.com/watch?v=34ZmIfWOb0U&list=PLADD_vxzPcZB595tXmu540KC6MTMqIndB)

<u>Interfaces and implementation</u>

Unlike other object-oriented progamming language, a Go interface does not require explicit implementation type. A variable of a given interface can receive any struct or function types that implicitly implemented all method signatures of the interface.

Here is a [working example](../examples/interf/ex1/main.go).

<u>Single method interfaces</u>

An interface with only one method signature, can be implemented by a struct with member methods and function types.

Here is a [working](../examples/interf/ex2/main.go).



