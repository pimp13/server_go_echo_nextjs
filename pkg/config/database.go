package config

import "sync"

type Database struct {
	Password   string
	Username   string
	Dbname     string
	Port       string
	Host       string
	Connection string
}

var (
	databaseConfigInstance *Database
	dbOnce                 sync.Once
)

func GetDatabaseConfig() *Database {
	dbOnce.Do(func() {
		databaseConfigInstance = &Database{
			Dbname:     getEnv("DB_NAME", "database-name"),
			Username:   getEnv("DB_USER", "username"),
			Password:   getEnv("DB_PASSWORD", "******"),
			Host:       getEnv("DB_HOST", "localhost"),
			Port:       getEnv("DB_PORT", "3306"),
			Connection: getEnv("DB_CONNECTION", "sqlite3"),
		}
	})
	return databaseConfigInstance
}
