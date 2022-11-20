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

Here are [working examples demonstrating recovery from panic](../example/error/recovery/recovery_test.go)

## Handling errors

There are two approaches to handling errors:

* Custom wrappers [working examples](../example/error/wrap/customs_test.go)
* Using [standard package errors](../example/error/wrap/errors_test.go)
