package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type App struct {
	Protocol     string
	Address      string
	Static       string
	Log          string
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	RunTest      bool
}

type Db struct {
	Driver                string
	Address               string
	Database              string
	User                  string
	Password              string
	MaxConnections        int
	MaxIdeConnections     int
	ConnectionMaxLifeTime time.Duration
}

type Pagination struct {
	PerPage  int
	UrlQuery string
}

type Jwt struct {
	Secret     string
	ExpireTime time.Duration
}

type Configuration struct {
	App        App
	Db         Db
	Pagination Pagination
	Jwt        Jwt
}

var config *Configuration
var once sync.Once

func LoadConfig() *Configuration {
	//通过单例加载配置文件
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatal("open config error", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatal("Decode Config Error ", err)
		}
	})

	return config
}
