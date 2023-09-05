package user

import (
	"time"
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

type UserDataInterface interface {
	Login(email string, password string) (dataLogin Core, err error)
	SelectUserById(id string) (dataUser Core, err error)
	CreateUser(idUser string, input Core) (dataRegister Core, err error)
	SelectAllUser() ([]Core, error)
	DeleteUser(idProject string) (dataUser Core, err error)
}

type UserServiceInterface interface {
	Login(email string, password string) (dataLogin Core, token string, err error)
	GetUserById(id string) (dataUser Core, err error)
	CreateUser(idUser string, input Core) (dataRegister Core, err error)
	GetAllUser() ([]Core, error)
	DeleteUser(idProject string) (dataUser Core, err error)
}
