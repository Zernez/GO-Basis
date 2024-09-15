package data

type Bird struct{}

func (a Bird) Eat() string {
	return "worms"
}

func (a Bird) Move() string {
	return "fly"
}

func (a Bird) Speak() string {
	return "peep"
}
