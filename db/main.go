package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

    "github.com/dulwin/ezebot/utils"
)

type Query struct {
    gorm.Model
    category string `gorm:"size:127;primary_key"`
    query string `gorm:"size:127;primary_key"`
    response string
}

func GetInstance() *gorm.DB {
    db, err := gorm.Open("sqlite3", "test.db")
    utils.CheckError(err)
    return db
}
