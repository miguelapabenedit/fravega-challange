package service

import "github.com/miguelapabenedit/fravega-challange/entity"

type Service interface {
	GetBranch(branchID int) (*entity.Branch, error)
	GetNearestDeliver(latitude float32, longitude float32) (*entity.Branch, error)
	AddBranch(branch *entity.Branch) error
}
