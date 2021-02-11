package config

import (
	"fmt"
	"log"

	"github.com/confetti-framework/support/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     env.Str("DB_HOST"),
		Port:     env.Int("DB_PORT"),
		User:     env.Str("DB_USERNAME"),
		Password: env.Str("DB_PASSWORD"),
		DBName:   env.Str("DB_DATABASE"),
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
func NewDB(params ...string) *gorm.DB {
	var err error
	conString := DbURL(BuildDBConfig())
	log.Print(conString)

	DB, err = gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
