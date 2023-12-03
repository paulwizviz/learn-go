package depend

import "log"

var _ int = generateInt()

func generateInt() int {
	log.Println("Initialising anonymous depend int var")
	return 1
}

func init() {
	log.Println("depend init called")
}

func Op() {
	log.Println("depend Op func called")
}
