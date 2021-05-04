package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lhsribas/calendar-residency-prediction-core/configuration"
	"github.com/lhsribas/calendar-residency-prediction-core/controller"
	"github.com/lhsribas/calendar-residency-prediction-core/repository"
)

var s = repository.SResidency{}
var c = configuration.Config{}

var port = ":3000"
var version = ""

func init() {
	c.Read()

	// replace the startup port of server
	port = c.Srvport

	// repplace the api versoning
	version = c.Version

	// load database consiguration
	s.Server = c.Server
	s.Database = c.Database

	// Creates the connection with the database
	s.Connect()

	//Create indexes of a collection
	s.CreateResidencyIndex()
}

/*
 * This func up a server in a specific port
 */
func serverUP(router *mux.Router) {
	fmt.Printf("Server Runing in port: %s", port)
	http.ListenAndServe(port, router)
}

/*
 * This func create a router
 */
func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	addRouter(router)
	return router
}

/*
 * This func add new routes to a router
 */
func addRouter(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})

	}).Methods("GET")

	subRouter := router.PathPrefix("/api/v" + version).Subrouter()
	subRouter.HandleFunc("/residency", controller.SaveResidency).Methods("POST")
	subRouter.HandleFunc("/customer", controller.SaveCustomer).Methods("POST")
	subRouter.HandleFunc("/staffMember", controller.SaveStaffMamber).Methods("POST")
	subRouter.HandleFunc("/note", controller.SaveResidencyNotes).Methods("POST")
}

/*
 * This func publish all endpoints
 */
func main() {
	serverUP(newRouter())
}
