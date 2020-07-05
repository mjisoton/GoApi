package models

//Some dependencies
import "database/sql"

//Third-Party dependencies
import _ "github.com/go-sql-driver/mysql"

//Global variable to connect to MariaDB (it's thread-safe, is it's OK to use as global)
var db *sql.DB

//Row Abstraction
type Row interface {
    Scan(...interface{}) error
}

//Scanner Abstraction
type RowScanner interface {
    ScanRow(Row) error
}

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

	return nil
}

//Helpers that simplify the row scanning proccess
func QueryRows(rs RowScanner, query string, params ...interface{}) error {
    rows, err := db.Query(query, params...)
    if err != nil {
        return err
    }

    defer rows.Close()

    for rows.Next() {
        if err := rs.ScanRow(rows); err != nil {
            return err
        }
    }
    return rows.Err()
}

func QueryRow(rs RowScanner,query string, params ...interface{}) error {
    return rs.ScanRow(db.QueryRow(query, params...))
}
