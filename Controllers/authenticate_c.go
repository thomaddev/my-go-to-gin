package controllers

import (
	"net/http"

	"my-go-to-bed/models"

	"github.com/gin-gonic/gin"
)

func (db *DBController) GetUser(c *gin.Context) {
	_type := c.Query("type")
	_where := map[string]interface{}{}

	if _type != "" {
		_where["type"] = _type
	}

	var collections []models.User
	db.Database.Where(_where).Find(&collections)

	// for i, _ := range collections {
	// 	db.Database.Model(collections[i]).Association("Groups").Find(&collections[i].Groups)
	// }

	c.JSON(http.StatusOK, gin.H{"results": &collections})
}

func (db *DBController) AddUser(c *gin.Context) {
	// var user models.User
	// user.Name = "test1"
	// user.UserName = "test_username"

	user := models.User{Id: 1, Name: "Jinzhu", UserName: "test_username"}

	result := db.Database.Create(&user)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"results": "SUCCESS"})
	}

	// if result := db.Database.Create(&user); result.Error != nil {
    //     c.AbortWithError(http.StatusNotFound, result.Error)
    //     return
    // }
}