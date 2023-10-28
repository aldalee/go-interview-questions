package main

import (
	"fmt"
	"go-interview-questions/learning-go/oop/io"
	"go-interview-questions/learning-go/oop/user"
)

func main() {
	james := &user.Person{FirstName: "James", LastName: "Gosling", Age: 68}
	rob := user.NewStudent(101, "Rob", "Pike", 40)
	russ := user.NewStudent(102, "Russ", "Cox", 30)

	fmt.Println(rob.Love(james))
	fmt.Println(rob.Love(&russ.Person))

	io.ReadFile(&io.Reader{})
	io.ReadFile(&io.Channel{})
}
