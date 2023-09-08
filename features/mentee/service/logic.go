package service

import (
	"immersive-dash-4/features/mentee"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type menteeService struct {
	menteeData mentee.MenteeDataInterface
	validate   *validator.Validate
}

// GetAllMentee implements mentee.MenteeServiceInterface.
func (service *menteeService) GetAllMentee(c echo.Context) ([]mentee.Core, error) {
	return service.menteeData.SelectAllMentee(c)
}

func New(repo mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
	}
}
