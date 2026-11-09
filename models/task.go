package models

// Task represents a student's academic task.
type Task struct {
	ID                 int
	Title              string
	Description        string
	Subject            string
	TaskType           string
	DueDate            string
	Priority           string
	CalculatedPriority string
	Status             string
}
