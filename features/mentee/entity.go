package mentee

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          string `gorm:"type:varchar(255)" json:"id"`
	Fullname    string `validate:"required"`
	Category    string
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
	SelectAllMentee(c echo.Context) ([]Core, error)
}

type MenteeServiceInterface interface {
	GetAllMentee(c echo.Context) ([]Core, error)
}
