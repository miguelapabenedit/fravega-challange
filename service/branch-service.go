package service

import (
	"fmt"
	"log"

	"github.com/miguelapabenedit/fravega-challange/entity"
	"github.com/miguelapabenedit/fravega-challange/infrastructure"
)

type service struct{}

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

func (*service) AddBranch(branch *entity.Branch) error {

	if branch == nil {
		return fmt.Errorf("The branch can't be empty")
	}

	if branch.BranchID != 0 {
		return fmt.Errorf("The Id must be 0")
	}

	if branch.Address == "" {
		return fmt.Errorf("Address cant be empty")
	}

	err := repo.SaveBranch(branch)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
