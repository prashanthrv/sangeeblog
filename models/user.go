package models

import (
	"database/sql"
	//"fmt"
	//"log"

	//"golang.org/x/crypto/bcrypt"
	//"gopkg.in/gorp.v1"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/golang/glog"
	"time"
)

type User struct {
    ID           int
    Birthday     time.Time
    Age          int
    Name         string  `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time

    Emails            []Email         // One-To-Many relationship (has many)
    BillingAddress    Address         // One-To-One relationship (has one)
    BillingAddressID  sql.NullInt64   // Foreign key of BillingAddress
    ShippingAddress   Address         // One-To-One relationship (has one)
    ShippingAddressID int             // Foreign key of ShippingAddress
    IgnoreMe          int `sql:"-"`   // Ignore this field
    Languages         []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

type Email struct {
    ID      int
    UserID  int     `sql:"index"` // Foreign key (belongs to), tag `index` will create index for this field when using AutoMigrate
    Email   string  `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
    Subscribed bool
}

type Address struct {
    ID       int
    Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
    //Address2 string         `sql:"type:varchar(100);unique"`
    //Post     sql.NullString `sql:"not null"`
}

type Language struct {
    ID   int
    Name string `sql:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
    Code string `sql:"index:idx_name_code"` // `unique_index` also works
}

// func (user *User) HashPassword(password string) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		glog.Fatalf("Couldn't hash password: %v", err)
// 		panic(err)
// 	}
// 	user.Password = hash
// }

// func GetUserByEmail(dbMap *gorp.DbMap, email string) (user *User) {
// 	err := dbMap.SelectOne(&user, "SELECT * FROM Users where Email = ?", email)
//
// 	if err != nil {
// 		glog.Warningf("Can't get user by email: %v", err)
// 	}
// 	return
// }
//
// func InsertUser(dbMap *gorp.DbMap, user *User) error {
// 	return dbMap.Insert(user)
// }
//
// func GetDbMap(user, password, hostname, port, database string) *gorp.DbMap {
// 	// connect to db using standard Go database/sql API
// 	// use whatever database/sql driver you wish

// 	db, err := sql.Open("mysql", fmt.Sprint(user, ":", password, "@(", hostname, ":", port, ")/", database, "?charset=utf8mb4"))
// 	checkErr(err, "sql.Open failed")
//
// 	// construct a gorp DbMap
// 	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}
//
// 	// add a table, setting the table name to 'posts' and
// 	// specifying that the Id property is an auto incrementing PK
// 	dbMap.AddTableWithName(User{}, "Users").SetKeys(true, "Id")
//
// 	// create the table. in a production system you'd generally
// 	// use a migration tool, or create the tables via scripts
// 	err = dbMap.CreateTablesIfNotExists()
// 	checkErr(err, "Create tables failed")
//
// 	return dbMap
// }
//
// func checkErr(err error, msg string) {
// 	if err != nil {
// 		log.Fatalln(msg, err)
// 	}
// }
