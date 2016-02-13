package main

import (
	"fmt"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
    "gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
)

var logger = loggo.GetLogger("main")
var rootLogger = loggo.GetLogger("")

func setupConfig() {
	viper.SetConfigType("yaml")
	// Setting default configuration in case no configuration file is porvided
	viper.SetDefault("Filename", "logs/client.log")
	// The max allowed file size before being archived is 100 MB
	viper.SetDefault("MaxSize", 500)
	viper.SetDefault("MaxBackups", 3)
	viper.SetDefault("MaxAge", 28)
	viper.SetDefault("AllowedClients", []string{"flex3", "html5"})
	viper.SetDefault("ServerPort", "8090")

	// Load application configuration
	viper.SetConfigName("config")                 // name of config file (without extension)
	viper.AddConfigPath(".")                      // path to look for the config file in
	viper.AddConfigPath("/etc/bbb-client-logger") // call multiple times to add many search paths

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	setupLogger()
}

func setupLogger() {
	fmt.Println("Setting up logger")
	loggo.ConfigureLoggers("<root>=TRACE")
	oldWriter, _, err := loggo.RemoveWriter("default")
	// err is non-nil if and only if the name isn't found.
	loggo.RegisterWriter("default", oldWriter, loggo.TRACE)
	fmt.Println("Log file: " + viper.GetString("Filename"))

	writer := &lumberjack.Logger{
		Filename:   viper.GetString("Filename"),
		MaxSize:    viper.GetInt("MaxSize"), // megabytes
		MaxBackups: viper.GetInt("MaxBackups"),
		MaxAge:     viper.GetInt("MaxAge"), // days
	}

	newWriter := loggo.NewSimpleWriter(writer, &ClientLogFormatter{})
	err = loggo.RegisterWriter("testfile", newWriter, loggo.TRACE)

	if err != nil {
		fmt.Println("Error registering logger")
	}
}

func main() {
	// Configure the server router
	setupConfig()
	
	router := NewRouter()
	// To init the debug logging uncomment the next line
	// Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":" + viper.GetString("ServerPort"), router))
}
