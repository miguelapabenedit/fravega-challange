package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"github.com/miguelapabenedit/fravega-challange/controller"
	_ "github.com/miguelapabenedit/fravega-challange/docs"
	"github.com/miguelapabenedit/fravega-challange/infrastructure"
	"github.com/miguelapabenedit/fravega-challange/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	apiBasePath    = "/api"
	branchBasePath = "/branch"
)

var (
	branchRepo       infrastructure.Repository = infrastructure.NewSQLRepository()
	branchservice    service.Service           = service.NewBranchService(branchRepo)
	branchController controller.Controller     = controller.NewBranchController(branchservice)
)

// @title Fravega Challange Deliver API
// @version 1.0
// @description This api serves for creation and retrieve of branches and other functionalities
// @contact.name API Support
// @contact.email miguell.beneditt@gmail.com
// @host localhost:5000
// @BasePath /api
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/branch/getNearestDeliver", branchController.GetNearestDeliver).Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s%s/{id}", apiBasePath, branchBasePath), branchController.Get).Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s%s", apiBasePath, branchBasePath), branchController.Post).Methods("POST")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	port := ":5000"
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
