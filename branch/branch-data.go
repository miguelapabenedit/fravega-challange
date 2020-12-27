package branch

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/miguelapabenedit/fravega-challange/database"
)

func getBranch(branchID int) (*Branch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := database.DbConn.QueryRowContext(ctx, `SELECT 
	branchId, 
	address, 
	latitude, 
	longitude 
	FROM Branches 
	WHERE branchId = @p1`, branchID)

	branch := &Branch{}
	err := row.Scan(
		&branch.BranchID,
		&branch.Address,
		&branch.Latitude,
		&branch.Longitude,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	return branch, nil
}

func getNearestBranch(latitude float64, longitude float64) (*Branch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := database.DbConn.QueryRowContext(ctx, `SELECT TOP 1 
		BranchID,
		Address,
		Latitude,
		Longitude,
		Round(acos(sin(RADIANS(latitude)) * sin(RADIANS(@p1))
		+ cos(RADIANS(latitude)) * cos(RADIANS(@p1)) 
		* cos(RADIANS(longitude) - RADIANS(@p2))) 
		* 6371,2) as distance FROM Branches ORDER BY distance`,
		latitude,
		longitude)

	branch := &Branch{}

	var distance float64
	err := row.Scan(
		&branch.BranchID,
		&branch.Address,
		&branch.Latitude,
		&branch.Longitude,
		&distance,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}

	return branch, nil
}

func insertBranch(branch *Branch) error {
	_, err := database.DbConn.Exec(`INSERT INTO Branches
	(address,
	latitude,
	longitude
	)VALUES(@p1, @p2, @p3)`,
		branch.Address,
		branch.Latitude,
		branch.Longitude)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
