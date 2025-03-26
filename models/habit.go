package models

type Habit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}
