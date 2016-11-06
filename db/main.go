package db

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dulwin/ezebot/utils"
    "github.com/dulwin/ezebot/models"
)

type EntityManager struct {
    manager *gorm.DB
}

func GetInstance() EntityManager {
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open("sqlite3", dbName)
	utils.CheckError(err)
	return EntityManager{manager: db}
}

func (e EntityManager) Close() {
    e.manager.Close()
}

func (e EntityManager) Migrate() {
	e.manager.AutoMigrate(&models.Query{})
}

func (e EntityManager) Insert(q models.Entity) {
	e.manager.Create(q)
}
