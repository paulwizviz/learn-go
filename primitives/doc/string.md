# Rune and Strings

Please refer to the following concepts to appreciate the inner workings of rune and string:

* [Strings, bytes, runes and characters in Go by Rob Pike](https://go.dev/blog/strings)
* [ASCII, Unicode, UTF-8: Explained Simply](https://www.youtube.com/watch?v=DntKZ9xJ1sM)

## Rune

A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' or '\n'. Within the quotes, any character may appear except newline and unescaped single quote. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats[1](https://go.dev/ref/spec).

The `rune` keyword is an alias for uint32 

Rune may also contain special characters and these are:

```
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000B vertical tab
\\   U+005C backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
```

Please refer to this [working example](../str/rune_test.go)

## String

A string literal represents a string constant obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals[1](https://go.dev/ref/spec)).

A string is a slice of bytes

Raw string are characters sequence in back quotes example \`raw string\`. Interpreted string literals are character sequences between double quotes, as in "bar" 

Please refer to this to understand the internals of [Strings, Bytes and Runes | Golang | intermediate level](https://www.youtube.com/watch?v=fVbI_3v0Zys&list=PLSozy2hb5kKPpIJnpZ2sSMfjVjP0tyJYG&index=2)
