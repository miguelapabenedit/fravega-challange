package main

import (
	"net/http"

	"github.com/miguelapabenedit/fravega-challange/branch"
	_ "github.com/miguelapabenedit/fravega-challange/docs"
)

const apiBasePath = "/api"

func main() {
	//go:generate swagger generate spec -o ./swagger.yaml --scan-models
	branch.SetUpRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
