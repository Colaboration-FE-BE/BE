package team

type Core struct {
	ID       string `gorm:"type:varchar(255)" json:"id"`
	Teamname string `validate:"required"`
}

// type UserDataInterface interface {
// 	Login(email string, password string) (dataLogin Core, err error)
// 	SelectUserById(id string) (dataUser Core, err error)
// }

// type UserServiceInterface interface {
// 	Login(email string, password string) (dataLogin Core, token string, err error)
// 	GetUserById(id string) (dataUser Core, err error)
// }
