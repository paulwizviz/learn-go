# Integer Literals

An integer literal is a sequence of digits representing an integer constant. The following are types of integers:

* 8bits  - `int8` or `uint8`
* 16bits -`int16` or `uint16`
* 32bits - `int32` or `uint32`
* 64bits - `int64` or `uint64` 

Please refer to this [working examples of integer literals](./integer_test.go).

## Binary

A binary is a base 2 integer. A binary literal consists of the prefix `0b` or `0B` followed by a series of `0`s and `1`s. 

Please refer to this [working examples of binary](./binary_test.go).

## Octal

An Octal is a base 8 integer. An octal literal consists of the prefix `0`, `0o`, or `0O` followed by a series of numbers `0-7`. 

Please refer to this [working examples of octal](./octal_test.go)

## Hex

A hexadecimal which is a base 16 integer. A hexadecimal literal consists of a prefix `0x` or `0X` followed by a combination of `0-9` or `A`(`a`) or `B`(`b`) or `C`(`c`) or `D`(`d`) or `E`(`e`) or `F`(`f`)  

Please refer to this [working exampls of hex](./hex_test.go)

## Byte

The key word `byte` is an alias for `uint8` represents 8 bit numeric value. 

Please refer to this [working example](./byte_test.go)

## Bitwise operators

The following are operators to enable you to perform bitwise operations:

* AND `&`
* OR `|`
* XOR `^`
* AND NOT `&^` 
* Left shift `<<`
* Right shift `>>`

Please refer to this [working examples of bit wise operations](./bitops_test.go)

## References

* [Bitwise operators [cheat sheet]](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/) to learn more about bit operations such as shift left or right.