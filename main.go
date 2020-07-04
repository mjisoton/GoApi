package main

//Some dependencies
import "fmt"
import "os"

//Custom packs
import "./utils"

//... and the magic starts
func main() {

	//Global with the app configuration
	var AppConfig utils.AppConfigType

	//Load config.json from the executable's directory
	if err := utils.LoadConfigFile(&AppConfig, "."); err != nil {
		fmt.Printf("[ERROR] Configuration error: %s", err)
		os.Exit(1)
	}

	//Just a simple greeter
	fmt.Println("###########################################################")
	fmt.Println("WebShopping API @ AgênciaNet - by @mjisoton, v0.0.1")
	fmt.Printf("MariaDB Config: %s:%d\n", AppConfig.Database_host, AppConfig.Database_port)
	fmt.Printf("Redis Config: %s\n", AppConfig.Redis_socket)
	fmt.Printf("Server Config: %d\n", AppConfig.Server_port)
	fmt.Println("###########################################################")


}
