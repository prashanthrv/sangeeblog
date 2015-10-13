package models

import (
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"time"
	"fmt"
	_ "github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
	"path"
)

var ReturnData struct {
	Posts      []Post
	Categories []Category
	Pages      []Page
}

func GetDB(filename string) *gorm.DB {
	db, err := gorm.Open("sqlite3", path.Join("data", filename+".db"))
	if err != nil {
		panic(err)
	}
	db.DB()

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// db.CreateTable(&User{})
	// db.CreateTable(&Address{})
	// db.CreateTable(&Email{})
	// db.CreateTable(&Language{})
	db.CreateTable(&Post{})
	db.CreateTable(&Page{})
	db.CreateTable(&Category{})
	//  category := Category{CategoryName: "Stories"}
	//  //db.First(&category, 5)
	//  page := Page{PageName: "Scribbles"}
	//  //db.First(&page, 5)
	//  post := Post{Author:"Sangee" ,PostTitle: "The War of the Two", Category: category, Content: "\"Just stop analyzing everything!!\", screamed the heart. \"Will you please stop getting emotional all the time\"? the mind screamed back. \"Oh please, without emotions, you'll be as good as dead!\", countered the heart. \"You know what, I would prefer being dead over fixating on this rubbish you call love\", smirked the mind. \"You cloud your thoughts, wondering too much about how a simple emotion can harm your day!\". \"Sweety, the mind sometimes just gets clouded. The heart gets overloaded, darkened and heavy. Which do you think is worse for the guy we belong to?\". And so on it went..<br /><br /><br /><br />\"Poor fellow, look at his state. He has just been letting us argue for hours!\". \"Aww, the poor heart has started feeling bad now\". \"Shut up, you are supposed to think the best for him. And all you do is figure out how to keep him from finding true happiness\". \"Ya right, all I am trying to do is protect him from getting hurt. Actually, did you realise I am protecting you? You might have forgotten, but I saw how you became after his last breakup\". \"What on earth is wrong with you?! I don't remember you telling the leg not to run too fast, although you knew he was in danger of falling down. Didn't the leg get hurt then? And didn't it move on, just the way everything and everyone in this world is supposed to?\".<br /><br />Really? Is that how it is really supposed to be? I found myself wondering. Actually if we think about it, that is exactly what we have been doing. Making our own mistakes, realising and regretting at a later stage, all the while convincing ourselves that time is the answer to all problems, time heals all wounds and that age makes us grow wiser. There's no questioning the logic of this life - it is obvious that with time, everything really does change. But then, if mistakes are what we are meant to do..is life then simply a continuous process of moving on from something or the other?<br /><br />\"Human beings have brains & a sixth sense for a reason, not just to blindly make mistakes and spend the rest of their lives moving on from them..\" my mind whispered gently. And my heart had an answer \"They spend their whole lives chasing happiness. If something gives them happiness for even a brief period of time, it cannot be termed a mistake\".<br /><br />True again. After all, aren't we all told to just \"live\" life, taking it one day at a time? As long as I do not become a mass murderer or a psychopath, I think I am perfectly allowed to make my own mistakes! Especially if they are mistakes in love. Don't we always find happiness in sharing love, whether it's with a parent or a sibling or a friend or a soulmate or even a pet? Sometimes things go wrong, we fall out with each other and suffer in loneliness. Otherwise, the period of happiness extends longer, only to be cut short by this thing called death. In the end, what's the point, really?<br /><br />\"Exactly. There's no point in the end!\", my heart exclaimed. My mind was at a loss for words. And my heart found its space to go on. \"Go on, tell her you love her. If things work out, good for you. Enjoy your time together, whether its days or months or years. Just make sure that when it ends, for whatever reason, what you had with her stays a lovely memory, always safe with me.\" \"I guess I have a big role to play in that\", my mind confessed and continued, \"Alright, I promise I won't think too much, I won't complicate your relationship and I will make sure your love for her stays above all else, including my ego.\"<br /><br />With my heart and mind finally at ease, I knocked her door. She opened it, smiling radiantly, looking splendid in her blue evening gown and long curly hair, with her dark black eyes looking deep into mine. For just a second, they showed a slight hint of confusion, as her gaze lowered, as I went down on my knees.<br /><br />The best memory I have of that day now is the blazing love I saw in her eyes, burning fiercely, just for me.<br /><br />Love truly is a magical feeling. And I won't mind living life a million times over, if I have someone I can fall in love with, every single day, over and over again.", Page: page, PostCreated:time.Now()}
  //  fmt.Println("Post Created")
	//  db.Create(&post)
	return &db
}
