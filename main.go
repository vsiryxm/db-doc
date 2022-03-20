package main

import (
	"db-doc/database"
	"db-doc/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var dbConfig model.DbConfig

func main() {

	type Yaml struct {
		App struct {
			Name    string `yaml:"name"`
			DocType int    `yaml:"doctype"`
		}
		Database struct {
			DbType   int    `yaml:"dbtype"`
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			DbName   string `yaml:"dbname"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"database"`
	}

	config := new(Yaml)
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if nil == err {
		err = yaml.Unmarshal(yamlFile, config)
		if err != nil {
			log.Fatalf("Unmarshal: %v when to struct", err)
		}
		dbConfig = model.DbConfig{
			DbType:   config.Database.DbType,
			DocType:  config.App.DocType,
			Host:     config.Database.Host,
			Port:     config.Database.Port,
			User:     config.Database.User,
			Password: config.Database.Password,
			Database: config.Database.DbName,
		}
	} else {
		fmt.Println("? Database type:\n1:MySQL\n2:SQL Server\n3:PostgreSQL")
		// db type
		fmt.Scanln(&dbConfig.DbType)
		if dbConfig.DbType < 1 || dbConfig.DbType > 3 {
			fmt.Println("wrong number, will exit ...")
			os.Exit(0)
		}
		GetDefaultConfig()
		// db host
		fmt.Println("? Database host (127.0.0.1) :")
		fmt.Scanln(&dbConfig.Host)
		// db port
		fmt.Printf("? Database port (%d) :\n", dbConfig.Port)
		fmt.Scanln(&dbConfig.Port)
		// db user
		fmt.Printf("? Database username (%s) :\n", dbConfig.User)
		fmt.Scanln(&dbConfig.User)
		// db password
		fmt.Println("? Database password (123456) :")
		fmt.Scanln(&dbConfig.Password)
		// db name
		fmt.Println("? Database name:")
		fmt.Scanln(&dbConfig.Database)
		// doc type
		fmt.Println("? Document type:\n1:Online\n2:Offline")
		fmt.Scanln(&dbConfig.DocType)
	}

	// generate
	database.Generate(&dbConfig)
}

// GetDefaultConfig get default config
func GetDefaultConfig() {
	dbConfig.Host = "127.0.0.1"
	dbConfig.Password = "123456"
	if dbConfig.DbType == 1 {
		dbConfig.Port = 3306
		dbConfig.User = "root"
	}
	if dbConfig.DbType == 2 {
		dbConfig.Port = 1433
		dbConfig.User = "sa"
	}
	if dbConfig.DbType == 3 {
		dbConfig.Port = 5432
		dbConfig.User = "postgres"
	}
}
