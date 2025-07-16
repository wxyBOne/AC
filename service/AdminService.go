package service

import (
	"ac/dao"
	"ac/model"
)

// LoginAdminService 管理员登录服务
func LoginAdminService(adminname string, password string) (uint, error) {
	return dao.LoginAdminDao(adminname, password)
}

// GetAllUsersService 获取所有用户
func GetAllUsersService() ([]model.User, error) {
	return dao.GetAllUsers()
}

// UpdateUserByAdminService 管理员更新用户信息
func UpdateUserByAdminService(uid uint, username string, password string) error {
	return dao.UpdateUserByAdmin(uid, username, password)
}

// DeleteUserService 删除用户
func DeleteUserService(uid uint) error {
	return dao.DeleteUser(uid)
}

// SearchUsersService 搜索用户
func SearchUsersService(query string) ([]model.User, error) {
	return dao.SearchUsers(query)
}

// GetPendingArtworksService 获取待审核艺术品
func GetPendingArtworksService() ([]model.ArtReviewItem, error) {
	return dao.GetPendingArtworks()
}

// ReviewArtworkService 审核艺术品
func ReviewArtworkService(artID uint, auditorID uint, status int, contractHash string) error {
	return dao.CreateReviewRecord(artID, auditorID, status, contractHash)
}

// GetTransactionCountService 获取交易总数
func GetTransactionCountService() (int64, error) {
	return dao.GetTransactionCount()
}

// GetTransactionTotalAmountService 获取已完成交易的总金额
func GetTransactionTotalAmountService() (float64, error) {
	return dao.GetTransactionTotalAmount()
}

// GetRecentTransactionsService 获取最近交易(最多10条)
func GetRecentTransactionsService() ([]model.Transaction, error) {
	return dao.GetRecentTransactions()
}
