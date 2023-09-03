package user

import "time"

type role string

type Core struct {
	ID          string `gorm:"type:varchar(255)" json:"id"`
	Fullname    string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Role        role   `sql:"type:ENUM('ADMIN', 'DEFAULT')" gorm:"column:role"`
	PhoneNumber string
	IdTeam      uint
	Status      bool
	IsDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserDataInterface interface {
}

type UserServiceInterface interface {
}
