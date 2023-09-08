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
	InsertClass(class Core) (input Core, err error)
	SelectClassById(id uint) (Core, error)
	DeleteClass(idClass int) error
}

type ClassServiceInterface interface {
	GetAllClass() ([]Core, error)
	CreateClass(class Core) (input Core, err error)
	GetClassById(id uint) (Core, error)
	DeleteClass(idClass int) error
}
