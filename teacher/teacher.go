package teacher

import "fmt"

type Teacher struct {
	Id         int
	Name       string
	Subject    string
	BaseSalary float64
	Bonus      float64
}

func (t Teacher) GetInfo() {
	fmt.Printf("Id : %d | Name : %s | Subject : %s | ToTal Salary : %.2f \n", t.Id, t.Name, t.Subject, t.TotalSalary())
}

func (t Teacher) TotalSalary() float64 {
	return t.BaseSalary + t.Bonus
}

func (t Teacher) GetId() int {
	return t.Id
}
