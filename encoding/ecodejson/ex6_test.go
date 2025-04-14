package encodejson

import (
	"encoding/json"
	"fmt"
	"os"
)

type Order struct {
	Base     string   `json:"base"`
	Proteins []string `json:"proteins"`
}

func Example_unarshalEx5() {
	f, err := os.Open("./ndjson")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var order Order
		if err := decoder.Decode(&order); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", order)
	}

	// Output:
	// {Base:white rice Proteins:[tofu]}
	// {Base:salad Proteins:[tuna salmon]}
}
