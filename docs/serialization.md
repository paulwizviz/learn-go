# Serialization

This section discusses techniques used to serialised data

## ASN.1

Abstract Syntax Notation One (ASN.1) is a standard interface description language for defining data structures that can be serialized and deserialized in a cross-platform way. It is broadly used in telecommunications and computer networking, and especially in cryptography.

* Source
    * [Introduction to ASN.1](https://www.itu.int/en/ITU-T/asn1/Pages/introduction.aspx)
    * [OSS Nokalva - ASN](https://www.oss.com/resources/resources.html)
    * [ASN1 Simple types](https://www.obj-sys.com/asn1tutorial/node10.html)
    * [A Layman's Guide to a Subset of ASN.1, BER, and DER](http://luca.ntop.org/Teaching/Appunti/asn1.html)

* Executable examples
    * [asn_test.go](../example/serial/asnser/asn1_test.go)

## Base64

In computer programming, Base64 is a group of binary-to-text encoding schemes that represent binary data (more specifically, a sequence of 8-bit bytes) in sequences of 24 bits that can be represented by four 6-bit Base64 digits[wiki](https://en.wikipedia.org/wiki/Base64)).

* Source
    * [The Base16, Base32, and Base64 Data Encodings](https://datatracker.ietf.org/doc/html/rfc4648)

* Executable examples
    * [base64_test.go](../example/serial/base64ser/base64_test.go)

## CBOR

The Concise Binary Object Representation (CBOR) -- RFC 8949-- is a data format whose design goals include the possibility of extremely small code size, fairly small message size, and extensibility without the need for version negotiation[CBOR](https://cbor.io/).

* Executable examples
    * [cbor_test.go](../example/serial/cborser/cbor_test.go)