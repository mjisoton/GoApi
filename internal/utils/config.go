package util

//Some Dependencies
import (
	"os"
	"errors"
	"strconv"
	"encoding/json"
	"io/ioutil"

)

//Validate the config struct according to predefined rules
func (config AppConfigType) Validate() error {

	//If there is no socket string, it will try to connect via TCP
	if config.Database_socket == "" {

		//TCP requires this parameters (which are not required by socket)
		if config.Database_host == "" || config.Database_pass == "" || config.Database_port == 0 {
			return errors.New("The config file is missing valid database parameters.")
		}
	}

	//Now, check the following fields as well, since they're required for both socket and TCP connection
	if config.Database_name == "" || config.Database_user == "" {
		return errors.New("The config file is missing valid database parameters.")
	}

	//Check if the redis parameters were filled
	if config.Redis_min_conn == 0 || config.Redis_socket == "" {
		return errors.New("The confif file is missing valid NoSQL parameters.")
	}

	//Checks if the Server port is filled
	if config.Server_port == 0 {
		return errors.New("The config file is missing the HTTP server port.")
	}

	return nil
}

//Load the config file if it exists, and populate the AppConfig struct with the values
func LoadConfigFile(AppConfig *AppConfigType, file string) error {

	//First of all, check if the file exists
	info, err := os.Stat(file)
	if os.IsNotExist(err) == true || info.IsDir() == true {
		 return err
 	}

	//Read file to a byte slice
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	//Decode to struct
	err = json.Unmarshal(content, &AppConfig)
	if err != nil {
		return err
	}

	return AppConfig.Validate()
}

//Gets the DSN string to connect to a relacional database
func (config AppConfigType) GetDSN() string {

	//If there is a socket declared, use it
	if config.Database_socket != "" {
		return config.Database_user + "@unix(" + config.Database_socket + ")/" + config.Database_name
	} else {
		return config.Database_user + ":" + config.Database_pass + "@tcp(" + config.Database_host + ":" + strconv.Itoa(config.Database_port) + ")/" + config.Database_name
	}
}
