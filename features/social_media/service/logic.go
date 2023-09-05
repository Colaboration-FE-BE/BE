package service

import (
	"immersive-dash-4/features/class"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData class.ClassDataInterface
	validate  *validator.Validate
}

func New(repo class.ClassDataInterface) class.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}

// GetAllUser implements user.UserServiceInterface.
func (service *classService) GetAllClass() ([]class.Core, error) {
	return service.classData.SelectAllClass()
}
