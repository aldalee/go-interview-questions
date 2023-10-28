package user

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func (p *Person) GetAge() int {
	return p.Age
}
