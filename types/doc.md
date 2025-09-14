# Basic Types

In this section of our journal, we turn our focus to the fundamentals: Go's basic types. The following notes and examples are our observations on how these types behave. All our examples have been validated against the Go version specified in [go.mod](./go.mod).

* [Array](#array)
* [Bitwise Operators](#bitwise-operators)
* [Byte](#byte)
* [Float](#float)
* [Integer](#integer)
* [Octal](#octal)
* [Hexadecimal](#hexadecimal)
* [Rune](#rune)
* [Slice](#slice)
* [String](#string)

## Array

For the **array**, our initial observation was that it's a fixed-size container. It appeared that every element must be of the same single type, and we concluded this fixed nature is its key characteristic. Our findings are documented in the [working example](./array_test.go).

## Binary

For **binary** numbers, we observed that Go allows for a base-2 representation of integers. It appeared that to define an integer using only ones and zeros, we had to use the `0b` or `0B` prefix. We concluded that this prefix signals the compiler to interpret the number as binary. Our findings on this are in the [working examples](./binary_test.go).

## Bitwise Operators

For **bitwise operations**, our exploration led us to situations requiring work at a lower level. We found that to manipulate the individual bits of an integer, Go provides a specific set of operators.

Our notes identify the following operators:

* AND `&`
* OR `|`
* XOR `^`
* AND NOT `&^`
* Left shift `<<`
* Right shift `>>`

Our findings are documented in the [working examples](./bitops_test.go). We also concluded that the [Bitwise operators cheat sheet](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/) was a useful external resource for further reading.

## Byte

For the `byte` keyword our initial thought was that it might be a distinct type, but it appears to be a simple alias type for `uint8`. We concluded it's used to represent an 8-bit unsigned integer, which seems perfect for raw binary data. Our observations are in the [working example](./byte_test.go).

## Float

For **floating-point** numbers, our observations focused on values that are not whole. We found that Go appears to provide two main types, `float32` and `float64`. We concluded that `float64` is the default, as it's the most commonly used unless memory conservation is a specific goal. Our investigation of float literals showed they are straightforward, supporting decimal points and `e` or `E` for exponent notation. Our findings are documented in the [working example](./float_test.go).

## Integer

For **integers**, our initial observation was that they represent whole numbers, with a literal being a simple sequence of digits like `123`.

Our investigation then focused on the variety of integer types. We found that Go appears to offer a high degree of control over an integer's size and range, providing types for 8, 16, 32, and 64-bit values. We concluded that each size comes in two distinct flavors: `signed` for numbers that can be positive or negative, and `unsigned` for numbers that are strictly non-negative.

Our findings on this breakdown are as follows:

* **8-bit:**  `int8` or `uint8`
* **16-bit:** `int16` or `uint16`
* **32-bit:** `int32` or `uint32`
* **64-bit:** `int64` or `uint64`

These observations are documented in our [working examples](./integer_test.go).

## Octal

For **octal** numbers, our investigation found that Go supports a base-8 representation. It appears that an octal literal can be prefixed with `0o` or `0O`. We also observed a legacy prefix, a single `0`, which can be a source of confusion. We concluded that while it's a supported feature, its use seems less common in modern Go code compared to hexadecimal or binary representations. Our findings are documented in the [working examples](./octal_test.go).

## Hexadecimal

For **hexadecimal** numbers, our investigation found that Go supports a base-16 representation for both integers and floats. This representation struck us as particularly useful when working with bitmasks or other low-level data, where base-16 is more common.

Our findings for both forms are documented in the [working examples](./hex_test.go).

### Hexadecimal Integers

Our observation for integers was that the literal format is straightforward, requiring a `0x` or `0X` prefix followed by hexadecimal digits (`0-9`, `a-f`).

### Hexadecimal Floats

For floats, we found the structure to be more detailed, adhering to the IEEE 754-2008 standard. Our conclusion was that a literal requires three parts:

1. A `0x` or `0X` prefix.
2. A mantissa (the number's significant digits).
3. A mandatory exponent, prefixed with `p` or `P`, which scales the mantissa by a power of 2.

## Rune

The concept of a `rune` was an interesting discovery. It appears to be Go's way of handling a single character or, more precisely, a Unicode code point. We found that the `rune` type is just an alias for `int32`. A rune literal is usually written in single quotes, like `'x'`. We've collected our notes on the various ways to express them, including special characters, in our [working example](./rune_test.go).

## Slice

`Slice` was one of the most interesting types we investigated. Our initial observation is that it's a more flexible version of an array. The documentation describes it as a "descriptor for a contiguous segment of an underlying array," which suggests it's a lightweight structure that points to a piece of an array. This seems to be a core, powerful concept in Go. Our initial findings are in the [working example](./slice_test.go).

* References
  * [Go Class: 10 Slices in details](https://www.youtube.com/watch?v=pHl9r3B2DFI)
  * [Memory layout and mechanics of arrays and slices | Golang | intermediate level](https://www.youtube.com/watch?v=RVTfPy_NELc)

## String

Finally, we looked at strings. The most important thing we observed is that in Go, a string seems to be an immutable slice of bytes. This implies that they are read-only and that when we see string indexing, we are accessing individual bytes, not characters. We found two ways to write them: interpreted (double quotes) and raw (backticks). Our notes on this are in the [working example](./string_test.go).

* References
  * [Strings, Bytes and Runes | Golang | intermediate level](https://www.youtube.com/watch?v=fVbI_3v0Zys&list=PLSozy2hb5kKPpIJnpZ2sSMfjVjP0tyJYG&index=2)
  * [Strings, Bytes and Runes | Golang | intermediate level](https://www.youtube.com/watch?v=pHl9r3B2DFI)