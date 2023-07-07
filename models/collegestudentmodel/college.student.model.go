package collegestudentmodel

import (
	"database/sql"

	"github.com/mhdianrush/go-crud-modal-bootstrap-ajax/config"
	"github.com/mhdianrush/go-crud-modal-bootstrap-ajax/entities"
)

type CollegeStudentModel struct {
	DB *sql.DB
}

func NewCollegeStudentModel() *CollegeStudentModel {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	return &CollegeStudentModel{
		DB: db,
	}
}

func (c *CollegeStudentModel) FindAll(CollegeStudent *[]entities.CollegeStudent) error {
	rows, err := c.DB.Query("select * from CollegeStudent")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var CollegeStudentData entities.CollegeStudent
		rows.Scan(
			&CollegeStudentData.Id,
			&CollegeStudentData.FullName,
			&CollegeStudentData.Gender,
			&CollegeStudentData.BirthPlace,
			&CollegeStudentData.BirthDay,
			&CollegeStudentData.Address,
		)
		*CollegeStudent = append(*CollegeStudent, CollegeStudentData)
	}
	return nil
}


