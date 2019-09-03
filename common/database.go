package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() *gorm.DB {
	var err error
	DB, err = gorm.Open("mysql", "root:@/naggi_club_mentoring?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	return DB
}

func TestDBInit() *gorm.DB {
	var err error
	DB, err = gorm.Open("mysql", "root:@/naggi_club_mentoring_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	return DB
}

func TestDBClean() {
	DB, err := gorm.Open("mysql", "root:@/naggi_club_mentoring_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	DB.Exec("DROP   DATABASE naggi_club_mentoring_test;")
	DB.Exec("CREATE DATABASE naggi_club_mentoring_test;")
}

func GetDB() *gorm.DB {
	return DB
}
