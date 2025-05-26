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

## References

* [What Makes Golang Go: The Power of Go Interfaces — Ricardo Gerardi](https://www.youtube.com/watch?v=TRoRluGIixs
