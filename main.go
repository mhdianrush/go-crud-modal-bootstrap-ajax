package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/mhdianrush/go-crud-modal-bootstrap-ajax/config"
	"github.com/mhdianrush/go-crud-modal-bootstrap-ajax/controllers/collegestudentcontroller"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/collegestudent", collegestudentcontroller.Index)
	r.HandleFunc("/collegestudent/get_form", collegestudentcontroller.GetForm)
	r.HandleFunc("/collegestudent/store", collegestudentcontroller.Store)
	r.HandleFunc("/collegestudent/delete", collegestudentcontroller.Delete)

	logger := logrus.New()
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.Println("Server Running on Port 8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
