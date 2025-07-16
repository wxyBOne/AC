package dao

import (
	"ac/model"
	"ac/utils"
	"errors"
)

// LoginAdminDao 管理员登录
func LoginAdminDao(adminname string, password string) (uint, error) {
	db := utils.DB
	var admin model.Admin

	result := db.Table("administrators").Where("adminname = ?", adminname).First(&admin)
	if result.Error != nil {
		return 0, errors.New("管理员不存在")
	}

	if admin.Adminpass != password {
		return 0, errors.New("密码错误")
	}

	return admin.AdminID, nil
}

// GetAllUsers 获取所有用户信息
func GetAllUsers() ([]model.User, error) {
	db := utils.DB
	var users []model.User
	result := db.Table("users").Find(&users)
	return users, result.Error
}

// UpdateUser 更新用户信息
func UpdateUserByAdmin(uid uint, username string, password string) error {
	db := utils.DB

	// 检查用户名是否已存在（排除当前用户）
	var existingUser model.User
	if err := db.Table("users").Where("username = ? AND uid != ?", username, uid).First(&existingUser).Error; err == nil {
		return errors.New("用户名已存在")
	}

	// 更新用户信息
	result := db.Table("users").Where("uid = ?", uid).Updates(map[string]interface{}{
		"username": username,
		"userpass": password,
	})
	return result.Error
}

// DeleteUser 删除用户
func DeleteUser(uid uint) error {
	db := utils.DB
	result := db.Table("users").Where("uid = ?", uid).Delete(&model.User{})
	return result.Error
}

// SearchUsers 搜索用户
func SearchUsers(query string) ([]model.User, error) {
	db := utils.DB
	var users []model.User
	result := db.Table("users").Where("username LIKE ? OR uid LIKE ?",
		"%"+query+"%", "%"+query+"%").Find(&users)
	return users, result.Error
}

// GetPendingArtworks 获取待审核的艺术品
func GetPendingArtworks() ([]model.ArtReviewItem, error) {
	db := utils.DB
	var artPieces []model.ArtPiece
	var reviewItems []model.ArtReviewItem

	// 获取未审核的艺术品（通过左连接查询没有审核记录的艺术品）
	if err := db.Table("art_pieces").
		Select("art_pieces.*, review_records.status").
		Joins("LEFT JOIN review_records ON art_pieces.aid = review_records.art_id").
		Where("review_records.rid IS NULL").
		Find(&artPieces).Error; err != nil {
		return nil, err
	}

	// 转换为前端展示模型
	for _, piece := range artPieces {
		creatorName, _ := GetUsernameByUserID(piece.CreatorID)
		reviewItem := model.ArtReviewItem{
			Aid:         piece.Aid,
			Title:       piece.Title,
			Description: piece.Description,
			Tag:         piece.Tag,
			ArtImage:    piece.ArtImage,
			ArtPrice:    piece.ArtPrice,
			CreatorName: creatorName,
		}
		reviewItems = append(reviewItems, reviewItem)
	}

	return reviewItems, nil
}

// CreateReviewRecord 创建审核记录
func CreateReviewRecord(artID uint, auditorID uint, status int, contractHash string) error {
	db := utils.DB
	review := model.ReviewRecord{
		ArtID:     artID,
		AuditorID: auditorID,
		Status:    status,
	}

	//更新艺术品合约标识
	var artPieces model.ArtPiece
	result := db.Table("art_pieces").Where("aid = ?", artID).First(&artPieces)
	if result.Error != nil {
		return result.Error
	}
	artPieces.ContractHash = contractHash
	if err := db.Table("art_pieces").Where("aid = ?", artID).Updates(&artPieces).Error; err != nil {
		return err // 更新失败
	}

	return db.Table("review_records").Create(&review).Error
}

// GetTransactionCountService 获取交易总数
func GetTransactionCount() (int64, error) {
	db := utils.DB
	var count int64
	result := db.Table("transactions").Count(&count)
	return count, result.Error
}

// GetTransactionTotalAmount 获取已完成交易的总金额
func GetTransactionTotalAmount() (float64, error) {
	db := utils.DB
	var result struct {
		TotalAmount float64 `json:"total_amount"`
	}

	err := db.Table("transactions").
		Where("is_completed = ?", 1).
		Select("COALESCE(SUM(price), 0) as total_amount"). // 表中没有符合条件的记录时，返回0
		Scan(&result).Error

	return result.TotalAmount, err
}

// GetRecentTransactions 获取最近交易(最多10条)
func GetRecentTransactions() ([]model.Transaction, error) {
	db := utils.DB
	var transactions []model.Transaction

	// 先获取总记录数
	var count int64
	db.Table("transactions").Where("is_completed = ?", 1).Count(&count)

	// 如果没有记录，返回空切片
	if count == 0 {
		return []model.Transaction{}, nil
	}

	// 限制获取最近10条记录
	limit := 10
	if count < int64(limit) {
		limit = int(count)
	}

	result := db.Table("transactions").
		Where("is_completed = ?", 1).
		Order("created_at desc").
		Limit(limit).
		Find(&transactions)

	return transactions, result.Error
}
