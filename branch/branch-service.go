package branch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	branch, err := getBranch(branchID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if branch == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		branchJSON, err := json.Marshal(&branch)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(branchJSON)
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
	case http.MethodGet:
		latitudeForm := r.FormValue("latitude")
		longitudeForm := r.FormValue("longitude")

		if latitudeForm == "" || longitudeForm == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		latitude, err := strconv.ParseFloat(latitudeForm, 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		longitude, err := strconv.ParseFloat(longitudeForm, 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		branch, err := getNearestBranch(latitude, longitude)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		branchJSON, err := json.Marshal(&branch)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(branchJSON)
		return
	case http.MethodPost:
		var newBranch Branch
		bodyBytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &newBranch)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if newBranch.BranchID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = insertBranch(&newBranch)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
