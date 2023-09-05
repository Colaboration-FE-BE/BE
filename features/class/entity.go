package class

import (
	"time"
)

type Core struct {
	ID           uint
	Name         string    `validate:"required"`
	PicId        string    `validate:"required"`
	StartDate    time.Time `validate:"required"`
	GraduateDate time.Time
}

type ClassDataInterface interface {
	SelectAllClass() ([]Core, error)
}

type ClassServiceInterface interface {
	GetAllClass() ([]Core, error)
}
