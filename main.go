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
	fmt.Printf("Database Config: %s:%d\n", AppConfig.Db_host, AppConfig.Db_port)
	fmt.Println("###########################################################")


}
