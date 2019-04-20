package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//Haobo : Auto generated from https://mholt.github.io/json-to-go/
type Configuration struct {
	Listen struct {
		ServerAddr string `json:"ServerAddr"`
		ServerPort string `json:"ServerPort"`
	} `json:"Listen"`
}

var Config Configuration

func LoadConfig(fileName string) error {
	configFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Print("LoadConfig() : Unable to open config.json, please check!")
		Config = Configuration{}
		return err
	}
	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		log.Print("LoadConfig() : Invalid config.json, please check!")
		Config = Configuration{}
		return err
	}
	return nil
}
