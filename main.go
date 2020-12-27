package main

import (
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/miguelapabenedit/fravega-challange/branch"
	"github.com/miguelapabenedit/fravega-challange/database"
	_ "github.com/miguelapabenedit/fravega-challange/docs"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	branch.SetUpRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
