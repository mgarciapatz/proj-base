package routers

import (
	"proj-base/controllers"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

func SampleRoutes(router *mux.Router) *mux.Router {
	//First Approach
	router.HandleFunc("/"+app_version+"/sample-routes",controllers.SampleController).Methods("POST")
	router.HandleFunc("/"+app_version+"/sample-insert-routes",controllers.SampleInsertController).Methods("POST")
	router.HandleFunc("/"+app_version+"/sample-update-routes",controllers.SampleUpdateController).Methods("POST")
	router.HandleFunc("/"+app_version+"/sample-delete-routes",controllers.SampleDeleteController).Methods("POST")
	router.HandleFunc("/"+app_version+"/sample-delete-routes",controllers.SampleDeleteController).Methods("GET")
	router.HandleFunc("/"+app_version+"/sample-read-routes",controllers.SampleReadController).Methods("GET")

	//Second Approach with JWT Authentication
	router.Handle("/"+app_version+"/sample/{action}",
		negroni.New(
			negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
			negroni.HandlerFunc(controllers.AllInOneCrudController),
		)).
	Methods("POST")

	//JWT Sample
	router.HandleFunc("/"+app_version+"/tokenizer",controllers.JWTokenizer).Methods("GET")
	router.Handle("/"+app_version+"/protected",
		negroni.New(
			negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
			negroni.HandlerFunc(controllers.ProtectedAction),
		)).
	Methods("GET")
	return router
}
