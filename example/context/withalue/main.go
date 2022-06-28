package main

import (
	"context"
	"fmt"
)

type key string
type value string

func main() {
	ctx := context.Background()

	k := key("hello")
	v := value("word")
	ctx = context.WithValue(ctx, k, v)
	fmt.Println(ctx.Value(k))
}
