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

func branchHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "/")
	branchID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])

	if err != nil || branchID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
