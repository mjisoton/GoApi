package utils

//Some Dependencies
import (
	"strconv"

	//Third Party
	"github.com/go-ozzo/ozzo-validation"
    "github.com/spf13/viper"
)

//Validate the config struct according to predefined rules
func (config AppConfigType) Validate() error {
    return validation.ValidateStruct(&config,

		//MariaDB
		validation.Field(&config.Database_socket, validation.NotNil),
		validation.Field(&config.Database_host, validation.Required),
		validation.Field(&config.Database_port, validation.Required),
		validation.Field(&config.Database_user, validation.Required),
		validation.Field(&config.Database_pass, validation.NotNil),
		validation.Field(&config.Database_name, validation.Required),

		//NoSQL
		validation.Field(&config.Redis_socket, validation.Required),

		//HTTP Server
		validation.Field(&config.Server_port, validation.Required),
    )
}

//Load the config file, scanning the list of directories passed as parameter
func LoadConfigFile(AppConfig *AppConfigType, path string) error {
	v := viper.New()

	//Looks for a config.json
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(path)

	//Try reading the files
	if err := v.ReadInConfig(); err != nil {
	    return err
	}

	//Try to decode the JSON file(s)
	if err := v.Unmarshal(&AppConfig); err != nil {
	    return err
	}

	//Validate the struct after the JSON decode
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
