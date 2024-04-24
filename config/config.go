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
	Port     string   `yaml:"port"`
	GrpcPort string   `yaml:"grpcPort"`
	Mode     string   `yaml:"mode"`
	Salt     string   `yaml:"salt"`
	Dict     []string `yaml:"dict"`
}

type databaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	SSL      string `yaml:"ssl"`
	Limit    int    `yaml:"limit"`
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

	fetchEnv()
}

func fetchEnv() {
	port := os.Getenv("PORT")
	if port != "" {
		Server.Port = port
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort != "" {
		Server.GrpcPort = grpcPort
	}

	dbName := os.Getenv("DB_NAME")
	if dbName != "" {
		Database.DbName = dbName
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost != "" {
		Database.Host = dbHost
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort != "" {
		Database.Port = dbPort
	}

	dbUsername := os.Getenv("DB_USER")
	if dbUsername != "" {
		Database.Username = dbUsername
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		Database.Password = dbPassword
	}
}
