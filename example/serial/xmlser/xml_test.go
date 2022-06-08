package xmlser

import (
	"encoding/xml"
	"fmt"
)

func Example_marshal() {
	p := Person{
		ID:        1,
		FirstName: "John",
		Surname:   "Doe",
		Address: Address{
			ID:     2,
			Sub:    "#1",
			Street: "hello street",
		},
	}

	pXML, _ := xml.Marshal(p)
	fmt.Println(string(pXML))
	// Output:
	// <person id="1"><firstname>John</firstname><surname>Doe</surname><address id="2" sub="#1"><street>hello street</street></address></person>
}

func Example_unmarshal() {
	data := []byte("<person id=\"1\"><firstname>John</firstname><surname>Doe</surname><address id=\"2\" sub=\"#1\"><street>hello street</street></address></person>")
	var p Person
	xml.Unmarshal(data, &p)
	fmt.Println(p)
	// Output:
	// {{ person} 1 John Doe {{ address} 2 #1 hello street}}
}

type Writer struct {
	Header string
}

func (w Writer) Write(p []byte) (int, error) {
	fmt.Printf("%s%s", w.Header, string(p))
	return 0, nil
}

func Example_encoder() {
	p := Person{
		ID:        1,
		FirstName: "John",
		Surname:   "Doe",
		Address: Address{
			ID:     2,
			Sub:    "#1",
			Street: "hello street",
		},
	}
	w := Writer{
		Header: xml.Header,
	}
	xml.NewEncoder(w).Encode(p)

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <person id="1"><firstname>John</firstname><surname>Doe</surname><address id="2" sub="#1"><street>hello street</street></address></person>
}
