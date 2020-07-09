package main

//Some native dependencies
import "log"

//Custom packages
import "github.com/mjisoton/GoApi/internal/utils"
import "github.com/mjisoton/GoApi/internal/models"
import "github.com/mjisoton/GoApi/internal/caching"
import "github.com/mjisoton/GoApi/internal/router"

//... and the magic starts
func main() {

	//Global with the app configuration
	var AppConfig util.AppConfigType

	//Load config.json from the executable's directory
	if err := util.LoadConfigFile(&AppConfig, "./config.json"); err != nil {
		log.Fatalf("[ERROR] Configuration error: %s", err)
	}

	//Just a simple greeter
	log.Println("###########################################################")
	log.Println("WebShopping API @ AgênciaNet - by @mjisoton, v0.0.1")
	log.Printf("MariaDB Connection: %s\n", AppConfig.GetDSN())
	log.Printf("Redis Connection: %s\n", AppConfig.Redis_socket)
	log.Printf("Server Config: %d\n", AppConfig.Server_port)
	log.Println("###########################################################\n")

	//Try and open the connection to the MariaDB database
	err := models.Connect(AppConfig.GetDSN())
	if err != nil {
		log.Fatal("[ERROR] Failed to connect to the MariaDB database: %s", err)
	} else {
		log.Printf("[SUCCESS] Connection with Mariadb database established.\n")
	}

	//Try and connect to the Redis database
	err = caching.Connect(AppConfig.Redis_socket, AppConfig.Redis_min_conn)
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to the Redis database: %s", err)
	} else {
		log.Printf("[SUCCESS] Connection with Redis database established.\n")
	}

	//After establishing the connections, start the HTTP server_port
	err = router.Start(AppConfig.Server_port)
	if err != nil {
		log.Fatal("[ERROR] Failed to start the HTTP server and router.")
	}

}
