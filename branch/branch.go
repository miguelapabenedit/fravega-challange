package branch

/*Branch represents a delivery location
 */
type Branch struct {
	BranchID  int     `json:"branchID"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
