package handler

import (
	"fmt"
	"immersive-dash-4/app/config"
	"immersive-dash-4/app/database"
	_classData "immersive-dash-4/features/class"
	_userData "immersive-dash-4/features/user"
	"immersive-dash-4/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type teamQuery struct {
	db *gorm.DB
}
type ClassHandler struct {
	classService _classData.ClassServiceInterface
}

func New(service _classData.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classService: service,
	}
}

func (handler *ClassHandler) GetAllClass(c echo.Context) error {

	var classResponse []ClassResponse
	// var user _userData.Core
	cfg := config.InitConfig()
	db := database.InitDBMysql(cfg)

	result, err := handler.classService.GetAllClass()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	for _, value := range result {
		fmt.Println("VALUE", value)
		var userData _userData.Core
		db.Raw("SELECT fullname FROM users WHERE id=?", value.PicId).Scan(&userData.Fullname)

		fmt.Println(userData.Fullname)
		classResponse = append(classResponse, ClassResponse{
			ID:           int(value.ID),
			Name:         value.Name,
			Pic:          userData.Fullname,
			StartDate:    value.StartDate,
			GraduateDate: value.GraduateDate,
		})
	}

	// fmt.Println(classResponse[0].Pic)
	// for _, v := range v {

	// }
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success get Class", classResponse))
}
