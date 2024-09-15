package data

type Cow struct{}

func (a Cow) Eat() string {
	return "grass"
}

func (a Cow) Move() string {
	return "walk"
}

func (a Cow) Speak() string {
	return "moo"
}
