package user

import "fmt"

type Student struct {
	id int
	Person
	god
}

func NewStudent(id int, firstName, lastName string, age int) *Student {
	return &Student{
		id:     id,
		Person: Person{firstName, lastName, age},
	}
}

func (s *Student) GetAge() int {
	return 18
}

func (s *Student) Love(p *Person) Result {
	return Result(fmt.Sprintf("%s and %s are in love.", s.FirstName, p.FirstName))
}

func (s *Student) Learn() {
	fmt.Println("Students Coding.")
}
