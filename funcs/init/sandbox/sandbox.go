package sandbox

import (
	"log"

	"learn-go/funcs/init/depend"
)

var _ int = generateInt()

func generateInt() int {
	log.Println("Initialising anonymous sandbox int var")
	return 1
}

func init() {
	log.Println("Sandbox init called")
}

func Op() {
	log.Println("Start of sandbox Op func called")
	defer log.Println("End of sandbox Op func called")
	depend.Op()
}
