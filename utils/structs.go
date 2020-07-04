package utils

//Application config
type AppConfigType struct {

	//Config values related to the relational database (MariaDB)
	Database_host string
	Database_port int
	Database_user string
	Database_pass string

	//Config values related to noSQL database (Redis)
	Redis_socket string

	//Config values related to the HTTP server itself
	Server_port int 
}
