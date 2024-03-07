# Embedding Files

This section discuss techniques using the `embed` package use to embed and read files.

## Mapping to []byte 

This example demonstrates a technique to map a specific file to []byte. Please refer to this [example](./ex1/main_test.go)

## Mapping to embed.FS

This example maps embed files to `embed.FS` type and the process to read content from `embed.FS`. Please refer to this [example](./ex2/main_test.go)

## Mapping to file or web server

You can use this to deliver via a fileserver or webserver. Please refer to the [networking/webserver](../networking/doc.md) to illustrate it as part of a fileserver.