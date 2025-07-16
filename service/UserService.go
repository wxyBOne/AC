package service

import (
	"ac/dao"
)

// RegisterUserService 用户注册
func RegisterUserService(username string, password string, privateKey string, publicKey string) (uint, error) {
	return dao.RegisterUserDao(username, password, privateKey, publicKey)
}

// LoginUserService 用户登录
func LoginUserService(username string, password string) (uint, string, error) {
	return dao.LoginUserDao(username, password)
}

// ResetPasswordService 重置密码
func ResetPasswordService(privateKey string, newPassword string) error {
	return dao.ResetPasswordDao(privateKey, newPassword)
}

// UpdateUserService 更新用户信息
func UpdateUserService(currentUsername string, newUsername string, newPassword string) (uint, error) {
	return dao.UpdateUser(currentUsername, newUsername, newPassword)
}
