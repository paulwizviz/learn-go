# Functions

A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is nil.

Within a list of parameters or results, the names must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-blank names in the signature must be unique. If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

The final incoming parameter in a function signature may have a type prefixed with .... A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.

```go
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

See [The Go Programming Language Specification - Function Types](https://go.dev/ref/spec#Function_types)

Please see this [working example](funcs_test.go)

## init

A function with the name `init` is intended to facilitate the initialisation of variables in a package

Please see this [working example](./init/main.go)

## defer

A function call within another function can be deferred.

Please refer to this [working example](./deferops_test.go)