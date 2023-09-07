package service

import (
	"immersive-dash-4/features/education"

	"github.com/go-playground/validator/v10"
)

type educationService struct {
	eduactionData education.EducationDataInterface
	validate      *validator.Validate
}

func New(repo education.EducationDataInterface) education.EducationServiceInterface {
	return &educationService{
		eduactionData: repo,
		validate:      validator.New(),
	}
}

// GetAllUser implements user.UserServiceInterface.
func (service *educationService) GetAllClass() ([]education.Core, error) {
	panic("inmkdlksld")
	// return service.eduactionData.
}
