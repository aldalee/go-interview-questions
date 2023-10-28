package user

type Result string

type Skill interface {
	Love(person *Person) Result
	Learn()
}
