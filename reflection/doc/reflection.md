# Reflection

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it’s a form of metaprogramming [Source: The Laws of Reflection](https://go.dev/blog/laws-of-reflection).

## Useful Resources

* [The Laws of Reflection](https://www.youtube.com/watch?v=Jvask1Hq_KE)
* [Learn to Use Go Reflection](https://www.youtube.com/watch?v=-U1oOmbKkx4)

## Rules of Reflection

* Rule 1  -Reflection goes from interface to reflection object.
* Rule 2 - Reflection goes from object to interface value.
* Rule 3 - To modify reflection object, value must be settable

<u>TypeOf</u>

TypeOf returns the reflection Type that represents the dynamic type of i. If i is a nil interface value, TypeOf returns nil.

```go
func TypeOf(i any) Type {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	return toType(eface.typ)
}
```

An empty interface has two components:

```go
// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}
```

typ -- refers to a struct
word --- refers to the value

Here are [working examples](../examples/typeof/main_test.go)

<u>valueOf</us>

ValueOf returns a new Value initialized to the concrete value stored in the interface i. ValueOf(nil) returns the zero Value.

Value is the reflection interface to a Go value.

Not all methods apply to all kinds of values. Restrictions, if any, are noted in the documentation for each method. Use the Kind method to find out the kind of value before calling kind-specific methods. Calling a method inappropriate to the kind of type causes a run time panic.

The zero Value represents no value. Its IsValid method returns false, its Kind method returns Invalid, its String method returns "<invalid Value>", and all other methods panic. Most functions and methods never return an invalid value. If one does, its documentation states the conditions explicitly.

A Value can be used concurrently by multiple goroutines provided that the underlying Go value can be used concurrently for the equivalent direct operations.

To compare two Values, compare the results of the Interface method. Using == on two Values does not compare the underlying values they represent.

A value type implementation is:

```go
type Value struct {
	// typ holds the type of the value represented by a Value.
	typ *rtype

	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
	ptr unsafe.Pointer

	// flag holds metadata about the value.
	//
	// The lowest five bits give the Kind of the value, mirroring typ.Kind().
	//
	// The next set of bits are flag bits:
	//	- flagStickyRO: obtained via unexported not embedded field, so read-only
	//	- flagEmbedRO: obtained via unexported embedded field, so read-only
	//	- flagIndir: val holds a pointer to the data
	//	- flagAddr: v.CanAddr is true (implies flagIndir and ptr is non-nil)
	//	- flagMethod: v is a method value.
	// If ifaceIndir(typ), code can assume that flagIndir is set.
	//
	// The remaining 22+ bits give a method number for method values.
	// If flag.kind() != Func, code can assume that flagMethod is unset.
	flag

	// A method value represents a curried method invocation
	// like r.Read for some receiver r. The typ+val+flag bits describe
	// the receiver r, but the flag's Kind bits say Func (methods are
	// functions), and the top bits of the flag give the method number
	// in r's type's method table.
}
```

Here are [working examples](../examples/valueof/main_test.go)

<u>Indirect<u>

Indirect returns the value that v points to.
If v is a nil pointer, Indirect returns a zero Value.
If v is not a pointer, Indirect returns v.

The implementation is:

```
func Indirect(v Value) Value {
	if v.Kind() != Pointer {
		return v
	}
	return v.Elem()
}
```

Here are [working examples](../examples/indirect/main_test.go)

<u>Inspection</u>

Here are [working examples](../examples/inspection/main_test.go) demonstrating techniques to inspecting struct types