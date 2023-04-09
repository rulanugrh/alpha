package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Host string
	Port string

	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}

	RabbitMQ struct {
		Host string
		Port string
	}
}

var app *App
var DB *gorm.DB

func GetConnect() *gorm.DB {

	conf := GetConfig()
	login := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(login), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	return db

}

func GetConfig() *App {
	if app == nil {
		initConfig()
	}

	return app
}
func initConfig() *App {
	conf := App{}
	err := godotenv.Load()
	if err != nil {
		conf.Host = "localhost"
		conf.Port = ""
		conf.Database.Host = "localhost"
		conf.Database.Port = "3306"
		conf.Database.User = "root"
		conf.Database.Pass = ""
		conf.Database.Name = "orders"
		conf.RabbitMQ.Host = "localhost"
		conf.RabbitMQ.Port = ""

		return &conf
	}

	conf.Host = os.Getenv("APP_HOST")
	conf.Port = os.Getenv("APP_PORT")
	conf.Database.Host = os.Getenv("MYSQL_HOST")
	conf.Database.Port = os.Getenv("MYSQL_PORT")
	conf.Database.User = os.Getenv("MYSQL_USER")
	conf.Database.Pass = os.Getenv("MYSQL_PASS")
	conf.Database.Name = os.Getenv("MYSQL_NAME")
	conf.RabbitMQ.Host = os.Getenv("RABBIT_HOST")
	conf.RabbitMQ.Port = os.Getenv("RABBIT_PORT")

	return &conf
}
