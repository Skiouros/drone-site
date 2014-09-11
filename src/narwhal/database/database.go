package database

import "log"
import "encoding/gob"
import "github.com/jinzhu/gorm"
import _ "github.com/lib/pq"

var schemas = [...]string {
	`CREATE TABLE IF NOT EXISTS "users"
	(
		id integer NOT NULL PRIMARY KEY,
		name text,
		pass text,
		UNIQUE(name)
	);`,
}

var DbMap *gorm.DB

func init() {
	DbMap = openDB()

	DbMap.Table("users").CreateTable(&User{})
	DbMap.Table("videos").CreateTable(&Video{})

	gob.Register(&User{})
	gob.Register(&Video{})
}

func openDB() (*gorm.DB) {
	DbMap, err := gorm.Open("postgres", "postgres://docker:docker@172.17.0.10:5432/docker?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	DbMap.LogMode(true)

	return &DbMap
}

func InitDB() {
	// for _, schema := range schemas {
	// 	_, err := DbMap.DB().Exec(schema)
	// 	if err != nil {
	// 		log.Printf("%q: %s\n", err, schema)
	// 		return
	// 	}
	// }
}

