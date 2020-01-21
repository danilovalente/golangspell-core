package config

import (
	"github.com/spf13/viper"
)

//Values stores the current configuration values
var Values Config

//Config contains the application's configuration values. Add here your own variables and bind it on init() function
type Config struct {
	//Port contains the port in which the application listens
	Port string
	//AppName for displaying in Monitoring
	AppName string
	//LogLevel - DEBUG or INFO or WARNING or ERROR or PANIC or FATAL
	LogLevel string
	//TestRun state if the current execution is a test execution
	TestRun bool
}

func init() {
	viper.BindEnv("TestRun", "TESTRUN")
	viper.SetDefault("TestRun", false)
	viper.BindEnv("Port", "PORT")
	viper.SetDefault("Port", "8080")
	viper.BindEnv("AppName", "APP_NAME")
	viper.SetDefault("AppName", "{{.AppName}}")
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.SetDefault("LogLevel", "INFO")
	viper.Unmarshal(&Values)
}
