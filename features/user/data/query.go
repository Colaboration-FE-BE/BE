package data

import (
	"errors"
	"fmt"
	"immersive-dash-4/app/middlewares"
	"strconv"

	_teamData "immersive-dash-4/features/team"
	_team "immersive-dash-4/features/team/data"
	"immersive-dash-4/features/user"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// SelectAllUser implements user.UserDataInterface.
func (repo *userQuery) SelectAllUser(c echo.Context) ([]user.Core, error) {
	var usersData []User
	var usersCore []user.Core
	var team _teamData.Core
	name := c.QueryParam("fullname")
	page := c.QueryParam("page")
	intPage, _ := strconv.Atoi(page)
	fmt.Println("FULLNAME", page)
	limit := 4
	offset := limit * (intPage - 1)
	if name != "" {
		tx := repo.db.Raw("SELECT*FROM users WHERE fullname LIKE?", "%"+name+"%").Scan(&usersData)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else {
		tx := repo.db.Raw("SELECT*FROM users WHERE is_deleted=0 ORDER BY created_at DESC LIMIT? OFFSET?", limit, offset).Scan(&usersData)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	for _, value := range usersData {
		fmt.Println("ID TEAM ", value.IdTeam)
		txTeam := repo.db.Raw("SELECT name FROM teams WHERE id=? ", value.IdTeam).Scan(&team.Teamname)

		if txTeam.Error != nil {
			return nil, txTeam.Error
		}

		// fmt.Println("Format Time", value.CreatedAt.Format("2006-01-02 3:4:5 pm"))
		// createdAt := value.CreatedAt.Format("2006-01-02 3:4:5 pm")
		// dateCreate, _ := time.Parse("2006-01-02 3:4:5 pm", createdAt)

		var userValue = user.Core{
			ID:          value.ID,
			Fullname:    value.Fullname,
			Email:       value.Email,
			Role:        value.Role,
			PhoneNumber: value.PhoneNumber,
			Team:        team.Teamname,
			Status:      value.Status,
			IsDeleted:   value.IsDeleted,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}

		usersCore = append(usersCore, userValue)
	}
	return usersCore, nil
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (dataLogin user.Core, err error) {
	var data User
	repo.db.Raw("SELECT id,role, email,password FROM users WHERE email=?", email).Scan(&data)
	samePassword := middlewares.CheckPassword(password, data.Password)
	// hashedPassword, err := middlewares.HashedPassword(password)
	// fmt.Println("HASHED PASSWORD", hashedPassword)

	if samePassword {
		tx := repo.db.Raw("SELECT id,role, email,password FROM users WHERE email=? AND password=?", email, data.Password).Scan(&data)
		if tx.Error != nil {
			return user.Core{}, tx.Error
		}
		if tx.RowsAffected == 0 {
			return user.Core{}, errors.New("data not found")
		}
		dataLogin = ModelToCore(data)
	} else {
		return user.Core{}, errors.New("data not found")
	}
	return dataLogin, nil
}

// SelectUserById implements user.UserDataInterface.
func (repo *userQuery) SelectUserById(id string) (user.Core, error) {
	var userData User
	tx := repo.db.Raw("SELECT id,fullname,email,role,status,id_team,created_at,updated_at FROM users WHERE id=?", id).Scan(&userData) // select * from users;
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	fmt.Println("USER DATA", userData.IdTeam)
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}

	//mapping dari struct gorm model ke struct core (entity)
	var usersCore = DetailUser(userData)

	return usersCore, nil
}

// CreateUser implements user.UserDataInterface.
func (repo *userQuery) CreateUser(idUser string, input user.Core) (dataRegister user.Core, err error) {
	var data User

	hashedPassword, _ := middlewares.HashedPassword(input.Password)

	input.ID = idUser
	input.Password = hashedPassword
	input.Role = "DEFAULT"
	input.Status = true

	fmt.Println("STATUS:", input.Status)

	userGorm := CoreToModel(input)
	tx := repo.db.Create(&userGorm)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("failed register new user")
	}
	dataRegister = ModelToCore(data)
	return dataRegister, nil
}

// DeleteUser implements user.UserDataInterface.
func (repo *userQuery) DeleteUser(idUser string) (dataUser user.DeleteCore, err error) {
	var userData User
	var teamData _team.Team
	repo.db.Raw("SELECT*FROM users WHERE id=?", idUser).Scan(&userData)

	tx := repo.db.Exec("update users SET is_deleted=? WHERE id=?", 1, &idUser)

	repo.db.Raw("SELECT*FROM teams WHERE id=?", userData.IdTeam).Scan(&teamData)
	fmt.Println("rsdjds", teamData)
	var userValue = user.DeleteCore{
		ID:        idUser,
		Team:      teamData.Name,
		Fullname:  userData.Fullname,
		Email:     userData.Email,
		Role:      userData.Role,
		IsDeleted: userData.IsDeleted,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
	if tx.Error != nil {
		return userValue, tx.Error
	}
	return userValue, nil
}

// EditUser implements user.UserDataInterface.
func (repo *userQuery) EditUser(c echo.Context, idUser string, input user.Core) (dataUser user.DeleteCore, err error) {
	var userData User
	var teamData _team.Team
	repo.db.Raw("SELECT*FROM users WHERE id=?", idUser).Scan(&userData)
	repo.db.Raw("SELECT*FROM users WHERE id=? AND is_deleted=0", idUser).Scan(&userData)

	fmt.Println("INPUT:::", input)
	if input.Fullname == "" || input.Email == "" || input.Password == "" || input.Role == "" || input.IdTeam == 0 {
		return dataUser, errors.New("validation error. project name description userid required")
	}
	userGorm := CoreToModel(input)
	hashedPassword, _ := middlewares.HashedPassword(userGorm.Password)
	// simpan ke DB
	tx := repo.db.Exec("update users SET fullname=?,email=?,password=?,role=?,id_team=?,status=? WHERE id=?", userGorm.Fullname, userGorm.Email, hashedPassword, userGorm.Role, userGorm.IdTeam, userGorm.Status, &idUser)

	repo.db.Raw("SELECT*FROM teams WHERE id=?", userData.IdTeam).Scan(&teamData)
	fmt.Println("rsdjds", teamData.Name)
	var userValue = user.DeleteCore{
		ID:        idUser,
		Team:      teamData.Name,
		Fullname:  userData.Fullname,
		Email:     userData.Email,
		Role:      userData.Role,
		IsDeleted: userData.IsDeleted,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
	if tx.Error != nil {
		return userValue, tx.Error
	}
	return userValue, nil
}
