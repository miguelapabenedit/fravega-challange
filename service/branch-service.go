package service

import (
	"log"

	"github.com/miguelapabenedit/fravega-challange/entity"
	"github.com/miguelapabenedit/fravega-challange/infrastructure"
)

type service struct{}

type Service interface {
	GetBranch(branchID int) (*entity.Branch, error)
	GetNearestDeliver(latitude float32, longitude float32) (*entity.Branch, error)
	SaveBranch(branch *entity.Branch) error
}

var repo infrastructure.Repository

func NewBranchService(repository infrastructure.Repository) Service {
	repo = repository
	return &service{}
}

func (*service) GetBranch(branchID int) (*entity.Branch, error) {
	branch, err := repo.GetBranch(branchID)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return branch, nil
}

func (*service) GetNearestDeliver(latitude float32, longitude float32) (*entity.Branch, error) {
	branch, err := repo.GetNearestDeliver(latitude, longitude)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return branch, nil
}

func (*service) SaveBranch(branch *entity.Branch) error {
	err := repo.SaveBranch(branch)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
