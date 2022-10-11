package embed

type SimpleStruct struct {
	ID   int64
	Name string
}

type CompoundStruct struct {
	SimpleStruct
	Name string
}
