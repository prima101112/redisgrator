package config

import (
	"log"
	"errors"
	"gopkg.in/gcfg.v1"
)

type RedisHostCfg struct {
	Origin	    string
	Destination string
}

type General struct {
	Port int
	SetToDestWhenGet int
}

type Config struct {
	General	  General
	RedisHost RedisHostCfg
}

var Cfg Config

func ReadConfig(path string) bool {
	err := gcfg.ReadFileInto(&Cfg, path+"config.ini")
	log.Printf("%+v",Cfg)
	if err = Cfg.Validate(); err != nil {
		log.Println("failed read config ", err.Error())
		return false
	}
	if err == nil {
		log.Println("read config from ", path)
		return true
	}
	log.Println("Error : ", err)
	return false
}

func (c *Config) Validate() error {
	if c.General.SetToDestWhenGet > 1 {
		return errors.New("SetToDestWhenGet must be 0 or 1")
	}
	return nil
}