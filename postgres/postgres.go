package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Db is our database struct for inherient
type Db struct {
	*sql.DB
}

// New make a new database using the connection string
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Db{db}, nil
}

// ConnString returns a connection string
func ConnString(host string, port int, user string, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbName)
}

// Duser shape
type Duser struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

// GetUsersByName is called within our user query for graphql
func (d *Db) GetUsersByName(name string) []Duser {
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("SELECT * FROM dusers WHERE name=$1")
	if err != nil {
		fmt.Println("GetUsersByName Query Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUsersByName Query Err: ", err)
	}

	// Create User struct for holding each row's data
	var r Duser
	// Create slice of Users for our response
	dusers := []Duser{}
	// Copy the columns from row into the values pointed at by r
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		dusers = append(dusers, r)
	}

	return dusers
}
