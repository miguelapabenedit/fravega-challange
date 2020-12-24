package branch

/*Branch represents a delivery location
 */
type Branch struct {
	BranchID  int     `json:"branchID"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
