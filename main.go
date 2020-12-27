package main

import (
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/miguelapabenedit/fravega-challange/controller"
	_ "github.com/miguelapabenedit/fravega-challange/docs"
	"github.com/miguelapabenedit/fravega-challange/infrastructure"
	"github.com/miguelapabenedit/fravega-challange/service"
)

const apiBasePath = "/api"

var branchRepo infrastructure.Repository = infrastructure.NewSQLRepository()
var branchservice service.Service = service.NewBranchService(branchRepo)

func main() {
	controller.SetUpRoutes(apiBasePath, branchservice)
	http.ListenAndServe(":5000", nil)
}
