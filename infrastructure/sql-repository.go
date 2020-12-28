package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/miguelapabenedit/fravega-challange/entity"
)

var (
	dbConn   *sql.DB
	server   = os.Getenv("DB_SERVER")
	port     = os.Getenv("DB_PORT")
	password = os.Getenv("DB_PASS")
	user     = os.Getenv("DB_USER")
	database = os.Getenv("DB_NAME")
)

type repo struct{}

func NewSQLRepository() Repository {
	var err error

	cs := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", server, user, password, port, database)
	dbConn, err = sql.Open("sqlserver", cs)

	if err != nil {
		log.Fatal(err)
	}

	dbConn.SetMaxOpenConns(4)
	dbConn.SetConnMaxLifetime(60 * time.Second)

	return &repo{}
}

func (*repo) GetBranch(branchID int) (*entity.Branch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := dbConn.QueryRowContext(ctx, `SELECT 
	branchId, 
	address, 
	latitude, 
	longitude 
	FROM Branches 
	WHERE branchId = @p1`, branchID)

	branch := &entity.Branch{}
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

func (*repo) GetNearestDeliver(latitude float32, longitude float32) (*entity.Branch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := dbConn.QueryRowContext(ctx, `SELECT TOP 1 
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

	branch := &entity.Branch{}

	var distance float32
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

func (*repo) SaveBranch(branch *entity.Branch) error {
	_, err := dbConn.Exec(`INSERT INTO Branches
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
