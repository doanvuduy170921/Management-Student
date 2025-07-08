package student

import "fmt"

type Student struct {
	Id     int
	Name   string
	Class  string
	Scores []float64
}

func (s Student) GetInfo() {
	fmt.Printf("Id : %d | Name : %s | Class : %s | Avg point : %.2f \n", s.Id, s.Name, s.Class, s.avgStudentScore())
}

func (s Student) avgStudentScore() float64 {
	total := 0.0
	for _, score := range s.Scores {
		total += score
	}
	return total / float64(len(s.Scores))
}

func (s Student) GetId() int {
	return s.Id
}
