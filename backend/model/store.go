package model

import (
	"github.com/jinzhu/gorm"
)

type Food struct {
	gorm.Model
	Foodid     	int	   `json:"foodid"`
	Shopid     	int    `json:"shopid"`
	Foodname   	string `json:"foodname"`
	Foodsize    string `json:"foodsize"`
	Foodsummary	string `json:"foodsummary"`
	Foodgenre	  string `json:"foodgenre"`
	Foodimage	  string `json:"foodimage"`
	Foodprice	  int    `json:"foodprice"`
}

type Shop struct {
	gorm.Model
	Userid      int    `json:"userid"`
	Shopid     	int    `json:"shopid"`
	Shopname   	string `json:"shopname"`
	Shopsummary string `json:"shopsummary"`
	Shoptel			string `json:"shoptel"`
	Shoptable	  string `json:"shoptable"`
	Shopaddress	string `json:"shopaddress"`
	Shopgenre	  string `json:"shopgenre"`
}

type User struct {
	Userid			 int 		`json:"userid"`
  Userpass		 string `json:"userpass"`
  Username	 	 string `json:"username"`
}

func GetAllFoods(db *gorm.DB) ([]Food, error) {
	foods := []Food{}
	if err := db.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func CreateFood(db *gorm.DB, food *Food) error {
	return db.Create(&food).Error
}

func GetAllShops(db *gorm.DB) ([]Shop, error) {
	shops := []Shop{}
	if err := db.Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func CreateShop(db *gorm.DB, shop *Shop) error {
	return db.Create(&shop).Error
}

func CreateUser(db *gorm.DB,user *User) ([]User,error,bool) {
	users := []User{}
	a			:= user.Userid
	b     := user.Userpass
  var   c bool
	var count = 0

	if err := db.Where("userid = ? AND userpass = ?", a,b).Find(&users).Count(&count).Error; err != nil {
		c = false
		return nil, err,c
	}

	if count == 0 {
        c = false
    } else {
        c = true
    }

	return users,nil,c
}

func UserShop(db *gorm.DB,user *User,shop *Shop) ([]Shop,error,bool) {
	shops := []Shop{}
	a			:= user.Userid
  var   c bool
	var count = 0

	if err := db.Where("userid = ?", a).Find(&shops).Count(&count).Error; err != nil {
		c = false
		return nil, err,c
	}

	if count == 0 {
        c = false
    } else {
        c = true
    }

	return shops,nil,c
}
