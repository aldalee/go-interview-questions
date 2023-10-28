package user

type god struct {
	dream string
}

func (g god) GetYourDream() string {
	return "Your dream can only be realized by yourself!"
}
