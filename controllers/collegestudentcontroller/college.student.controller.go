package collegestudentcontroller

import (
	"bytes"
	"encoding/json"
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

	var collegestudent []entities.CollegeStudent

	err := collegestudentmodel.NewCollegeStudentModel().FindAll(&collegestudent)
	if err != nil {
		panic(err)
	}

	for i := range collegestudent {
		// 20006-01-02 is a default number to construct
		dateTime, _ := time.Parse("2006-01-02T00:00:00Z", collegestudent[i].BirthDay)
		collegestudent[i].BirthDay = dateTime.Format("02-01-2006")
	}

	data := map[string]any{
		"collegestudent": collegestudent,
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

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		var collegestudent entities.CollegeStudent

		collegestudent.FullName = r.Form.Get("full_name")
		collegestudent.Gender = r.Form.Get("gender")
		collegestudent.BirthPlace = r.Form.Get("birth_place")
		collegestudent.BirthDay = r.Form.Get("birth_day")
		collegestudent.Address = r.Form.Get("address")

		err := collegestudentmodel.NewCollegeStudentModel().Create(&collegestudent)
		if err != nil {
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		data := map[string]any{
			"message": "data has been saved",
			"data":    template.HTML(GetData()),
		}
		ResponseJson(w, http.StatusOK, data)
	}
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"Error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
