package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	log "github.com/sirupsen/logrus"
)

var _once sync.Once
var cfg *Configuration

// automatically run whenever package is imported
func init() {
	_once.Do(func() {
		log.SetLevel(log.DebugLevel)
		// Executed only once in app lifetime
		log.Info("loading config.json ..")
		file, e := ioutil.ReadFile("./config.json")
		if e != nil {
			log.Fatalf("error reading config.json: %v\n", e)
		}
		e = json.Unmarshal(file, &cfg)
		if e != nil {
			log.Fatalf("error parsing config.json: %v\n", e)
		}
		if !cfg.Debug {
			log.SetLevel(log.InfoLevel)
		}
		GetDb()
		initCache()
	})
}

// Configuration holds the app configuration, struct makes it easier to load & save it as json
type Configuration struct {
	Debug         bool   `json:"debug"`
	Port          uint   `json:"port"`
	DbHost        string `json:"dbHost"`
	DbPort        uint   `json:"dbPort"`
	DbUser        string `json:"dbUser"`
	DbPassword    string `json:"dbPassword"`
	DbSchema      string `json:"dbSchema"`
	CacheHost     string `json:"cacheHost"`
	CachePort     uint   `json:"cachePort"`
	CachePassword string `json:"cachePassword"`
	CacheDb       uint   `json:"cacheDb"`
}

// Get returns the global configuration object
func Get() *Configuration {
	return cfg
}
