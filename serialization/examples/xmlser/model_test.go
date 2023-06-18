package xmlser

import "encoding/xml"

// <person id="1">
//   <firstname>John</firstname>
//   <surname>Doe</surname>
//   <address id="2" sub="#1">
//      <street>Hello</street>
//	 </address>
// </person>

type Address struct {
	XMLName xml.Name `xml:"address"`
	ID      int      `xml:"id,attr"`
	Sub     string   `xml:"sub,attr"`
	Street  string   `xml:"street"`
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"firstname"`
	Surname   string   `xml:"surname"`
	Address   Address
}
