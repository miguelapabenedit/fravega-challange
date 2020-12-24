// Package service thats provides creation and branches near by
//
// Documentatin for Branch Api
//
// Schemes: http
// BasePath: /branch
// version: 1.0.0
// contact: miguell.beneditt@gmail.com
//
// Consumes:
// -application/json
//
// Produces:
// -application/json
// swagger:meta
package branch

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const branchBasePath = "/branch"

/*SetUpRoutes registers the handlers of the service with the given an apiBasePath string*/
func SetUpRoutes(apiBasePath string) {
	handleBranch := http.HandlerFunc(branchHandler)
	handleBranches := http.HandlerFunc(branchesHandler)
	fmt.Print(branchBasePath)
	http.Handle(fmt.Sprintf("%s%s/", apiBasePath, branchBasePath), handleBranch)
	http.Handle(fmt.Sprintf("%s%s", apiBasePath, branchBasePath), handleBranches)

}

// BranchApi serves the Api for the branches
func branchHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "/")
	branchID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])

	if err != nil || branchID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// swagger:route Get /branch/{id}
	//
	// Returns a Branch from the system base on the Id
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// -name: id
	//  description: id of the branch
	//  required:true
	//  type:int
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// swagger:route Post /branch Branch
// Creates a Branch
// responses:
// 200: OKStatus
// 400: BadRequest
func branchesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// A BranchResponse is a response that contains a branch data
// swagger:response branchGetResponse
type BranchResponseWrapper struct {
	// in: body
	Body Branch
}
