package model

type Card struct {
	Id          int
	Name        string
	Description string
	Exercises   []Exercises
}

type Exercises struct {
	Statement string
	Answer    string
	Correct   bool
}
