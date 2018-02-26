package main

import (
	"fmt"
	"hxline/model"
	sv "hxline/services"
)

func main() {
	people := model.Person{"7a2dd175-7360-4c44-9559-42f7c4cae251", "GuRenrd", 31}
	sv.Create(people)
	// sv.Update(people)
	sv.Delete("b6974e6a-7add-49f0-bf9c-9cce77a296cb")

	rows := sv.ReadAllPeople()
	for rows.Next() {
		var id string
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		fmt.Println(id, " | ", name, " | ", age)
	}
}
