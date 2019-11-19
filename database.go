package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// PostgreSQLConnection is a struct that contains an SQL connection
type PostgreSQLConnection struct {
	db *sql.DB
}

var (
	connection *PostgreSQLConnection
)

// PGSQLInit initializes the database
func PGSQLInit() {
	file, err := ioutil.ReadFile("database/database.sql")

	if err != nil {
		log.Fatalf("[+] Couldn't init database. Reason %v", err)
	}

	db := PGSQLConnect()
	if ok, err := db.Health(); !ok {
		log.Fatalf("[!] Database is not healthy. Reason %v", err)
	}

	// Executes the file d
	_ = db.Execute(string(file), nil)
}

// PGSQLConnect to the PostgreSQL database configured
func PGSQLConnect() *PostgreSQLConnection {
	if connection == nil {
		dbHost := configuration.Database.Host
		dbPort := configuration.Database.Port
		dbUser := configuration.Database.User
		dbPassword := configuration.Database.Password
		database := configuration.Database.DatabaseName

		conn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s"+
			" sslmode=disable", dbHost, dbPort, dbUser, dbPassword, database)

		db, err := sql.Open("postgres", conn)

		// Based on the users of the application
		db.SetMaxOpenConns(10)

		if err != nil {
			log.Printf("[!] Couldn't connect to the database. Reason %v\n", err)
			return nil
		}

		err = db.Ping()

		if err != nil {
			log.Printf("[!] Couldn't ping to the database. Reason %v\n", err)
			return nil
		}
		connection = &PostgreSQLConnection{}
		connection.db = db
	}

	return connection
}

// Health returns the status of the database connection c
func (c *PostgreSQLConnection) Health() (bool, error) {
	timer := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), timer)
	defer cancel()

	err := c.db.PingContext(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Close closes the connection with the database. From the SQL package: It
// is rarely necessary to close a DB.
func (c *PostgreSQLConnection) Close() {
	try := 0
	err := c.db.Close()

	if err != nil {
		log.Printf("[!] Failed to close the connection. Reason %v\n", err)
		for ; try < 3; try++ {
			log.Printf("[-] Trying to close the connection again\n")
			err = c.db.Close()
		}

		if err != nil {
			log.Printf("[!] Database connection couldn't be closed. Reason %v\n",
				err)
		}
	}
}

// Query queries the database for a given query with its arguments
func (c *PostgreSQLConnection) Query(query string, arguments []interface{}) *sql.Rows {
	stmt, err := c.db.Prepare(query)

	if err != nil {
		log.Printf("[!] Couldn't prepare statement. Reason %v", err)
		return nil
	}

	result, err := stmt.Query(arguments...)

	if err != nil {
		log.Printf("[!] Couldn't execute query. Reason %v", err)
		return nil
	}

	return result
}

// Execute queries the database for a given query and returns the result of the
// execution.
func (c *PostgreSQLConnection) Execute(query string, arguments []interface{}) sql.Result {
	if arguments == nil {
		result, err := c.db.Exec(query)

		if err != nil {
			log.Printf("[!] Couldn't execute query. Reason %v", err)
			return nil
		}

		return result
	}

	stmt, err := c.db.Prepare(query)

	if err != nil {
		log.Printf("[!] Couldn't prepare statement. Reason %v", err)
		return nil
	}

	result, err := stmt.Exec(arguments...)

	if err != nil {
		log.Printf("[!] Couldn't execute query. Reason %v", err)
		return nil
	}

	return result
}
