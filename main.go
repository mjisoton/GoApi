package main

//Some native dependencies
import "log"
import "time"

//Custom packages
import "./utils"
import "./models"
import "./caching"
import "./router"

//... and the magic starts
func main() {

	//Global with the app configuration
	var AppConfig utils.AppConfigType

	//Load config.json from the executable's directory
	if err := utils.LoadConfigFile(&AppConfig, "./config.json"); err != nil {
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





	/*
	uList := new(models.SQLUserList)
	models.QueryRows(uList, `SELECT id, nome, email FROM usuarios WHERE ativo = ?`, 1)

	if uList.Len > 0 {
		for _, v := range uList.Res {
			log.Println(uList.Len, v.Nome)
		}
	}
	*/


}
