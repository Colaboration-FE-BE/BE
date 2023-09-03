package data

import (
	"errors"
	"fmt"
	"immersive-dash-4/app/middlewares"
	"immersive-dash-4/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
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
	tx := repo.db.Raw("SELECT id,fullname,email,role,status,id_team FROM users WHERE id=?", id).Scan(&userData) // select * from users;
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
