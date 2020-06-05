package models

import (
  "time"
)


type Carlog struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Make      string    `json:"make"`
	Model     string    `json:"model"`
	Reg       string    `json:"reg"`
	Year      string    `json:"year"`
	Active    string    `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
