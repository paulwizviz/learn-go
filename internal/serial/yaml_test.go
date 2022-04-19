package serial

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Models
type SimpleStruct struct {
	Field1 string
	Field2 string
}

type StructWithTag struct {
	Field1 string `yaml:"Field1,omitempty"`
	Field2 string `yaml:"Field2"`
}

type EmbeddedStruct struct {
	SimpleStruct
}

type LongDescription struct {
	Description string `yaml:"Description"`
}

type ComplexStruct struct {
	Project     string         `yaml:"Project"`
	Description string         `yaml:"Description"`
	Tags        []string       `yaml:"Tags"`
	Subprojects []SimpleStruct `yaml:"Subprojects"`
}

// Test cases for marshalling
func Example_marshalSimpleStruct() {
	s := SimpleStruct{}
	output, err := yaml.Marshal(s)
	if err != nil {
		fmt.Println("This should not happens")
	}
	fmt.Println(string(output))
	// Output:
	// field1: ""
	// field2: ""
}

func Example_marshalTagStruct() {
	d := StructWithTag{
		Field2: "field 2",
	}
	output, err := yaml.Marshal(d)
	if err != nil {
		fmt.Println("This should not happens")
	}
	fmt.Println(string(output))
	// Output:
	// Field2: field 2
}

func Example_marshalEmbeddedStruct() {
	d := EmbeddedStruct{
		SimpleStruct{
			Field1: "field1",
		},
	}
	output, err := yaml.Marshal(d)
	if err != nil {
		fmt.Println("This should not happens")
	}
	fmt.Println(string(output))
	// Output:
	// simplestruct:
	//   field1: field1
	//   field2: ""
}

func Example_marshalLongDescription() {
	d := LongDescription{
		Description: `This is a long description
with multiple lines`,
	}
	output, err := yaml.Marshal(d)
	if err != nil {
		fmt.Println("This should not happens")
	}
	fmt.Println(string(output))
	// Output:
	// Description: |-
	//   This is a long description
	//   with multiple lines
}

func Example_marshalMappingSequence() {
	d := []map[string]string{
		{
			"Field1": "field1.1",
			"Field2": "field1.2",
		},
		{
			"Field1": "field2.1",
			"Field2": "field2.2",
		},
	}
	output, err := yaml.Marshal(d)
	if err != nil {
		fmt.Println("This should not happens")
	}
	fmt.Println(string(output))
	// Output:
	// - Field1: field1.1
	//   Field2: field1.2
	// - Field1: field2.1
	//   Field2: field2.2
}

func Example_marshalComplexStruct() {
	d := ComplexStruct{
		Project: "Alpha",
		Description: `This is a project to send
capsule Alpha to the moon`,
		Tags: []string{
			"Space",
			"Capsule",
		},
		Subprojects: []SimpleStruct{
			{
				Field1: "field1",
				Field2: "field2",
			},
			{
				Field1: "field1.1",
				Field2: "field2.1",
			},
		},
	}

	output, err := yaml.Marshal(d)
	if err != nil {
		fmt.Println("This should not happens")
	}
	fmt.Println(string(output))
	// Output:
	// Project: Alpha
	// Description: |-
	//   This is a project to send
	//   capsule Alpha to the moon
	// Tags:
	// - Space
	// - Capsule
	// Subprojects:
	// - field1: field1
	//   field2: field2
	// - field1: field1.1
	//   field2: field2.1
}

// Test cases for unmarshalling
func Example_unmarshalNumeric() {
	b := []byte(`
field1: 1
field2: 1.0
`)
	var d map[string]interface{}
	err := yaml.Unmarshal(b, &d)
	if err != nil {
		fmt.Println("This is not be expected")
	}

	for _, v := range d {
		switch vv := v.(type) {
		case int:
			fmt.Printf("Type: int Value: %v\n", vv)
		case float32, float64:
			fmt.Printf("Type: float Value: %.2f\n", vv)
		}
	}

	// Output:
	// Type: int Value: 1
	// Type: float Value: 1.00
}

func Example_unmarshalSimpleStruct() {
	b := []byte(`
field1: field1
field2: field2	
`)
	var d SimpleStruct
	err := yaml.Unmarshal(b, &d)
	if err != nil {
		fmt.Println("This is not be expected")
	}
	fmt.Println(d)
	// Output:
	// {field1 field2}
}

func Example_unmarshalLongDescriptionLineBreak() {
	b := []byte(`
Description: |
  This is a long description
  with multiline.
`)
	var d LongDescription
	err := yaml.Unmarshal(b, &d)
	if err != nil {
		fmt.Println("This is not be expected")
	}
	fmt.Println(d)
	// Output:
	// {This is a long description
	// with multiline.
	// }

}

func Example_unmarshalLongDescriptionSpanning() {
	b := []byte(`
Description: >
  This is a long description
  with multiline.
`)
	var d LongDescription
	err := yaml.Unmarshal(b, &d)
	if err != nil {
		fmt.Println("This is not be expected")
	}
	fmt.Println(d)
	// Output:
	// {This is a long description with multiline.
	// }

}

func Example_unmarshalMappingSequence() {
	b := []byte(`
- field1: field1.1
  field2: field1.2
- field1: field2.1
  field2: field2.2`)

	var d interface{}
	err := yaml.Unmarshal(b, &d)
	if err != nil {
		fmt.Println("This is not be expected")
	}
	fmt.Println(d)
	// Output:
	// [map[field1:field1.1 field2:field1.2] map[field1:field2.1 field2:field2.2]]
}

func Example_unmarshalMappingSequenceSimpleStruct() {
	b := []byte(`
- field1: field1.1
  field2: field1.2
- field1: field2.1
  field2: field2.2`)

	var d []SimpleStruct
	err := yaml.Unmarshal(b, &d)
	if err != nil {
		fmt.Println("This is not be expected")
	}
	for _, item := range d {
		fmt.Printf("Type: %T Value: %v\n", item, item)
	}

	// Output:
	// Type: goyaml.SimpleStruct Value: {field1.1 field1.2}
	// Type: goyaml.SimpleStruct Value: {field2.1 field2.2}
}
