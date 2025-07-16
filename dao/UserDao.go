package dao

import (
	"ac/model"
	"ac/utils"
	"errors"
)

// RegisterUserDao 使用gorm实现用户注册
func RegisterUserDao(username string, password string, privateKey string, publicKey string) (uint, error) {
	db := utils.DB

	// 检查用户名是否已存在
	var existingUser model.User
	if err := db.Table("users").Where("username = ?", username).First(&existingUser).Error; err == nil {
		return 0, errors.New("username already exists")
	}

	newUser := model.User{
		Username:   username,
		Userpass:   password,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
	result := db.Table("users").Create(&newUser)
	if result.Error != nil {
		return 0, result.Error
	}
	return newUser.Uid, nil
}

// LoginUserDao 使用gorm实现用户登录
func LoginUserDao(username string, password string) (uint, string, error) {
	db := utils.DB
	var user model.User
	result := db.Table("users").Where("username = ?", username).First(&user)
	if result.Error != nil {
		return 0, "", result.Error // 用户名未找到
	}
	if user.Userpass != password {
		return 0, "", errors.New("invalid password") // 密码错误
	}
	return user.Uid, user.PrivateKey, nil // 返回用户 ID 和私钥
}

// ResetPasswordDao 重置密码
func ResetPasswordDao(privateKey string, newPassword string) error {
	db := utils.DB
	var user model.User
	result := db.Table("users").Where("private_key = ?", privateKey).First(&user)
	if result.Error != nil {
		return result.Error
	}

	user.Userpass = newPassword
	if err := db.Table("users").Where("private_key = ?", privateKey).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户信息
func UpdateUser(currentUsername, newUsername, newPassword string) (uint, error) {
	db := utils.DB

	// 检查用户名是否已存在
	if newUsername != currentUsername {
		var existingUser model.User
		if err := db.Table("users").Where("username = ?", newUsername).First(&existingUser).Error; err == nil {
			return 1, errors.New("username already exists") // 用户名已存在
		}
	}

	var user model.User
	result := db.Table("users").Where("username = ?", currentUsername).First(&user)
	if result.Error != nil {
		return 2, result.Error // 用户名不存在
	}

	// 更新用户信息
	user.Username = newUsername
	user.Userpass = newPassword
	if err := db.Table("users").Where("username = ?", currentUsername).Updates(&user).Error; err != nil {
		return 3, err // 更新失败
	}

	return 0, nil // 更新成功
}

// 根据用户名获取用户ID
func GetUserIDByUsername(username string) (uint, error) {
	db := utils.DB

	var user model.User
	result := db.Table("users").Where("username =?", username).First(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Uid, nil
}

// 根据用户ID获取用户名
func GetUsernameByUserID(uid uint) (string, error) {
	db := utils.DB

	var user model.User
	result := db.Table("users").Where("uid =?", uid).First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Username, nil
}
