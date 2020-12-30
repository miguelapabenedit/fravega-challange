package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

// GetBranch godoc
// @Summary Get a branch
// @ID get-string-by-int
// @Description Gets a branch base on an branchId
// @Tags branch
// @Accept json
// @Produce  json
// @Param id path int true "Branch ID"
// @Success 200 {object} entity.Branch
// @Failure 400,404 {object} http.Response
// @Failure 500 {object} http.Response
// @Router /branch/{id} [get]
func (*controller) Get(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "/")
	branchID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])

	if err != nil || branchID <= 0 {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	branch, err := serv.GetBranch(branchID)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if branch == nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	branchJSON, err := json.Marshal(&branch)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(branchJSON)
	return
}

// GetNearestDeliver godoc
// @Summary Get the nearest deliver
// @ID get-nearest-branch
// @Description Get the nearest deliver branch base on a position(latitude and longitude)
// @Tags branch
// @Accept json
// @Produce  json
// @Param latitude query number  true "Latitude"
// @Param longitude query number  true "Longitude"
// @Success 200 {object} entity.Branch
// @Failure 400,404 {object} http.Response
// @Failure 500 {object} http.Response
// @Router /branch/getNearestDeliver [get]
func (*controller) GetNearestDeliver(w http.ResponseWriter, r *http.Request) {
	latitudeForm := r.FormValue("latitude")
	longitudeForm := r.FormValue("longitude")

	if latitudeForm == "" || longitudeForm == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	latitude, err := strconv.ParseFloat(latitudeForm, 32)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(longitudeForm, 32)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	branch, err := serv.GetNearestDeliver(float32(latitude), float32(longitude))

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	branchJSON, err := json.Marshal(&branch)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(branchJSON)
	return
}

// Post godoc
// @Summary Creates a new branch
// @ID post-branch
// @Description Get the nearest deliver branch base on a position(latitude and longitude)
// @Tags branch
// @Accept json
// @Produce  json
// @Param branch body entity.Branch true "Branch"
// @Success 200 {object} entity.Branch
// @Failure 400,404 {object} http.Response
// @Failure 500 {object} http.Response
// @Router /branch [post]
func (*controller) Post(w http.ResponseWriter, r *http.Request) {
	var newBranch entity.Branch
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyBytes, &newBranch)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if newBranch.BranchID != 0 {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = serv.AddBranch(&newBranch)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
