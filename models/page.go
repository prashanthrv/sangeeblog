package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

type Page struct {
  gorm.Model
	PageName    string
}

func GetPages(db *gorm.DB) interface{}{
  db.Find(&ReturnData.Pages)
  return ReturnData.Pages
}
