package utils

//Application config
type AppConfigType struct {

	//Config values related to the relational database (MariaDB)
	Db_host string
	Db_port int
	Db_user string
	Db_pass string

	//Config values related to noSQL database (Redis)
	R_socket string

	//Config values related to the HTTP server itself
	S_Port int 
}
