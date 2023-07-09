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

func (c *CollegeStudentModel) FindAll(collegestudent *[]entities.CollegeStudent) error {
	rows, err := c.DB.Query("select * from collegestudent")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var collegestudentdata entities.CollegeStudent
		rows.Scan(
			&collegestudentdata.Id,
			&collegestudentdata.FullName,
			&collegestudentdata.Gender,
			&collegestudentdata.BirthPlace,
			&collegestudentdata.BirthDay,
			&collegestudentdata.Address,
		)
		*collegestudent = append(*collegestudent, collegestudentdata)
	}
	return nil
}

func (c *CollegeStudentModel) Create(collegestudent *entities.CollegeStudent) error {
	result, err := c.DB.Exec(
		"insert into collegestudent(full_name, gender, birth_place, birth_day, address) values(?,?,?,?,?)",
		collegestudent.FullName,
		collegestudent.Gender,
		collegestudent.BirthPlace,
		collegestudent.BirthDay,
		collegestudent.Address,
	)
	if err != nil {
		return err
	}
	lastInsertId, _ := result.LastInsertId()
	collegestudent.Id = lastInsertId
	return nil
}
