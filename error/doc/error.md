# Error and Panic

This section illustrates techniques for handling errors and panic.

## Panic and recover

The pattern to recover from panic is:

```
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered")
    }
}()
```

[working examples](../examples/recovery/recovery_test.go)

## Error wrappers

[working examples](../examples/wrap/customs_test.go)
[standard package errors](../examples/wrap/errors_test.go)

