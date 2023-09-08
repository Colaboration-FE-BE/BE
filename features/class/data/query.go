package data

import (
	"errors"
	"fmt"
	"immersive-dash-4/features/class"
	"time"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}

// SelectAllClass implements class.ClassDataInterface.
func (repo *classQuery) SelectAllClass() ([]class.Core, error) {
	var classData []Class
	tx := repo.db.Raw("SELECT*FROM classes").Scan(&classData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping dari struct gorm model ke struct core (entity)
	var classCore = ListModelToCore(classData)

	return classCore, nil
}

// InsertClass implements class.ClassDataInterface.
func (repo *classQuery) InsertClass(class class.Core) (input class.Core, err error) {
	var zeroTime time.Time
	if class.Name == "" || class.PicId == "" || class.GraduateDate == zeroTime || class.StartDate == zeroTime {
		return input, errors.New("All field arr required")
	}
	classGorm := CoreToModel(class)
	var ID int64
	tx := repo.db.Exec("INSERT INTO classes(name,pic_id,start_date,graduate_date) VALUES(?,?,?,?) ", classGorm.Name, classGorm.PicId, class.StartDate, class.GraduateDate)
	// tx := repo.db.Create(&classGorm)
	repo.db.Raw("SELECT LAST_INSERT_ID()").Scan(&ID)
	classGorm.ID = uint(ID)
	//mapping dari struct core ke struct gorm model
	var modelClass = ModelToCore(classGorm)
	if tx.Error != nil {
		return modelClass, tx.Error
	}
	return modelClass, nil

}

// SelectClassById implements class.ClassDataInterface.
func (repo *classQuery) SelectClassById(id uint) (class.Core, error) {
	var classData Class

	tx := repo.db.Raw("SELECT*FROM classes WHERE id=?", id).Scan(&classData)

	if tx.Error != nil {
		return class.Core{}, tx.Error
	}

	var class = class.Core{
		ID:           classData.ID,
		Name:         classData.Name,
		PicId:        classData.PicId,
		StartDate:    classData.StartDate,
		GraduateDate: classData.GraduateDate,
	}

	fmt.Println(class)

	return class, nil
}

// DeleteClass implements class.ClassDataInterface.
func (repo *classQuery) DeleteClass(idClass int) error {

	tx := repo.db.Exec("DELETE FROM classes WHERE id=?", idClass)

	if tx.Error != nil {
		return tx.Error
	}

	return nil

}

// EditClass implements class.ClassDataInterface.
func (repo *classQuery) EditClass(c echo.Context, idClass int, input class.Core) (dataClass class.Core, err error) {

	if input.Name == "" || input.PicId == "" {
		return dataClass, errors.New("validation error.all field are required")
	}
	tx := repo.db.Exec("UPDATE classes SET name= ?, pic_id=?,start_date=?,graduate_date=? WHERE id=?", input.Name, input.PicId, input.StartDate, input.GraduateDate, input.ID)
	fmt.Println("TARKWRRKA", tx)
	if tx.Error != nil {
		return dataClass, tx.Error
	}
	return dataClass, nil
}
