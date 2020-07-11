package models

//Third-Party dependencies
import "database/sql"
import "log"

func SearchStoreByToken(token string) (string, bool) {
	var id_loja string

	row := db.QueryRow(
	`
		SELECT
			id
		FROM loj_lojas
		WHERE ativo = 1 AND token = ?
		LIMIT 1
	`, token)

	err := row.Scan(&id_loja)

	if err != nil {
	    if err == sql.ErrNoRows {
	        return "", false
	    } else {
			log.Fatalf("Failed to request the token from the database.")
		}
	}

	return id_loja, true
}
