package mysql

import (
	"axe/gl"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func InitMySQL() (err error) {
	db, err = sqlx.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {
		gl.App.Logger().Fatal("connect server failed, err: %v\n", err)
		return
	}
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})
	db.SetMaxOpenConns(viper.GetInt("db.active"))
	db.SetMaxIdleConns(viper.GetInt("db.idle"))
	return
}
