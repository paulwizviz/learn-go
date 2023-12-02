# Integer

An integer literal is a sequence of digits representing an integer constant. An optional prefix sets a non-decimal base: 0b or 0B for binary, 0, 0o, or 0O for octal, and 0x or 0X for hexadecimal. A single 0 is considered a decimal zero. In hexadecimal literals, letters a through f and A through F represent values 10 through 15[1](https://go.dev/ref/spec).

Here are [working examples](../integer/integer_test.go) of integer literals and variable types

## Binary

Binary literals are represented with the prefix `0b` or `0B`.

To get a string version of binary use the following string format `fmt.Printf("%0[number of bits]", ...)`. For example to only show 3 bits binary run `fmt.Printf("%3", ...)`

See [Working example](../integer/binary_test.go)

## Octal

Octals are base 8 numbers (0-7) and are represented with the zero in front of an integer. To get a string version of an Octal use the following string format `fmt.Print("%0o", ...)`

See [Working example](../integer/octal_test.go)

## Hex

Hex are base 16 numbers (0-9ABCDEF) and are represented with a prefix 0x or 0X.

To get a string version of an Octal use the following string format `fmt.Print("%X", ...)` to display [A-F] or `fmt.Print("%x", ...)` to display [a-f].

See [Working example](../integer/hex_test.go)

## Bytes

Bytes are `uint8` types and there is an alias `byte`.

See [Working example](../integer/byte_test.go)

## Bitwise operations

Please refer to this [cheat sheet](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/) for details.

See [Working examples](../integer/bitops_test.go)