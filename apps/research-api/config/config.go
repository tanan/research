package config

type Config struct {
	Database Database
	Verbose  bool
}

type Database struct {
	host   string
	port   string
	dbname string
	schema string
	user   string
	pass   string
}
