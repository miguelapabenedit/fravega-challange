package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/miguelapabenedit/fravega-challange/entity"
	"github.com/miguelapabenedit/fravega-challange/service"
)

var serv service.Service

type controller struct{}

func NewBranchController(service service.Service) Controller {
	serv = service
	return &controller{}
}

func (*controller) Get(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "/")
	branchID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])

	if err != nil || branchID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	branch, err := serv.GetBranch(branchID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if branch == nil {
		w.WriteHeader(http.StatusNotFound)
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
}

func (*controller) GetNearestDeliver(w http.ResponseWriter, r *http.Request) {
	latitudeForm := r.FormValue("latitude")
	longitudeForm := r.FormValue("longitude")

	if latitudeForm == "" || longitudeForm == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	latitude, err := strconv.ParseFloat(latitudeForm, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeForm, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	branch, err := serv.GetNearestDeliver(float32(latitude), float32(longitude))

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
}

func (*controller) Post(w http.ResponseWriter, r *http.Request) {
	var newBranch entity.Branch
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

	err = serv.SaveBranch(&newBranch)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
