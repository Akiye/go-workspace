package models

import (
	"go-workspace/crud-api/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	Email        string `json:"email"`
	Title        string `json:"title"`
	DisplayName  string `json:"display_name"`
	Height       int    `json:"height"`
	Weight       int    `json:"weight"`
	BodyFat      int    `json:"body_fat"`
	TargetWeight int    `json:"target_weight"`
	TestDate     string `json:"test_date"`
	NextTestDate string `json:"next_test_date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID = ?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID = ?", ID).Delete(user)
	return user
}
