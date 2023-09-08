package class

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	Name         string    `validate:"required"`
	PicId        string    `validate:"required"`
	StartDate    time.Time `validate:"required"`
	GraduateDate time.Time
}
type ClassCore struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	PicName      string    `json:"pic_name"`
	StartDate    time.Time `json:"start_date"`
	GraduateDate time.Time `json:"graduate_date"`
}

type ClassDataInterface interface {
	SelectAllClass() ([]Core, error)
	InsertClass(class Core) (input Core, err error)
	SelectClassById(id uint) (Core, error)
	EditClass(c echo.Context, idClass int, input Core) (dataClass Core, err error)
	DeleteClass(idClass int) error
}

type ClassServiceInterface interface {
	GetAllClass() ([]Core, error)
	CreateClass(class Core) (input Core, err error)
	GetClassById(id uint) (Core, error)
	UpdateClass(c echo.Context, idClass int, input Core) (dataUser Core, err error)
	DeleteClass(idClass int) error
}
