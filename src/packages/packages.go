package main

import (
	//	"fmt"
	"packages/employees"
)

func main() {
	e1 := new(employees.Person)
	e1.Name = "Frank"
	e1.SetPersonId(32)
	e1.Print()
	e2 := employees.GetEmployee()
	e2.Name = "Ralph"
	e2.SetPersonId(5)
	e2.SetEmployeeId(3)
	e2.Print()
	e3 := employees.GetEmployee()
	e3.Name = "Lucy"
	e3.SetEmployeeId(120)
	e3.SetPersonId(500)
	e3.Person.Print()
	e3.Print()
	e2.Print()

}
