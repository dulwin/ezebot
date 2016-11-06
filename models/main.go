package models

import (
    "github.com/jinzhu/gorm"
)

type Query struct {
    gorm.Model
    Category string `gorm:"size:127;primary_key"`
    Query    string `gorm:"size:127;primary_key"`
    Response string
}
