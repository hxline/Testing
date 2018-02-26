package main

import (
	// model "hxline/model"
	"hxline/model"
	sv "hxline/services"
)

func main() {
	people := model.Person{"3bf9cffc-40a1-4612-8a3c-1f420fea3d91", "Hxasd", 21}
	// sv.Create(people)
	sv.Update(people)
	// sv.Delete("9cbe6341-63f7-4c33-a2a0-168b68c4520f")
}
