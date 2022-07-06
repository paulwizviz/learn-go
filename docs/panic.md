# Defer and Panic

This section illustrates technique to defer operations and recover from panic.

The pattern recover from panic is:

```
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered")
    }
}()
```

## Working examples

* [Deferring operations](../example/panic/deferops/deferops_test.go)
* [Non recovered panic](../example/panic/norecovery/main.go)
* [Recover from panic](../example/panic/recovery)