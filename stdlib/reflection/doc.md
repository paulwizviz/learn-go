# Reflection

Reflection in computing is the ability of a program to examine its own structure, particularly through types; it’s a form of metaprogramming [Source: The Laws of Reflection](https://go.dev/blog/laws-of-reflection).

There are three rules to reflection:

* Rule 1 - Reflection goes from interface to reflection object.
* Rule 2 - Reflection goes from object to interface value.
* Rule 3 - To modify reflection object, value must be settable

## Indirect

* Purpose:
	* Returns the reflect.Value that v points to.
	* If v is not a pointer, reflect.Indirect() returns v itself.
	* It is used to get the value that a pointer reflects to.
* Use Cases:
	* Working with pointers and interfaces that might contain pointers.
	* Dereferencing pointers to access the underlying values.

## Typeof

* Purpose:
	* Returns the `reflect.Type` of the value passed to it.
	* The `reflect.Type` interface provides methods for inspecting the type's properties (e.g., name kind, methods, fields).
	* It tells you about the static type of the value, or the dynamic type of the value held by an interface.   
* Use Cases:
	* Inspecting the structure of a type.
	* Determining the type's kind (e.g., struct, slice, map).   
	* Retrieving information about a type's methods or fields.
* [Working Example](typeof_test.go)

## ValueOf

* Purpose:
	* Returns the `reflect.Value` of the value passed to it.
	* The `reflect.Value` struct provides methods for accessing and manipulating the value itself.
	* It is used to get the actual value that is passed in.   
* Use Cases:
	* Accessing and modifying the value of variables.
	* Calling methods of a value.
	* Setting the fields of a struct.
* [Working Example](valueof_test.go)

## Useful Resources

* [The Laws of Reflection](https://www.youtube.com/watch?v=Jvask1Hq_KE)
* [Learn to Use Go Reflection](https://www.youtube.com/watch?v=-U1oOmbKkx4)