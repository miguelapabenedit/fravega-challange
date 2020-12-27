package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

/*DbConn represents a connection to the db
 */
var DbConn *sql.DB
var server = "localhost"
var port = 1434
var password = "go123"
var user = "go"
var database = "FravegaChallange"

/*SetupDatabase configures and start a new connection to the database
 */
func SetupDatabase() {
	var err error
	cs := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)
	DbConn, err = sql.Open("sqlserver", cs)

	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(4)
	DbConn.SetMaxOpenConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
