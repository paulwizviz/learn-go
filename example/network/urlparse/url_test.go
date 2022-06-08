package urlparse

import (
	"fmt"
	"net/url"
)

const (
	urlString = "postgres://user:pass@host.com:5432/path?k=v#f"
)

func Example_scheme() {
	urlElems, _ := url.Parse(urlString)
	scheme := urlElems.Scheme
	fmt.Println(scheme)

	// Output:
	// postgres
}

func Example_user() {
	urlElems, _ := url.Parse(urlString)
	user := urlElems.User
	fmt.Printf("User: %v\n", user)

	fmt.Printf("Username: %v\n", user.Username())

	value, ok := user.Password()
	fmt.Printf("Value: %v Ok: %v", value, ok)

	// Output:
	// User: user:pass
	// Username: user
	// Value: pass Ok: true
}

func Example_path() {
	urlElems, _ := url.Parse(urlString)
	path := urlElems.Path
	fmt.Printf("Path: %v", path)

	// Output:
	// Path: /path
}

func Example_query() {
	urlElems, _ := url.Parse(urlString)
	queryValues := urlElems.Query()
	fmt.Println(queryValues)

	// Output:
	// map[k:[v]]
}

func Example_fragment() {
	urlElems, _ := url.Parse(urlString)
	frag := urlElems.Fragment
	fmt.Println(frag)

	// Output:
	// f
}
