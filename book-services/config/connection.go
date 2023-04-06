package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
	RunnApp struct {
		Host string
		Port string
	}
	JWTSecret string
	RabbitMQ  struct {
		Host string
		Port string
		User string
		Pass string
	}
}

var app *App
var DB *gorm.DB

func GetConnection() *gorm.DB {
	conf := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	return db
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}
func initConfig() *App {
	conf := App{}
	err := godotenv.Load()
	if err != nil {
		conf.Database.Host = "localhost"
		conf.Database.Port = "3306"
		conf.Database.User = "root"
		conf.Database.Pass = ""
		conf.Database.Name = "book"
		conf.JWTSecret = ""
		conf.RunnApp.Host = "localhost"
		conf.RunnApp.Port = "3000"
		conf.RabbitMQ.Host = ""
		conf.RabbitMQ.Port = ""
		conf.RabbitMQ.User = "guest"
		conf.RabbitMQ.Pass = "guest"

		return &conf
	}

	conf.Database.Host = os.Getenv("MYSQL_HOST")
	conf.Database.Port = os.Getenv("MYSQL_PORT")
	conf.Database.Name = os.Getenv("MYSQL_NAME")
	conf.Database.User = os.Getenv("MYSQL_USER")
	conf.Database.Pass = os.Getenv("MYSQL_PASS")
	conf.JWTSecret = os.Getenv("JWT_SECRET")
	conf.RabbitMQ.Host = os.Getenv("RABBIT_HOST")
	conf.RabbitMQ.Port = os.Getenv("RABBIT_PORT")
	conf.RunnApp.Host = os.Getenv("APP_HOST")
	conf.RunnApp.Port = os.Getenv("APP_PORT")
	conf.RabbitMQ.User = os.Getenv("RABBITMQ_USER")
	conf.RabbitMQ.Pass = os.Getenv("RABBITMQ_PASS")

	return &conf

}
