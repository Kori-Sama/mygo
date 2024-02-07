package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Server     serverConfig     `yaml:"server"`
	Database   databaseConfig   `yaml:"database"`
	Blockchain blockchainConfig `yaml:"blockchain"`
	JwtConfig  jwtConfig        `yaml:"jwt"`
}

type serverConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

type databaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	Charset  string `yaml:"charset"`
}

type blockchainConfig struct {
	GethClient      string `yaml:"gethClient"`
	KeystorePath    string `yaml:"keystorePath"`
	ContractAddress string `yaml:"contractAddress"`
	OwnerPassword   string `yaml:"ownerPassword"`
	OwnerAddress    string `yaml:"ownerAddress"`
}

type jwtConfig struct {
	TokenExpire   int64  `yaml:"tokenExpire"`
	RefreshExpire int64  `yaml:"refreshExpire"`
	Secret        string `yaml:"secret"`
}

var Server serverConfig
var Database databaseConfig
var Blockchain blockchainConfig
var JwtConfig jwtConfig

func InitLog() {
	log.SetPrefix("MyGO: ")

	f, err := os.OpenFile("log/sys.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		if f, err = os.Create("log/server.log"); err != nil {
			log.Fatalf("Failed to create log file: %v", err)
		}
	}
	log.SetOutput(f)

	// f, err = os.OpenFile("log/http.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	if f, err = os.Create("log/http.log"); err != nil {
	// 		log.Fatalf("Failed to create log file: %v", err)
	// 	}
	// }
	// return io.MultiWriter(gin.DefaultWriter, f)
}

func InitConfig() {
	bytes, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	config := config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("Failed to parser config file: %v", err)
	}

	Server = config.Server
	Database = config.Database
	Blockchain = config.Blockchain
	JwtConfig = config.JwtConfig

}

const Logo = `
  __  ____     _______  ____    _   _   _   _   _ 
 |  \/  \ \   / / ____|/ __ \  | | | | | | | | | |
 | \  / |\ \_/ / |  __| |  | | | | | | | | | | | |
 | |\/| | \   /| | |_ | |  | | | | | | | | | | | |
 | |  | |  | | | |__| | |__| | |_| |_| |_| |_| |_|
 |_|  |_|  |_|  \_____|\____/  (_) (_) (_) (_) (_)

`
