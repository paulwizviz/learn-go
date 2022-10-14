package embed

type SimpleStruct struct {
	ID   int64
	Name string
}

func (s SimpleStruct) StructName() string {
	return s.Name
}

type CompoundStruct struct {
	SimpleStruct
	Name string
}
