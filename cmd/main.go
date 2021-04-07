package main

import (
	"axe/dao/mysql"
	"axe/gl"
	"axe/v1/router"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	pflag.String("app.cacheDir", "./cache/", "cache directory")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	//fmt.Println(viper.GetString("app.cacheDir"))

	mysql.InitMySQL()

	app := iris.Default()
	gl.App = app

	router.InitRouter(app)
}
