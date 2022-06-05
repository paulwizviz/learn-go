package override

import "fmt"

type data struct {
	ID int
}

func (d *data) changeID(newID int) {
	d.ID = newID
}

func Example_data() {
	d := &data{}
	d.changeID(20)
	fmt.Println(d.ID)
	// Output:
	// 20
}

type dataWrapper struct {
	data
}

func Example_dataWrapper() {
	w := dataWrapper{
		data{},
	}
	w.changeID(20)
	fmt.Println(w.ID)
	// Output:
	// 20
}

type dataOverride struct {
	data
}

func (o *dataOverride) changeID(newID int) {
	o.data.ID = newID * 10
}

func Example_dataOverride() {
	o := dataOverride{
		data{},
	}
	o.changeID(10)
	fmt.Println(o.ID)
	// Output:
	// 100
}
