package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"loan-server/config"
)

type MyDb struct {
	MyDbConfig *config.Db
	Db         *gorm.DB
}

func Init(dbConfig *config.Db) (*MyDb, error) {
	// ---- Init Db ----
	gormDb, err := gorm.Open(mysql.Open(dbConfig.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	db := MyDb{
		MyDbConfig: dbConfig,
		Db:         gormDb,
	}
	return &db, nil
}
