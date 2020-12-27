package docs

import "github.com/miguelapabenedit/fravega-challange/entity"

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

// swagger:route GET /branch/{branchID} branchId
// Gets a branch base on an branchId
// responses:
// 200: branchResponse
// 400:

// swagger:route GET /branch currentPosition
// Gets the nearest branch for deliver base on a position
// responses:
// 200: branchResponse
// 400:

// swagger:route Post /branch branchParams
// Creates a new branch
// responses:
// 200:
// 400:

// BranchResponseWrapper is a response that contains a branch data
// swagger:response branchResponse
type BranchResponseWrapper struct {
	// in:body
	Body entity.Branch
}

// BranchParamsWrapper holds the structure to create or update a branch
// swagger:parameters branchParams
type BranchParamsWrapper struct {
	// required
	// in: body
	Body entity.Branch
}

// GetBranchParamsWrapper holds the params needs to get a branch by id
// swagger:parameters branchId
type GetBranchParamsWrapper struct {
	// required
	// in: path
	BranchID int32 `json:"branchID"`
}

// CurrentPositionParamsWrapper holds a current position
// swagger:parameters currentPosition
type CurrentPositionParamsWrapper struct {
	// required
	// in: query
	Latitude float32 `json:"latitude"`
	// required
	Longitude float32 `json:"longitude"`
}
