package db

import (
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/dulwin/ezebot/utils"
)

type Query struct {
	gorm.Model
	Category string `gorm:"size:127;primary_key"`
    Query string `gorm:"size:127;primary_key"`
    Response string
}
//TODO: Refactor to using classes and type

func GetInstance() *gorm.DB {
    dbName := os.Getenv("DB_NAME")
    db, err := gorm.Open("sqlite3", dbName)
    utils.CheckError(err)
    return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Query{})
}

func Insert(db *gorm.DB, q *Query) {
	db.Create(q)
}
