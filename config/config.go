package config

import (
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file" + err.Error())
	}
}

func InitFromTest() {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))

	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("test")
	config.AddConfigPath("../config/")
	config.AddConfigPath(filepath.Dir(d) + "/config/")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file" + err.Error())
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
