package controller

import "net/http"

type Controller interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	GetNearestDeliver(w http.ResponseWriter, r *http.Request)
}
