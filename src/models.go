package main

type Card struct {
	ID          int
	Title       string
	Description string
	Effort      int
	Value       int
	Status      Status
	Iteration   string
}

type Status int

const (
	New        Status = 1
	InProgress Status = 2
	Testin     Status = 3
	Done       Status = 4
)

type Board struct {
	Stories []Card
}
