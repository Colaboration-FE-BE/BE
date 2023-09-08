package handler

import (
	"fmt"
	_menteeData "immersive-dash-4/features/mentee"
	"immersive-dash-4/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

type MenteeHandler struct {
	menteeService _menteeData.MenteeServiceInterface
}

func New(service _menteeData.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		menteeService: service,
	}
}

func (handler *MenteeHandler) GetAllMentee(c echo.Context) error {
	result, err := handler.menteeService.GetAllMentee(c)
	fmt.Println("RESULT MENTEE", result)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}

	var menteeResponse []MenteeResponse
	for _, value := range result {
		fmt.Println("CLASS ID ", value.Category)
		var gender string
		if value.Gender == false {
			gender = "Male"
		} else {
			gender = "Female"
		}
		menteeResponse = append(menteeResponse, MenteeResponse{
			ID:        value.ID,
			Fullname:  value.Fullname,
			Class:     value.Class,
			Status:    value.Status,
			Category:  value.Category,
			Gender:    gender,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read datahandler", menteeResponse))
}
