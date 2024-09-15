package data

type Snake struct{}

func (a Snake) Eat() string {
	return "mice"
}

func (a Snake) Move() string {
	return "slither"
}

func (a Snake) Speak() string {
	return "hsss"
}
