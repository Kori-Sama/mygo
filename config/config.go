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
	LogConfig  logConfig        `yaml:"log"`
}

type serverConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
	Salt string `yaml:"salt"`
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

type logConfig struct {
	Level   string       `yaml:"level"`
	Console bool         `yaml:"console"`
	File    bool         `yaml:"file"`
	SysPath string       `yaml:"sysPath"`
	GinPath string       `yaml:"ginPath"`
	Format  formatConfig `yaml:"format"`
}

type formatConfig struct {
	Prefix    string `yaml:"prefix"`
	Timestamp string `yaml:"timestamp"`
}

var Server serverConfig
var Database databaseConfig
var Blockchain blockchainConfig
var Jwt jwtConfig
var Log logConfig

func InitConfig() {
	configPath := "config/config.yaml"

	bytes, err := os.ReadFile(configPath)
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
	Jwt = config.JwtConfig
	Log = config.LogConfig
}
