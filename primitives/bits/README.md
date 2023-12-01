# Binary, Hex, Octal and Bytes

This chapter discuss representation of Bits, Hex and Bytes literals. These literals are `uint8`, `int` or `unsigned int` types.

## Binary

Binary literals are represented with the prefix `0b` or `0B`.

To get a string version of binary use the following string format `fmt.Printf("%0[number of bits]", ...)`. For example to only show 3 bits binary run `fmt.Printf("%3", ...)`

See [Working example](./literals_test.go)

## Octal

Octals are base 8 numbers (0-7) and are represented with the zero in front of an integer. To get a string version of an Octal use the following string format `fmt.Print("%0o", ...)`

See [Working example](./literals_test.go)

## Hex

Hex are base 16 numbers (0-9ABCDEF) and are represented with a prefix 0x or 0X.

To get a string version of an Octal use the following string format `fmt.Print("%X", ...)` to display [A-F] or `fmt.Print("%x", ...)` to display [a-f].

See [Working example](./literals_test.go)

## Bytes

Bytes are `uint8` types or alias `byte` type.

See [Working example](./literals_test.go)

## Bitwise operations

Please refer to this [cheat sheet](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/) for details.

See [Working examples](./bitops_test.go)