package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=20241202 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	var exists bool

	err := db.Model(&User{}).Select("count(*)>0").Where(User{Username: "admin", Password: "password"}).Find(&exists).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exists)
}
