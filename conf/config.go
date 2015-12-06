package conf

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

var (
	filepath     = "conf/v1.json"
	config       *Config
	ERR_NO_MONGO = errors.New("not found mongo config.")
)

type Config struct {
	Name  string            `json:"name"`
	Mongo map[string]*Mongo `json:"mongo"`
	Redis map[string]*Redis `json:"redis"`
}

var Mongo struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Redis struct {
	Name           string
	Host           string
	Port           int
	Authentication string `json:"authentication"`
}

func GetConf() (c *Config, err error) {
	if config != nil {
		return config
	}
	f, err := os.Open(filepath)
	if err != nil {
		retrun
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return
	}
	config = c
	return
}

func GetMGOConfig(name string) (mc *Mongo, err error) {
	c, err := GetConf()
	if err != nil {
		return
	}
	if mc, ok := c.Mongo[name]; ok {
		return mc, nil
	}
	return nil, ERR_NO_MONGO
}
