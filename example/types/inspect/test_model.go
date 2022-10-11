package inspect

import (
	"fmt"
	"reflect"
)

type FieldElement struct {
	Type string
	Kind string
}

type StructMeta interface {
	Name() string
	Fields() []reflect.StructField
	AppendField(field reflect.StructField)
}

type structMeta struct {
	name   string
	fields []reflect.StructField
}

func (s *structMeta) Name() string {
	return s.name
}

func (s *structMeta) Fields() []reflect.StructField {
	return s.fields
}

func (s *structMeta) AppendField(field reflect.StructField) {
	s.fields = append(s.fields, field)
}

func NewStructMeta(structName string) StructMeta {
	return &structMeta{
		name: structName,
	}
}

func ExtractStructMeta(data interface{}) (StructMeta, error) {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not struct")
	}

	strutName := value.Type().String()
	sm := NewStructMeta(strutName)
	for index := 0; index < value.NumField(); index++ {
		sm.AppendField(value.Type().Field(index))
	}

	return sm, nil
}

func ExtractFieldNames(data interface{}) map[string]FieldElement {

	fieldElms := make(map[string]FieldElement)

	elm := reflect.Indirect(reflect.ValueOf(data))

	for index := 0; index < elm.NumField(); index++ {
		fn := elm.Type().Field(index).Name
		ft := elm.Type().Field(index).Type

		if ft.Kind() == reflect.Struct {
			fieldElms[fn] = FieldElement{
				Type: fmt.Sprintf("%v", ft),
				Kind: fmt.Sprintf("%v", ft.Kind()),
			}
			fd := elm.Type().Field(index)
			for i := 0; i < fd.Type.NumField(); i++ {
				sfn := fd.Type.Field(i).Name
				sft := fd.Type.Field(i).Type
				fieldElms[sfn] = FieldElement{
					Type: fmt.Sprintf("%v", sft),
					Kind: fmt.Sprintf("%v", sft.Kind()),
				}
			}
		} else {
			fieldElms[fn] = FieldElement{
				Type: fmt.Sprintf("%v", ft),
				Kind: fmt.Sprintf("%v", ft.Kind()),
			}
		}
	}
	return fieldElms
}
