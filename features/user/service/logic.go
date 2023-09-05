package service

import (
	"immersive-dash-4/app/middlewares"
	"immersive-dash-4/features/user"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
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

// GetAllUser implements user.UserServiceInterface.
func (service *userService) GetAllUser() ([]user.Core, error) {
	return service.userData.SelectAllUser()
}

// DeleteUser implements user.UserServiceInterface.
func (service *userService) DeleteUser(idProject string) (dataUser user.Core, err error) {
	return service.userData.DeleteUser(idProject)
}
