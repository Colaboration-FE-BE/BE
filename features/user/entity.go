package user

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          string `gorm:"type:varchar(255)" json:"id"`
	Fullname    string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Role        string `sql:"type:ENUM('ADMIN', 'DEFAULT')" gorm:"column:role"`
	PhoneNumber string
	IdTeam      int
	Team        string
	Status      bool
	IsDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DeleteCore struct {
	ID        string `json:"id"`
	Team      string `json:"team"`
	Fullname  string `json:"full_name"`
	Email     string `validate:"required,email" json:"email"`
	Role      string `validate:"required,role" json:"role"`
	IsDeleted bool   `validate:"required,is_deleted" json:"is_deleted"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface {
	Login(email string, password string) (dataLogin Core, err error)
	SelectUserById(id string) (dataUser Core, err error)
	CreateUser(idUser string, input Core) (dataRegister Core, err error)
	SelectAllUser(c echo.Context) ([]Core, error)
	DeleteUser(idProject string) (dataUser DeleteCore, err error)
	EditUser(c echo.Context, idUser string, input Core) (dataUser DeleteCore, err error)
}

type UserServiceInterface interface {
	Login(email string, password string) (dataLogin Core, token string, err error)
	GetUserById(id string) (dataUser Core, err error)
	CreateUser(idUser string, input Core) (dataRegister Core, err error)
	GetAllUser(c echo.Context) ([]Core, error)
	DeleteUser(idProject string) (dataUser DeleteCore, err error)
	UpdateUser(c echo.Context, idUser string, input Core) (dataUser DeleteCore, err error)
}
