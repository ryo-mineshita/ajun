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
	kensakuid	  := user.Userid
	kensakupass := user.Userpass
  var hantei bool
	var count = 0

	if err := db.Where("userid = ? AND userpass = ?", kensakuid,kensakupass).Find(&users).Count(&count).Error; err != nil {
		hantei = false
		return nil, err,hantei
	}

	if count == 0 {
        hantei = false
    } else {
        hantei = true
    }

	return users,nil,hantei
}

func UserShop(db *gorm.DB,user *User,shop *Shop) ([]Shop,error,bool) {
	shops := []Shop{}
	kensaku := user.Userid
  var hantei bool
	var count = 0

	if err := db.Where("userid = ?", kensaku).Find(&shops).Count(&count).Error; err != nil {
		hantei = false
		return nil, err,hantei
	}

	if count == 0 {
        hantei = false
    } else {
        hantei = true
    }

	return shops,nil,hantei
}

func ShopFood(db *gorm.DB,food *Food,shop *Shop) ([]Food,error,bool) {
	foods := []Food{}
	kensaku := shop.Shopid
  var hantei bool
	var count = 0

	if err := db.Where("shopid = ?", kensaku).Find(&foods).Count(&count).Error; err != nil {
		hantei = false
		return nil, err,hantei
	}

	if count == 0 {
        hantei = false
    } else {
        hantei = true
    }

	return foods,nil,hantei
}
