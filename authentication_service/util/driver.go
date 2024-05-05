package util

import (
	"fmt"

	"authentication_service/config"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseDriver string

const (
	PostgreSQL DatabaseDriver = "postgres"

	// Static DatabaseDriver = "static"

	// SqlServer DatabaseDriver = "sqlserver"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	PostgreSQL *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Driver1 {
	case "PostgreSQL":
		db.Driver = PostgreSQL
		db.PostgreSQL = NewPostgreSQL(config)
	default:
		panic("Database driver not supported")
	}
	return &db
}
func NewPostgreSQL(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DB_Host1,
		config.DB_User1,
		config.DB_Pass1,
		config.DB_Name1,
		config.DB_Port1)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	log.Debug().Msg(dsn)

	if err != nil {
		panic(err)
	}

	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.PostgreSQL != nil {
		db, _ := db.PostgreSQL.DB()
		db.Close()
	}
}

// connection 2
type DatabaseDriver2 string

const (
	PostgreSQL2 DatabaseDriver = "postgres"
)

type DatabaseConnection2 struct {
	Driver DatabaseDriver

	PostgreSQL2 *gorm.DB
}

func NewConnectionDatabase2(config *config.AppConfig) *DatabaseConnection2 {
	var db DatabaseConnection2

	switch config.Driver2 {
	case "PostgreSQL":
		db.Driver = PostgreSQL
		db.PostgreSQL2 = NewPostgreSQL2(config)
	default:
		panic("Database driver not supported")
	}
	return &db
}
func NewPostgreSQL2(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DB_Host2,
		config.DB_User2,
		config.DB_Pass2,
		config.DB_Name2,
		config.DB_Port2)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	return db
}

func (db *DatabaseConnection2) CloseConnection2() {
	if db.PostgreSQL2 != nil {
		db, _ := db.PostgreSQL2.DB()
		db.Close()
	}
}
