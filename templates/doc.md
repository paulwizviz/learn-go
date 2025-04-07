# Templates

Go provides data driven templates for generating textual output.

## Working Examples

* [Conditional](./conditional_test.go)
* [Pipes](./pipping_test.go)
* [Range](./range_test.go)
* [Reuse](./reuse_test.go)
* [Structs](./structs_test.go)

## Inbuilt Functions

The following are inbuilt functions in Go templates.

### Comparison Functions:

* `eq`: Returns true if the arguments are equal (arg1 == arg2).
* `ne`: Returns true if the arguments are not equal (arg1 != arg2).
* `lt`: Returns true if the first argument is less than the second (arg1 < arg2).
* `le`: Returns true if the first argument is less than or equal to the second (arg1 <= arg2).
* `gt`: Returns true if the first argument is greater than the second (arg1 > arg2).
* `ge`: Returns true if the first argument is greater than or equal to the second (arg1 >= arg2).

### Boolean Functions:

* `and`: Returns the boolean AND of its arguments.
* `or`: Returns the boolean OR of its arguments.
* `not`: Returns the boolean negation of its argument.

### Indexing Functions:

* `index`: Returns the result of indexing its first argument by the following arguments. Example: {{index .Slice 2}} or {{index .Map "key"}}.

### Formatting Functions:

* `printf`: Returns a formatted string using the standard Go fmt.Sprintf function.

### String Functions:

* `print`: Equivalent to fmt.Sprint.
* `println`: Equivalent to fmt.Sprintln.
* `printf`: Equivalent to fmt.Sprintf.
* `html`: Returns the escaped HTML equivalent of its arguments.
* `js`: Returns the JavaScript escaped equivalent of its arguments.
* `urlquery`: Returns the escaped value of its arguments in a url query.

### Data Context

The . (dot) represents the current data context.

### Other Functions

* `len`: Returns the length of its argument (string, slice, map, etc.).
* `call`: Calls the function whose name is the first argument with the remaining arguments as parameters.
* `default`: Returns the second argument if the first is empty, otherwise returns the first argument.
* `slice`: returns a slice.
* `fileExists`: returns if a file exists.

### Conditional

`{{if}}`, `{{else}}`, `{{else if}}`, and `{{end}}` are template actions that control the flow of execution.

## References

* [Discover Go template syntax](https://developer.hashicorp.com/nomad/tutorials/templates/go-template-syntax)