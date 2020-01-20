package model

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm.Model
	Storename   string `json:"storename"`
	Loc         string `json:"loc"`
	Genre       string `json:"ganre"`
	Tel         string `json:"tel"`
	Information string `json:"information"`
	Tables      string `json:"tables"`
}

func GetAllStores(db *gorm.DB) ([]Store, error) {
	stores := []Store{}
	if err := db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}
