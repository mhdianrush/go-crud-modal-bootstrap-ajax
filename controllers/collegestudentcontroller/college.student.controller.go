package collegestudentcontroller

import (
	"bytes"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/collegestudent/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

// we'll use on Index and Delete Data
func GetData() string {
	buffer := &bytes.Buffer{}

	template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/collegestudent/data.html")
}
