package routers

import (
	"my-go-to-bed/controllers" // import controllers package

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetCollectionRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctrls := controllers.DBController{Database: db}

	router.GET("users", ctrls.GetUser) // GET
	router.GET("user/add", ctrls.AddUser) // GET
	// router.GET("collections/:id", ctrls.GetCollectionById)  // GET BY ID
	// router.POST("collections", ctrls.CreateCollection)  // POST
	// router.PATCH("collections", ctrls.UpdateCollection)  // PATCH
	// router.DELETE("collections/:id", ctrls.DeleteCollection)  // DELETE
}
