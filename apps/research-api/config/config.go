package config

import "os"

var config Config

type Config struct {
	database Database
	verbose  bool
}

type Database struct {
	host   string
	port   string
	dbname string
	schema string
	user   string
	pass   string
}

func init() {
	// RESEARCH_DB_HOST
	if host := os.Getenv("RESEARCH_DB_HOST"); host == "" {
		config.database.host = "research-db-svc"
	} else {
		config.database.host = host
	}
	// RESEARCH_DB_PORT
	if port := os.Getenv("RESEARCH_DB_PORT"); port == "" {
		config.database.port = "5432"
	} else {
		config.database.port = port
	}
	// RESEARCH_DB_NAME
	if name := os.Getenv("RESEARCH_DB_NAME"); name == "" {
		config.database.dbname = "research"
	} else {
		config.database.dbname = name
	}
	// RESEARCH_DB_SCHEMA
	if schema := os.Getenv("RESEARCH_DB_SCHEMA"); schema == "" {
		config.database.schema = "research"
	} else {
		config.database.schema = schema
	}
	// RESEARCH_DB_USER
	if user := os.Getenv("RESEARCH_DB_USER"); user == "" {
		config.database.user = "dolphin"
	} else {
		config.database.user = user
	}
	// RESEARCH_DB_PASSWORD
	if pass := os.Getenv("RESEARCH_DB_PASSWORD"); pass == "" {
		config.database.pass = "dolphin"
	} else {
		config.database.pass = pass
	}
}

func GetDatabaseConnInfo() string {
	return "host=" + config.database.host + " port=" + config.database.port + " user=" + config.database.user + " dbname=" + config.database.dbname + " password=" + config.database.pass + " sslmode=disable"
}

func GetVerbose() bool {
	return config.verbose
}
