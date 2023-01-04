package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

// User用户模型
type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
