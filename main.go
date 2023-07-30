package main

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

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

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r,
	}

	if err = server.ListenAndServe(); err != nil {
		logger.Printf("failed connect to server %s", err.Error())
	}

	logger.Printf("Server Running on Port %s", os.Getenv("SERVER_PORT"))
}
