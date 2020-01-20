package controller

import (
	"net/http"

	// "image/png"
	"crud_go/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/boombuler/barcode"
	// "github.com/boombuler/barcode/qr"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) GetStorePage(c *gin.Context) {
	errMsg := ""
	stores, err := model.GetAllStores(s.DB)
	if err != nil {
		errMsg = "エラー発生"
		stores = []model.Store{}
	}
	c.JSON(http.StatusOK,gin.H{
		"stores": stores,
		"errMsg": &errMsg,
	})
}
