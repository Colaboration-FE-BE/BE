package handler

import (
	"immersive-dash-4/features/mentee"
	"time"
)

type MenteeRequest struct {
	Fullname string `json:"fullname"`
	Class    int    `json:"class"`
	Status   string `json:"status"`
	Category string `json:"category"`
	Gender   int    `json:"gender"`
}

func RequestToCore(input MenteeRequest) mentee.Core {
	return mentee.Core{
		ID:          "",
		Fullname:    "",
		IdClass:     0,
		Class:       "",
		Status:      "",
		Address:     "",
		HomeAddress: "",
		Gender:      false,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
}
