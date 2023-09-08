package handler

import (
	"fmt"
	"immersive-dash-4/app/config"
	"immersive-dash-4/app/database"
	_classData "immersive-dash-4/features/class"
	_userData "immersive-dash-4/features/user"
	"immersive-dash-4/helpers"
	"net/http"
	"strconv"
	"strings"

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

func (handler *ClassHandler) CreateClass(c echo.Context) error {
	classInput := new(ClassRequest)

	errBind := c.Bind(&classInput)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	classCore := RequestToCore(*classInput)
	fmt.Println("RESULT HANDLER", classCore)
	result, err := handler.classService.CreateClass(classCore)
	fmt.Println("RESULT HANDLER====", result)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))

		}
	}
	resultResponse := ClassResponse{
		ID:           int(result.ID),
		Name:         result.Name,
		Pic:          result.PicId,
		StartDate:    result.StartDate,
		GraduateDate: result.GraduateDate,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success register new class", resultResponse))
}

func (handler *ClassHandler) GetClassById(c echo.Context) error {
	id := c.Param("class_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}

	result, err := handler.classService.GetClassById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}

	resultResponse := ClassResponse{
		ID:           idConv,
		Name:         result.Name,
		Pic:          result.PicId,
		StartDate:    result.StartDate,
		GraduateDate: result.GraduateDate,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read class", resultResponse))
}

func (handler *ClassHandler) DeleteClass(c echo.Context) error {
	id := c.Param("class_id")

	idClass, errConv := strconv.Atoi(id)
	fmt.Println("ID CLASS", idClass)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error. id should be number",
		})
	}

	err := handler.classService.DeleteClass(idClass)
	fmt.Println("DEBUG", err)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success Delete Class", nil))
}
