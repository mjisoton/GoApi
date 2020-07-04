package utils

//Some Dependencies
import (

	"github.com/go-ozzo/ozzo-validation"
    "github.com/spf13/viper"
)

//Validate the config struct according to predefined rules
func (config AppConfigType) Validate() error {
    return validation.ValidateStruct(&config,

		//MariaDB
		validation.Field(&config.Database_host, validation.Required),
		validation.Field(&config.Database_port, validation.Required),
		validation.Field(&config.Database_user, validation.Required),
		validation.Field(&config.Database_pass, validation.NotNil),

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
