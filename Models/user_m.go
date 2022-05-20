package models
import "github.com/jinzhu/gorm"
type User struct {
	gorm.Model
	Id       uint   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Name     string `json:"name" orm:"size(128)"`
	UserName string `json:"user_name" orm:"size(128)"`
}
