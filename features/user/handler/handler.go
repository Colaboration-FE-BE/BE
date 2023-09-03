package handler

import (
	"fmt"
	"immersive-dash-4/app/config"
	"immersive-dash-4/app/database"
	_teamData "immersive-dash-4/features/team"
	_userData "immersive-dash-4/features/user"
	"immersive-dash-4/helpers"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type teamQuery struct {
	db *gorm.DB
}
type UserHandler struct {
	userService _userData.UserServiceInterface
}

func New(service _userData.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	response := map[string]any{
		"id":    dataLogin.ID,
		"email": dataLogin.Email,
		"token": token,
		"role":  dataLogin.Role,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("user_id")
	fmt.Println("ID PARAMS", id)

	result, err := handler.userService.GetUserById(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusNotFound, "User Not found", nil))

		}
	}
	cfg := config.InitConfig()
	db := database.InitDBMysql(cfg)
	var teamData _teamData.Core
	db.Raw("SELECT name FROM teams WHERE id=?", result.IdTeam).Scan(&teamData.Teamname)

	fmt.Println(teamData.Teamname)

	resultResponse := UserResponse{
		ID:        result.ID,
		Fullname:  result.Fullname,
		Email:     result.Email,
		Team:      teamData.Teamname,
		Role:      result.Role,
		Status:    result.Status,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success get specific user data", resultResponse))
}

func (handler *UserHandler) CreateNewUser(c echo.Context) error {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	fmt.Println(userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	userCore := RequestToCore(*userInput)

	id := uuid.New()
	idUser := id.String()
	dataRegister, err := handler.userService.CreateUser(idUser, userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error register", nil))

		}
	}
	fmt.Print("DATA REGISTER", dataRegister)
	// fmt.Print(token)
	response := map[string]any{
		"user_id":   idUser,
		"full_name": userInput.Fullname,
		"password":  userInput.Password,
		"role":      "DEFAULT",
		"id_team":   userInput.IdTeam,
		// "user_name": dataRegister.Username,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success register", response))
}
