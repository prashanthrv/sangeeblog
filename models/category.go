package models

import (
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/jinzhu/gorm"
	_ "github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
)

type Category struct {
	gorm.Model
	CategoryName string
}

func GetCategories(db *gorm.DB) interface{} {
	db.Find(&ReturnData.Categories)
	return ReturnData.Categories
}
