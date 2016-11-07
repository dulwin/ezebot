package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dulwin/ezebot/models"
	"github.com/dulwin/ezebot/utils"
)

type EntityManager struct {
	instance *gorm.DB
}

func GetInstance() *EntityManager {
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open("sqlite3", dbName)
	utils.CheckError(err)
	return &EntityManager{instance: db}
}

func (e *EntityManager) Close() {
	e.instance.Close()
}

func (e *EntityManager) Migrate() {
	e.instance.AutoMigrate(&models.Query{})
}

func (e *EntityManager) Insert(q models.Entity) {
	e.instance.Create(q)
}
