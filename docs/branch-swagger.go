package docs

import "github.com/miguelapabenedit/fravega-challange/branch"

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
// 400: http.BadRequest

// swagger:route Post /branch branchParams
// Creates a new branch
// responses:
// 200: branchResponse
// 400: http.BadRequest

// swagger:response branchResponse
type BranchResponseWrapper struct {
	// BranchResponseWrapper is a response that contains a branch data
	// in:body
	Body branch.Branch
}

// BranchParamsWrapper holds the structure to create or update a branch
// swagger:parameters branchParams
type BranchParamsWrapper struct {
	// required
	// in: body
	Body branch.Branch
}

// GetBranchParamsWrapper holds the params needs to get a branch by id
// swagger:parameters branchId
type GetBranchParamsWrapper struct {
	// required
	// in: path
	BranchID int32 `json:"branchID"`
}
