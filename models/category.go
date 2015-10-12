package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

type Category struct {
  gorm.Model
	CategoryName    string
}

func GetCategories(db *gorm.DB) interface{}{
  db.Find(&ReturnData.Categories)
  return ReturnData.Categories
}
