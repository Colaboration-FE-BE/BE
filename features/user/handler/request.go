package handler

import "immersive-dash-4/features/user"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Fullname    string `json:"full_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Status      bool   `json:"status"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	IdTeam      int    `json:"id_team"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		Fullname:    input.Fullname,
		Role:        input.Role,
		Email:       input.Email,
		Status:      input.Status,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		IdTeam:      input.IdTeam,
	}
}
