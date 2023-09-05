package mentee

import (
	"time"
)

type Core struct {
	ID          string `gorm:"type:varchar(255)" json:"id"`
	Fullname    string `validate:"required"`
	IdClass     int
	Class       string
	Status      string `sql:"type:ENUM('PLACEMENT', 'ACTIVE','ELIMINATE')" gorm:"column:status"`
	Address     string
	HomeAddress string
	Gender      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type MenteeDataInterface interface {
	SelectAllMentee() ([]Core, error)
}

type MenteeServiceInterface interface {
	GetAllMentee() ([]Core, error)
}
