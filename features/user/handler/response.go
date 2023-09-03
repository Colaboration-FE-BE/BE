package handler

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	Team      string    `json:"team"`
	Role      string    `json:"role"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"uptadeted_at"`
}
