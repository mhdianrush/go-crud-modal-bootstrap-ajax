package collegestudentcontroller

import (
	"bytes"
	"html/template"
	"net/http"
	"time"

	"github.com/mhdianrush/go-crud-modal-bootstrap-ajax/entities"
	"github.com/mhdianrush/go-crud-modal-bootstrap-ajax/models/collegestudentmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"data": template.HTML(GetData()),
	}
	temp, err := template.ParseFiles("views/collegestudent/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

// we'll use on Index and Delete Data
func GetData() string {
	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/collegestudent/data.html")

	var CollegeStudent []entities.CollegeStudent

	err := collegestudentmodel.NewCollegeStudentModel().FindAll(&CollegeStudent)
	if err != nil {
		panic(err)
	}

	for i := range CollegeStudent {
		// 20006-01-02 is a default number to construct
		dateTime, _ := time.Parse("2006-01-02T00:00:00Z", CollegeStudent[i].BirthDay)
		CollegeStudent[i].BirthDay = dateTime.Format("02-01-2006")
	}

	data := map[string]any{
		"collegestudent": CollegeStudent,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)
	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/collegestudent/form.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
