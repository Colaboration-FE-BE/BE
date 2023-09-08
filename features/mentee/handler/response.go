package handler

import "time"

type MenteeResponse struct {
	ID        string    `json:"id"`
	Fullname  string    `json:"fullname"`
	Class     string    `json:"class"`
	Status    string    `json:"status"`
	Category  string    `json:"category"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
