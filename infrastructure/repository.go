package infrastructure

import "github.com/miguelapabenedit/fravega-challange/entity"

type Repository interface {
	GetBranch(branchID int) (*entity.Branch, error)
	GetNearestBranch(latitude float32, longitude float32) (*entity.Branch, error)
	SaveBranch(branch *entity.Branch) error
}
