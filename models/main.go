package models

import (
    "github.com/jinzhu/gorm"
)

type Entity interface {

}

type Query struct {
    gorm.Model
    Entity

    Category string `gorm:"size:127;primary_key"`
    Query    string `gorm:"size:127;primary_key"`
    Response string
}
