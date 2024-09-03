package todo

import "time"

type Todo struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}

type TodoUpdate struct {
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
	Title     *string   `json:"title"`
	Completed *bool     `json:"completed"`
}
