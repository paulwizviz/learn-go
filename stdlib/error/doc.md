# Error and Panic

This section illustrates techniques for handling errors and panic.

## Panic and recover

The pattern to recover from panic is:

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered")
    }
}()
```

[working examples](./recovery_test.go)

## Error wrappers

[working examples](./customs_test.go)
[standard package errors](./errors_test.go)
