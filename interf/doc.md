# Interfaces

A Golang interface is a custon type from a collection of method signatures. It specifies the behavoiour but does not implement a type.

## Implementation

Unlike other object-oriented progamming language, a Go interface does not require explicit implementation type. A variable of a given interface can receive any struct or function types that implicitly implemented all method signatures of the interface.

Here is a [working example](./ex1/ex1_test.go).

## Embeded interface

These examples demonstrate technique composition to overide behaviour.

Here is a [working example](./ex3/ex3_test.go)

## Useful references

* [What Makes Golang Go: The Power of Go Interfaces — Ricardo Gerardi](https://www.youtube.com/watch?v=TRoRluGIixs)
* [Ardan Labs - interfaces 101](https://www.youtube.com/watch?v=34ZmIfWOb0U&list=PLADD_vxzPcZB595tXmu540KC6MTMqIndB)
