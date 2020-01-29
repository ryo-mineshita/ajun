package controller

import (
	"net/http"
	// "image/png"
	"reactgo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/boombuler/barcode"
	// "github.com/boombuler/barcode/qr"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) GetFoodPage(c *gin.Context) {
	errMsg := ""
	foods, err := model.GetAllFoods(s.DB)
	if err != nil {
		errMsg = "エラー発生"
		foods = []model.Food{}
	}
	c.JSON(http.StatusOK,gin.H{
		"foods": foods,
		"errMsg": &errMsg,
	})
}

func (s *Server) CreateFoodHandler(c *gin.Context) {
	food := new(model.Food)
	c.BindJSON(&food)
	model.CreateFood(s.DB, food)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (s *Server) GetShopPage(c *gin.Context) {
	errMsg := ""
	shops, err := model.GetAllShops(s.DB)
	if err != nil {
		errMsg = "エラー発生"
		shops = []model.Shop{}
	}
	c.JSON(http.StatusOK,gin.H{
		"shops": shops,
		"errMsg": &errMsg,
	})
}

func (s *Server) CreateShopHandler(c *gin.Context) {
	shop := new(model.Shop)
	c.BindJSON(&shop)
	model.CreateShop(s.DB, shop)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (s *Server) CreateUserHandler(c *gin.Context) {
	user := new(model.User)
	c.BindJSON(&user)
	user2 , _ ,hantei := model.CreateUser(s.DB, user)
	if(hantei == true){
		c.JSON(http.StatusOK,gin.H{
			"login": hantei,
			"user" : user2,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (s *Server) UserShopHandler(c *gin.Context) {
	user := new(model.User)
	shop := new(model.Shop)
	c.BindJSON(&user)

	shops , _ ,hantei := model.UserShop(s.DB, user, shop)
	if(hantei == true){
		c.JSON(http.StatusOK,gin.H{
			"hantei": hantei,
			"shops" : shops,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func (s *Server) ShopFoodHandler(c *gin.Context) {
	food := new(model.Food)
	shop := new(model.Shop)
	c.BindJSON(&shop)

	foods , _ ,hantei := model.ShopFood(s.DB, food, shop)
	if(hantei == true){
		c.JSON(http.StatusOK,gin.H{
			"hantei": hantei,
			"foods" : foods,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
