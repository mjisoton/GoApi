package utils

//Some Dependencies
import (

	"github.com/go-ozzo/ozzo-validation"
    "github.com/spf13/viper"
)

//Validate the config struct according to predefined rules
func (config AppConfigType) Validate() error {
    return validation.ValidateStruct(&config,

		validation.Field(&config.Db_host, validation.Required),
		validation.Field(&config.Db_port, validation.Required),
		validation.Field(&config.Db_user, validation.Required),
		validation.Field(&config.Db_pass, validation.NotNil),
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
