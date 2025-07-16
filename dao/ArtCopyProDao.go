package dao

import (
	"ac/model"
	"ac/utils"
	"fmt"
)

// GetArtItems 获取已审核通过且未售出的艺术品列表
func GetArtItem() ([]model.ArtItem, error) {
	db := utils.DB
	var artPieces []model.ArtPiece

	// 使用JOIN查询获取已审核通过且未售出的艺术品
	if err := db.Table("art_pieces").
		Joins("JOIN review_records ON art_pieces.aid = review_records.art_id").
		Where("art_pieces.is_sold = ? AND review_records.status = ?", 0, 1).
		Find(&artPieces).Error; err != nil {
		return nil, err // 获取失败
	}

	var artItems []model.ArtItem
	for _, artPiece := range artPieces {
		creatorName, err := GetUsernameByUserID(artPiece.CreatorID)
		if err != nil {
			return nil, err
		}

		artItem := model.ArtItem{
			Aid:          artPiece.Aid,
			Title:        artPiece.Title,
			Description:  artPiece.Description,
			Tag:          artPiece.Tag,
			ArtImage:     artPiece.ArtImage,
			ArtPrice:     artPiece.ArtPrice,
			CreatorName:  creatorName,
			ContractHash: artPiece.ContractHash,
		}
		artItems = append(artItems, artItem)
	}

	return artItems, nil // 返回艺术品列表
}

// AddArtPiece 添加艺术品
func AddArtPiece(artPiece model.ArtPiece) (uint, error) {
	db := utils.DB
	result := db.Table("art_pieces").Create(&artPiece)
	if result.Error != nil {
		return 0, result.Error // 添加失败
	}
	return artPiece.Aid, nil // 返回新艺术品的 ID
}

// GetArtProByID 根据 ID 获取艺术品
func GetArtProByID(aid uint) (model.ArtItem, error) {
	db := utils.DB
	var artPiece model.ArtPiece
	var artPro model.ArtItem

	if err := db.Table("art_pieces").Where("aid = ?", aid).First(&artPiece).Error; err != nil {
		return artPro, err // 获取失败
	}

	creatorName, err := GetUsernameByUserID(artPiece.CreatorID)
	if err != nil {
		return artPro, err
	}

	artPro = model.ArtItem{
		Aid:          artPiece.Aid,
		Title:        artPiece.Title,
		Description:  artPiece.Description,
		Tag:          artPiece.Tag,
		ArtImage:     artPiece.ArtImage,
		ArtPrice:     artPiece.ArtPrice,
		CreatorName:  creatorName,
		ContractHash: artPiece.ContractHash,
	}

	return artPro, nil // 返回艺术品列表
}

// GetArtProByID 根据 ID 获取艺术品
func GetArtPieceByID(aid uint) (model.ArtPiece, error) {
	db := utils.DB
	var artPiece model.ArtPiece

	if err := db.Table("art_pieces").Where("aid = ?", aid).First(&artPiece).Error; err != nil {
		return artPiece, err // 获取失败
	}

	return artPiece, nil // 返回艺术品列表
}

// AddToCart 添加商品到购物车
func AddToCart(userID uint, artID uint) error {
	db := utils.DB
	cartItem := model.ShoppingCart{
		UserID: userID,
		ArtID:  artID,
	}

	// 尝试插入，如果已存在则返回错误
	if err := db.Table("shopping_carts").Create(&cartItem).Error; err != nil {
		return fmt.Errorf("item already in cart") // 返回自定义错误
	}

	return nil // 成功
}

// GetArtPiecesByIDs 根据 ArtID 列表获取艺术品
func GetArtPiecesByIDs(artIDs []uint) ([]model.ArtPiece, error) {
	db := utils.DB
	var artPieces []model.ArtPiece

	// 使用 ? 占位符时，确保传递的是切片
	if err := db.Table("art_pieces").Where("aid IN (?)", artIDs).Find(&artPieces).Error; err != nil {
		return nil, err // 获取失败
	}

	return artPieces, nil // 返回艺术品列表
}

// GetCartItems 获取购物车商品
func GetCartItems(userID uint) ([]model.ShoppingCart, error) {
	db := utils.DB
	var cartItems []model.ShoppingCart

	if err := db.Table("shopping_carts").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return nil, err // 获取失败
	}

	return cartItems, nil // 返回购物车商品
}

// RemoveFromCart 删除购物车商品
func RemoveFromCart(artID uint, userID uint) error {
	db := utils.DB
	if err := db.Table("shopping_carts").Where("art_id = ? AND user_id = ?", artID, userID).Delete(&model.ShoppingCart{}).Error; err != nil {
		return err // 删除失败
	}
	return nil // 成功
}

// RemoveCartPiece 删除购物车中的商品
func RemoveCartPiece(artID uint) error {
	db := utils.DB
	if err := db.Table("shopping_carts").Where("art_id = ?", artID).Delete(&model.ShoppingCart{}).Error; err != nil {
		return err // 删除失败
	}

	return nil // 成功
}

// AddTransaction 添加交易记录
func AddTransaction(transaction model.Transaction) (model.Transaction, error) {
	db := utils.DB
	if err := db.Table("transactions").Create(transaction).Error; err != nil {
		return transaction, err // 添加失败
	}
	//获取生成的tid
	var newTransaction model.Transaction
	if err := db.Table("transactions").Where("art_id = ?", transaction.ArtID).First(&newTransaction).Error; err != nil {
		return transaction, err // 获取失败
	}

	return newTransaction, nil // 成功
}

// AddOrder 添加订单记录
func AddOrder(order model.Order) error {
	db := utils.DB
	if err := db.Table("orders").Create(&order).Error; err != nil {
		return err // 添加失败
	}
	return nil // 成功
}

// UpdateArtPiece 更新艺术品
func UpdateArtPiece(artPiece model.ArtPiece) error {
	db := utils.DB
	if err := db.Table("art_pieces").Where("aid = ?", artPiece.Aid).Updates(artPiece).Error; err != nil {
		return err // 更新失败
	}
	return nil // 成功
}

// GetPublishedArtCount 获取已发布艺术品数量
func GetPublishedArtCount(userID uint) (int64, error) {
	db := utils.DB
	var count int64
	if err := db.Table("art_pieces").Joins("INNER JOIN review_records ON art_pieces.aid = review_records.art_id").
		Where("art_pieces.creator_id = ? AND review_records.status = ?", userID, 1).Count(&count).Error; err != nil {
		return 0, err // 获取失败
	}
	return count, nil // 返回数量
}

// GetPurchasedArtCount 获取已购入艺术品数量
func GetPurchasedArtCount(userID uint) (int64, error) {
	db := utils.DB
	var count int64
	if err := db.Table("transactions").Where("buyer_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err // 获取失败
	}
	return count, nil // 返回数量
}

// GetSoldArtCount 获取已售出艺术品数量
func GetSoldArtCount(userID uint) (int64, error) {
	db := utils.DB
	var count int64
	if err := db.Table("transactions").Where("seller_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err // 获取失败
	}
	return count, nil // 返回数量
}

// GetTransactionsByBuyerID 根据 buyer_id 获取交易记录
func GetTransactionsByBuyerID(buyerID uint) ([]model.Transaction, error) {
	db := utils.DB
	var transactions []model.Transaction
	if err := db.Table("transactions").Where("buyer_id = ?", buyerID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionsBySellerID 根据 seller_id 获取交易记录
func GetTransactionsBySellerID(sellerID uint) ([]model.Transaction, error) {
	db := utils.DB
	var transactions []model.Transaction
	if err := db.Table("transactions").Where("seller_id = ?", sellerID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetArtPiecesByCreatorID 根据 creator_id 获取艺术品
func GetArtPiecesByCreatorID(creatorID uint) ([]model.ArtPiece, error) {
	db := utils.DB
	var artPieces []model.ArtPiece
	// 使用连接查询获取已审核通过的艺术品
	if err := db.Table("art_pieces").
		Joins("INNER JOIN review_records ON art_pieces.aid = review_records.art_id").
		Where("art_pieces.creator_id = ? AND review_records.status = ?", creatorID, 1).
		Find(&artPieces).Error; err != nil {
		return nil, err
	}
	return artPieces, nil
}
