package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConnString struct {
	DBUser     string `json:"DBUser"`
	DBPassword string `json:"DBPassword"`
	DBName     string `json:"DBName"`
	DBHost     string `json:"DBHost"`
	DBPort     string `json:"DBPort"`
	DBType     string `json:"DBType"`
}

func GetConnString() ConnString {
	var cfg ConnString

	cfgFile, _ := os.Open("config/connString.json")
	defer cfgFile.Close()

	jsonParser := json.NewDecoder(cfgFile)
	jsonParser.Decode(&cfg)
	return cfg
}

func GetDBType() string {
	return GetConnString().DBType
}

func GetPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		GetConnString().DBHost,
		GetConnString().DBPort,
		GetConnString().DBUser,
		GetConnString().DBName,
		GetConnString().DBPassword)
	return dataBase
}
