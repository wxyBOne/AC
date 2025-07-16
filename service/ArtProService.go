package service

import (
	"ac/dao"
	"ac/model"
)

// GetArtPiecesService 获取艺术品列表
func GetArtItemService() ([]model.ArtItem, error) {
	return dao.GetArtItem() // 调用 DAO 方法获取艺术品
}

// AddArtPieceService 添加艺术品
func AddArtPieceService(artPiece model.ArtPiece) (uint, error) {
	return dao.AddArtPiece(artPiece) // 调用 DAO 方法添加艺术品
}

// GetArtProByIDService 根据 ID 获取艺术品
func GetArtProByIDService(aid uint) (model.ArtItem, error) {
	return dao.GetArtProByID(aid) // 调用 DAO 方法获取艺术品
}

// GetArtProByIDService 根据 ID 获取艺术品
func GetArtPieceByIDService(aid uint) (model.ArtPiece, error) {
	return dao.GetArtPieceByID(aid) // 调用 DAO 方法获取艺术品
}

// AddToCartService 添加商品到购物车
func AddToCartService(userID uint, artID uint) error {
	return dao.AddToCart(userID, artID) // 调用 DAO 方法添加到购物车
}

// GetCartItemsService 获取购物车商品
func GetCartItemsService(userID uint) ([]model.ArtPiece, error) {
	cartItems, err := dao.GetCartItems(userID) // 获取购物车商品
	if err != nil {
		return nil, err
	}

	// 提取 ArtID 列表
	var artIDs []uint
	for _, item := range cartItems {
		artIDs = append(artIDs, item.ArtID)
	}

	// 根据 ArtID 获取艺术品
	artPieces, err := dao.GetArtPiecesByIDs(artIDs)
	if err != nil {
		return nil, err
	}

	return artPieces, nil // 返回艺术品列表
}

// RemoveFromCartService 删除购物车商品
func RemoveFromCartService(artID uint, userID uint) error {
	return dao.RemoveFromCart(artID, userID) // 调用 DAO 方法删除购物车商品
}

// RemoveCartPieceService 删除购物车中的商品
func RemoveCartPieceService(artID uint) error {
	return dao.RemoveCartPiece(artID) // 调用 DAO 方法删除艺术品
}

// AddTransaction 添加交易记录
func AddTransaction(transaction model.Transaction) (model.Transaction, error) {
	return dao.AddTransaction(transaction) // 调用 DAO 方法添加交易记录
}

// AddOrder 添加订单记录
func AddOrder(order model.Order) error {
	return dao.AddOrder(order) // 调用 DAO 方法添加订单记录
}

// UpdateArtPiece 更新艺术品
func UpdateArtPiece(artPiece model.ArtPiece) error {
	return dao.UpdateArtPiece(artPiece) // 调用 DAO 方法更新艺术品
}

// GetPublishedArtCount 获取已发布艺术品数量
func GetPublishedArtCount(userID uint) (int64, error) {
	return dao.GetPublishedArtCount(userID)
}

// GetPurchasedArtCount 获取已购入艺术品数量
func GetPurchasedArtCount(userID uint) (int64, error) {
	return dao.GetPurchasedArtCount(userID)
}

// GetSoldArtCount 获取已售出艺术品数量
func GetSoldArtCount(userID uint) (int64, error) {
	return dao.GetSoldArtCount(userID)
}

// GetTransactionsByBuyerID 根据 buyer_id 获取交易记录
func GetTransactionsByBuyerID(buyerID uint) ([]model.Transaction, error) {
	return dao.GetTransactionsByBuyerID(buyerID)
}

// GetTransactionsBySellerID 根据 seller_id 获取交易记录
func GetTransactionsBySellerID(sellerID uint) ([]model.Transaction, error) {
	return dao.GetTransactionsBySellerID(sellerID)
}

// GetUsernameByUserID 根据 user_id 获取用户名
func GetUsernameByUserID(userID uint) (string, error) {
	return dao.GetUsernameByUserID(userID)
}

// GetArtPieceByID 根据 art_id 获取艺术品
func GetArtPieceByID(artID uint) (model.ArtPiece, error) {
	return dao.GetArtPieceByID(artID)
}

// GetArtPiecesByCreatorID 根据 creator_id 获取艺术品
func GetArtPiecesByCreatorID(creatorID uint) ([]model.ArtPiece, error) {
	return dao.GetArtPiecesByCreatorID(creatorID)
}
