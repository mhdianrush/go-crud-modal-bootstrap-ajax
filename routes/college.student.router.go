package routes

import (
	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-crud-modal-bootstrap-ajax/controllers/collegestudentcontroller"
)

func NewRouter(r *mux.Router) {
	router := r.PathPrefix("/student").Subrouter()

	router.HandleFunc("", collegestudentcontroller.Index)
	router.HandleFunc("/collegestudent/get_form", collegestudentcontroller.GetForm)
	
}
