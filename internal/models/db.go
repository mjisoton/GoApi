package models

//Some dependencies
import "database/sql"
import "time"
import "encoding/json"
import "log"
import "reflect"

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


func queryToJson(db *sql.DB, query string, args ...interface{}) ([]byte, error) {
	var objects []map[string]interface{}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}

		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			v := reflect.New(column.ScanType()).Interface()
			switch v.(type) {
			case *[]uint8:
				v = new(string)
			default:
				// use this to find the type for the field
				// you need to change
				log.Printf("%v: %T", column.Name(), v)
			}

			object[column.Name()] = v
			values[i] = object[column.Name()]
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		objects = append(objects, object)
	}

	// indent because I want to read the output
	return json.MarshalIndent(objects, "", "\t")
}
