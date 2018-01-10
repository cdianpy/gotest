package models

import (
	"log"
	"fmt"
	db "database"
)

type Person struct {
Id 			int 	`json:"id" form:"id"`
FirstName 	string	`json:"first_name" form:"first_name"`
LastName	string	`json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() bool {
	rs,err :=
}
