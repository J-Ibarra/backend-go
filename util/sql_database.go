package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type properties struct {
	SqlitePath string `env:"SQL_LITE_PATH,default=.data/sqlite.db"`
}

var (
	// DB connection
	DB *gorm.DB
)

func init() {
	var prop properties
	LoadFromEnv(&prop)

	db, err := gorm.Open(sqlite.Open(prop.SqlitePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
