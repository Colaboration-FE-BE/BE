package service

import (
	"immersive-dash-4/features/education"

	"github.com/go-playground/validator/v10"
)

type educationService struct {
	classData education.EducationDataInterface
	validate  *validator.Validate
}

func New(repo education.EducationDataInterface) education.EducationServiceInterface {
	return &educationService{
		classData: repo,
		validate:  validator.New(),
	}
}

// GetAllUser implements user.UserServiceInterface.
func (service *educationService) GetAllEducation() ([]education.Core, error) {
	panic("Unimplement")
	// return service.classData.SelectAllClass()
}
