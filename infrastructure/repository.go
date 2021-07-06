package infrastructure

import "github.com/miguelapabenedit/fravega-challange/entity"

/*Repository interface defines basic behavior for crud of a branch
 */
type Repository interface {
	GetBranch(branchID int) (*entity.Branch, error)
	GetNearestDeliver(latitude float32, longitude float32) (*entity.Branch, error)
	SaveBranch(branch *entity.Branch) error
}
