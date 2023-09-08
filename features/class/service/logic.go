package service

import (
	"fmt"
	"immersive-dash-4/features/class"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
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

// CreateClass implements class.ClassServiceInterface.
func (service *classService) CreateClass(class class.Core) (input class.Core, err error) {
	fmt.Println("INPUT SERVICE", class)
	result, err := service.classData.InsertClass(class)
	fmt.Println("SERVICE ID SDJDKADK", result.ID)
	return result, err
}

// GetClassById implements class.ClassServiceInterface.
func (service *classService) GetClassById(id uint) (class.Core, error) {
	return service.classData.SelectClassById(id)
}

// DeleteClass implements class.ClassServiceInterface.
func (service *classService) DeleteClass(idClass int) error {
	err := service.classData.DeleteClass(idClass)
	return err
}

// UpdateUser implements class.ClassServiceInterface.
func (service *classService) UpdateClass(c echo.Context, idClass int, input class.Core) (dataUser class.Core, err error) {
	result, err := service.classData.EditClass(c, idClass, input)
	return result, err
}
