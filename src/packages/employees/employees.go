package employees

import "fmt"

type Person struct {
	Name       string
	persId int
}

type employee struct {
	Person
	empId int
}

func GetEmployee() employee {
	return employee{}
}

func (p *Person) SetPersonId(id int) {
	p.persId = id 
}

func (e *employee) SetEmployeeId(id int) {
	e.empId = id
}

func (p Person) Print () {
	fmt.Println("type = Person")
	fmt.Printf("Id: %d - Name: %s\n",p.persId, p.Name)
}

func (e employee) Print () {
	fmt.Println("type = Employee")
	fmt.Printf("EId: %d - PId: %d - Name: %s\n",e.empId, e.persId, e.Name)
	
}