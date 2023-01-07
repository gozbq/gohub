package user

import "gohub/pkg/database"

// IsEmailExist 判断Email已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ? ", phone).First(&userModel)
	return
}
func GetByMulti(loginID string) (userModel User) {
	database.DB.Where("phone = ? ", loginID).Or("email = ? ", loginID).Or("name= ?", loginID).First(&userModel)
	return
}
