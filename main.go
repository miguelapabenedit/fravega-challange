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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/branch/getNearestDeliver", branchController.GetNearestDeliver).Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s%s/{id}", apiBasePath, branchBasePath), branchController.Get).Methods("GET")
	r.HandleFunc(fmt.Sprintf("%s%s", apiBasePath, branchBasePath), branchController.Post).Methods("POST")

	port := ":5000"
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
