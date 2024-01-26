package db

import (
	"fmt"

	logs "github.com/ashiqsabith123/love-bytes-proto/log"
	"github.com/ashiqsabith123/notification-svc/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(config config.Config) *gorm.DB {
	connstr := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.Postgres.Host, config.Postgres.User, config.Postgres.Database, config.Postgres.Port, config.Postgres.Paswword)
	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})

	if err != nil {
		logs.ErrLog.Fatal("Failed to connect database - ", err)
		return nil
	}

	err = db.AutoMigrate()

	if err != nil {
		logs.ErrLog.Fatalln(err)
	}

	logs.GenLog.Println("Database connected succesfully....")

	return db
}
