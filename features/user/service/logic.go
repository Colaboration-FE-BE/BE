package service

import (
	"immersive-dash-4/app/middlewares"
	"immersive-dash-4/features/user"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// GetAllUser implements user.UserServiceInterface.
func (service *userService) GetAllUser(c echo.Context) ([]user.Core, error) {
	return service.userData.SelectAllUser(c)
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (dataLogin user.Core, token string, err error) {
	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return user.Core{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID)
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

// GetUserById implements user.UserServiceInterface.
func (service *userService) GetUserById(id string) (user.Core, error) {
	result, err := service.userData.SelectUserById(id)
	return result, err
}

// CreateUser implements user.UserServiceInterface.
func (service *userService) CreateUser(idUser string, input user.Core) (dataRegister user.Core, err error) {
	dataRegister, err = service.userData.CreateUser(idUser, input)
	if err != nil {
		return user.Core{}, err
	}
	//  err = middlewares.CreateToken(int(dataRegister.ID))
	// if err != nil {
	// 	return user.Core{},  err
	// }
	return dataRegister, nil
}

// DeleteUser implements user.UserServiceInterface.
func (service *userService) DeleteUser(idUser string) (dataUser user.DeleteCore, err error) {
	result, err := service.userData.DeleteUser(idUser)
	return result, err
}

// UpdateUser implements user.UserServiceInterface.
func (service *userService) UpdateUser(c echo.Context, idUser string, input user.Core) (dataUser user.DeleteCore, err error) {
	result, err := service.userData.EditUser(c, idUser, input)
	return result, err
}
