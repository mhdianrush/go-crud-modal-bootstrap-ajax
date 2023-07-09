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

func (c *CollegeStudentModel) Find(id int64, collegestudent *entities.CollegeStudent) error {
	row := c.DB.QueryRow("select * from collegestudent where id = ?", id)
	return row.Scan(
		&collegestudent.Id,
		&collegestudent.FullName,
		&collegestudent.Gender,
		&collegestudent.BirthPlace,
		&collegestudent.BirthDay,
		&collegestudent.Address,
	)
}

func (c *CollegeStudentModel) Update(collegestudent entities.CollegeStudent) error {
	_, err := c.DB.Exec(
		`update collegestudent set
		full_name = ?,
		gender = ?,
		birth_place = ?,
		birth_day = ?,
		address = ?
		where id = ?`,
		collegestudent.FullName,
		collegestudent.Gender,
		collegestudent.BirthPlace,
		collegestudent.BirthDay,
		collegestudent.Address,
		collegestudent.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *CollegeStudentModel) Delete(id int64) error {
	_, err := c.DB.Exec(`delete from collegestudent where id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}
