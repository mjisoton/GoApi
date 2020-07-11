package models

//Some dependencies
import "database/sql"
import "time"

//Third-Party dependencies
import _ "github.com/go-sql-driver/mysql"

//Global variable to connect to MariaDB (it's thread-safe, is it's OK to use as global)
var db *sql.DB

//Initialize a connection with the database
func Connect(dsn string) error {
	var err error
    db, err = sql.Open("mysql", dsn)

	if err != nil {
        return err
    }

    if err = db.Ping(); err != nil {
        return err
    }

	//https://www.alexedwards.net/blog/configuring-sqldb
	//Max connections to the DB (idle + in use)
	db.SetMaxOpenConns(12)

	//Max connections in idle state
	db.SetMaxIdleConns(12)

	//Max connection lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	return nil
}
