# Interface

An **interface** is a type that specifies method signatures. Any type that implements those methods is said to **satisfy** the interface.

## Single Method Interfaces

A single method interface is an interface with a single method.

Why use single method interface?

* They’re perfect for mocking, decoupling, and interface-based design.
* Go’s idiomatic use of interfaces encourages “small interfaces”, often with just one method.

| Benefit | Why It Matters |
| --- | --- |
| Loose coupling | Use only what you need |
| Composability | Build bigger interfaces from small ones |
| Easier testing | Mocks are quick and isolated |
| Implicit satisfaction | Less boilerplate |
| Cleaner APIs | Smaller, more focused contracts |

There are a number of single method interfaces in the standard library.

### `io` packages

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

```go
type Closer interface {
    Close() error
}
```

```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

```go
type ByteReader interface {
    ReadByte() (byte, error)
}
```

```go
type ByteWriter interface {
    WriteByte(c byte) error
}
```

### `fmt` package

```go
type Stringer interface {
    String() string
}
```

### `net/http` package

```go
type Handler interface {
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```

### `errors` package

```go
type Wrapper interface {
    Unwrap() error
}
```

## Implement Any Types by Interfaces

Any Go types can be implement an Go interfaces. You can implement methods against:

* Alias native types like `int`, `float32`, etc
* Struct types
* Function types

Here is a [working example](./implement_test.go).

## Grouping Types by Interfaces

We can use interfaces group types by methods. For example, we could define aggregate types by interface.

Here is the [working example](./grouping_test.go).

## Embedding Interfaces

Embedding an interface in a struct is a way to delegate method calls to a concrete implementation of that interface. However, embedding interfaces in structs is **not** idiomatic in Go. The preferred approach is to define interfaces where they are needed, usually on the **consumer** side.

There are some cases where embedding an interface might be acceptable:

* If you’re writing a framework or middleware layer, and want to provide default forwarding behavior.
* If you are building decorators/wrappers where forwarding is natural.
* In tests, to simplify mock composition.

But even in those cases, named fields are usually clearer.

## References

* [What Makes Golang Go: The Power of Go Interfaces — Ricardo Gerardi](https://www.youtube.com/watch?v=TRoRluGIixs)
