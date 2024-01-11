package db

import (
	"fmt"
	"log"
	"time"

	"github.com/atefeh-syf/E-Wallet/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func InitDB(cfg *config.Config) error {
	var err error
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.SSLMode,
	)

	dbClient, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, err := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	//migrate tables
	err = MigrateEntities(dbClient)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Db connection established")
	return nil
}

func GetDb() *gorm.DB {
	return DBClient
}

func CloseDb() {
	conn, _ := DBClient.DB()
	conn.Close()
}
