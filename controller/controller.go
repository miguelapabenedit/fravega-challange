package controller

import "net/http"

/*Controller interface implements the bassic handlers for the branch api
 */
type Controller interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	GetNearestDeliver(w http.ResponseWriter, r *http.Request)
}
