package config

import (
	"encoding/json"
	"log"
	"os"

	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func manageError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() Config {
	var c Config
	file, err := os.Open("./config.json")
	manageError(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	manageError(err)
	return c
}

func GetConnection() *gorm.DB {
	c := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	manageError(err)
	//defer db.Close()
	return db
}
