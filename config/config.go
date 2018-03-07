package config

import (
	"io/ioutil"
	"fmt"

	"flag"
	"encoding/json"
)

type mysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type mongoConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type Config struct {
	Mysql mysqlConfig `json:"mysql"`
	Mongo mongoConfig `json:"mongo"`
}

var config *Config = nil;

func loadConfig() *Config {
	flag.Parse()
	p := flag.String("p", "dev", "environment")
	filePath := "config/dev.json"
	if *p == "pro" {
		filePath = "pro.json"
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Read Conif Error: %s\n", err.Error()))
	}
	configJson := string(buf)
	var config *Config
	json.Unmarshal([]byte(configJson), &config)
	return config
}

func GetConfig() Config {

	if config == nil {
		config = loadConfig();
	}
	return *config;
}
