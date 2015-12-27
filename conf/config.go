package conf

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var (
	filepath     = ""
	config       *Config
	ERR_NO_MONGO = errors.New("not found mongo config.")
)

type Config struct {
	Name  string            `json:"name"`
	Mongo map[string]*Mongo `json:"mongo"`
	Mysql map[string]*Mysql `json:"mysql"`
	Redis map[string]*Redis `json:"redis"`
}

type Mongo struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Redis struct {
	Name           string `json:"name"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Authentication string `json:"authentication"`
}

type Mysql struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	if filepath == "" {
		filepath = os.Getenv("GOPATH") + "/src/github.com/black-tech/realtime/conf/v1.json"
	}
	_, err := getConf()
	if err != nil {
		log.Fatal("config failed, ", err)
	}
}

func getConf() (c *Config, err error) {
	if config != nil {
		return config, nil
	}
	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	c = new(Config)
	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		log.Println(err)
		log.Println(c)
		return
	}
	config = c
	return
}

func GetConfig() *Config {
	return config
}

func GetMysqlConfig(name string) *Mysql {
	return config.Mysql[name]
}

func GetMGOConfig(name string) (mc *Mongo, err error) {
	c, err := getConf()
	if err != nil {
		return
	}
	if mc, ok := c.Mongo[name]; ok {
		return mc, nil
	}
	return nil, ERR_NO_MONGO
}
