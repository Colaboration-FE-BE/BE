package service

import (
	"immersive-dash-4/features/class"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData class.ClassDataInterface
	validate  *validator.Validate
}

// CreateClass implements class.ClassServiceInterface.
func (*classService) CreateClass(class class.Core) (input class.Core, err error) {
	panic("unimplemented")
}

// GetClassById implements class.ClassServiceInterface.
func (*classService) GetClassById(id uint) (class.Core, error) {
	panic("unimplemented")
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
