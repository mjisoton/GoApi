package models

//Third-Party dependencies
import "database/sql"
import "log"

type StoreTokenType struct {
		token string
		id int

}

func SearchStoreByToken(token string) (string, error) {
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
	        return "", err
	    } else {
			log.Println(err)
		}
	}

	return id_loja, nil
}
