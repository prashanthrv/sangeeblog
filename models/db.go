package models

import (
	"github.com/prashanthrv/sangeeblog/Godeps/_workspace/src/github.com/jinzhu/gorm"
	//"time"
	//"fmt"
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
	//  category := Category{}
	//  db.First(&category, 5)
	//  page := Page{}
	//  db.First(&page, 5)
	//  post := Post{Author:"Sangee" ,PostTitle: "Golu - the celebration of Navratri", Category: category, Content: "It was that time of the year again when her mother would call her every evening to say that she misses her. And as can be expected from a traditional Indian mother, it had something to do with food. As her mother continued raving on about the delicious sundal she had made, her mind drifted to memories of the past.<br /><br />The young girl was excited. It was Navratri season again. As usual, she began screaming in her sister’s ears and shaking her to wake her up. Running away just in time to escape from her wrath, she went next to her father, tugging at his shirt, calling him to take down the cardboards of dolls from the loft. It was time for ‘Golu’, the South Indian traditional festival of dolls. At her home, this was a ritual which happened year after year – as her mom and dad set up the steps of the ‘Golu’, she and her sister would painstakingly remove all the newspaper coverings of each doll safely preserved up in the lofts for a whole year. It was always a race between them to search for the tiny birthday cake, which was a part of the birthday doll set. The joy seen in the one who found the cake amidst all the clutter had no limits.<br /><br />'Childhood is truly a magical phase', she reminisced. Her mother's shriek startled her back to reality. She started regretting for accepting the video call, as her mother moved on to the ‘You have lost so much weight’, 'You are not eating at all’ monologue, commonly heard by every son/daughter living away from home. To divert the topic, she asked to see that year’s Golu at her place on video. After her initial shock on seeing the small cupboard Golu, she began to understand how alone her mother really was.<br /><br />Golu at her place had always been a magnificent affair, with 9 colorful steps full of huge dolls along with sprawling zoos, parks, processions & what not, spread around on either side of the main Golu. The whole family used to spend hours together on setting it all up, showcasing various levels of creativity in making each year’s output look different. The neighborhood stationery store enjoyed high sales during this festival - colored chart papers, thermocol sheets & balls were purchased in great vigor to create roads & borders for their dream doll city. The cricket set put forth a never-ending challenge; it was fun to design new world-class stadiums every year for the wonderful Plaster-of-Paris players. In addition to these, her mother used Golu as an opportunity to do her part for the society – she created her own humble awareness-creating social campaign every year, right from within the walls of her cherished home. Being the perfect host she was, she loved attending to guests on all 9 days of the festival. The house turned lively every evening, with the ragas of divine Carnatic music echoing through the rooms and the heavenly aroma of tasty sundal mesmerizing every visitor. Nobody ever left empty-handed, her mother herself handpicked small memorable gifts to be given to each guest during those 9 days.<br /><br />As the memories filled her mind, she started counting back to when she had last experienced the joy of Navratri Golu. It had been 5 years since she had left home, in pursuit of good education and new places. 5 years since she had searched for the little cake with her sister; her sister had since gotten married and had now started keeping her own Golu at her in-laws’ place. 5 years since she had happily worn colorful ‘pattu pavadais’ and visited relatives’ houses, meeting cousins and mocking the elders by singing movie songs at the Golu. It had been 5 years, and as she had moved on with her life, attending Dandiya festivals in the north, her mother had been left behind. With both daughters away, she had lost her enthusiasm in the festivities. Now, arranging the dolls had become a burden, all bursts of creativity had been long gone and she made sundals in the forlorn kitchen, thinking of her daughter miles away who loved them.", Page: page, PostCreated:time.Now()}
	//  db.Create(&post)
	return &db
}
