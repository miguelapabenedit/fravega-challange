package service

import (
	"testing"

	"github.com/miguelapabenedit/fravega-challange/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBranchIdNotEqualT0(t *testing.T) {
	testService := NewBranchService(nil)
	testBranch := entity.Branch{BranchID: 12, Address: "test", Latitude: 0, Longitude: 0}

	err := testService.AddBranch(&testBranch)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The Id must be 0")
}

func TestAddBranchEmtpyAdrress(t *testing.T) {
	testService := NewBranchService(nil)
	testBranch := entity.Branch{BranchID: 0, Address: "", Latitude: 10, Longitude: 10}

	err := testService.AddBranch(&testBranch)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Address cant be empty")
}

func TestAddBranchEmptyBranch(t *testing.T) {
	testService := NewBranchService(nil)

	err := testService.AddBranch(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The branch can't be empty")
}

func TestAddBranch(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("SaveBranch").Return(nil)

	testService := NewBranchService(mockRepo)
	testBranch := entity.Branch{BranchID: 0, Address: "San Andres", Latitude: 1254, Longitude: 6652}

	err := testService.AddBranch(&testBranch)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) SaveBranch(branch *entity.Branch) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockRepository) GetNearestDeliver(latitude float32, longitude float32) (*entity.Branch, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Branch), args.Error(1)
}

func (mock *MockRepository) GetBranch(branchID int) (*entity.Branch, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Branch), args.Error(1)
}
