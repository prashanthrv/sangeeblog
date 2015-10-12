package models

import (
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"time"
	//"fmt"
	_ "github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
	//"strconv"
)

type Post struct {
	gorm.Model
	Author      string
	PostTitle   string
	Category    Category
	CategoryID  int
	Content     string
	Page        Page
	PageID      int
	PostCreated time.Time
}

func GetPosts(db *gorm.DB) interface{} {
	db.Preload("Category").Find(&ReturnData.Posts)
	return ReturnData.Posts
}

func GetPost(db *gorm.DB, postid string) interface{} {
	post := Post{}
	db.Preload("Category").Find(&post, "id =?", postid)
	return post
}

func GetPostsByID(db *gorm.DB, categoryid string) interface{} {
	db.Preload("Category", "ID = ?", categoryid).Find(&ReturnData.Posts, "category_id =?", categoryid)
	return ReturnData.Posts
}

func GetPostsByPage(db *gorm.DB, pageid string) interface{} {
	db.Preload("Category").Find(&ReturnData.Posts, "page_id =?", pageid)
	return ReturnData.Posts
}
