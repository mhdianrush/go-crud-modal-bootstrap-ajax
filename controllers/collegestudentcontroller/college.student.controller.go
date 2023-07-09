package collegestudentcontroller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
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
	// to get id for edit
	queryString := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(queryString, 10, 64)

	var data map[string]any

	if err != nil {
		data = map[string]any{
			"title": "add college student data",
		}
	} else {
		var collegestudent entities.CollegeStudent

		err = collegestudentmodel.NewCollegeStudentModel().Find(id, &collegestudent)
		if err != nil {
			panic(err)
		}

		data = map[string]any{
			"title":          "edit college student data",
			"collegestudent": collegestudent,
		}
	}

	temp, err := template.ParseFiles("views/collegestudent/form.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		var collegestudent entities.CollegeStudent

		collegestudent.FullName = r.FormValue("full_name")
		collegestudent.Gender = r.FormValue("gender")
		collegestudent.BirthPlace = r.FormValue("birth_place")
		collegestudent.BirthDay = r.FormValue("birth_day")
		collegestudent.Address = r.FormValue("address")

		// to check if there is exist an id in the URL
		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]any

		if err != nil {
			// if id isn't exist ==> Create
			err = collegestudentmodel.NewCollegeStudentModel().Create(&collegestudent)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]any{
				"message": "data has been saved",
				"data":    template.HTML(GetData()),
			}
		} else {
			// if id exist ==> Update
			collegestudent.Id = id
			err = collegestudentmodel.NewCollegeStudentModel().Update(collegestudent)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]any{
				"message": "data has been changed",
				"data":    template.HTML(GetData()),
			}
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

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = collegestudentmodel.NewCollegeStudentModel().Delete(id)
	if err != nil {
		panic(err)
	}
	data := map[string]any{
		"message": "data successfully deleted",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}
