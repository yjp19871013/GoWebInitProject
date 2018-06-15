package models

import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
	"book_borrow/config"
)

var db *gorm.DB

func init() {
	fmt.Println("Start database init")
	
	var err error
	db, err = gorm.Open("mysql", config.DB_USERNAME + ":" + config.DB_PASSWORD + 
		"@tcp(" + config.DB_HOST + ")/" + config.DB_SCHEMA + 
		"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&User{}, &Role{}, &Permission{})
	
	initTable()
	
	fmt.Println("End database init")
}

func initTable() {
	initPermissions()
	fmt.Println("permissions table inited")
	
	initRoles()
	fmt.Println("roles table inited")
	
	initAdminUser()
	fmt.Println("users table inited")
}

