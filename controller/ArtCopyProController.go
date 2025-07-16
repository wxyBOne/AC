package controller

import (
	"ac/model"
	"ac/req"
	"ac/service"
	"ac/utils"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ArtCopyProController struct{}

func (ac *ArtCopyProController) ArtRouterController(engine *gin.Engine) {
	//艺术品交易模块
	engine.POST("/add-art-piece", ac.AddArtPiece)        // 添加艺术品
	engine.GET("/art-pieces", ac.GetArtItem)             // 获取艺术品列表
	engine.GET("/art-pieces/:id", ac.GetArtProByID)      // 根据 ID 获取艺术品
	engine.POST("/add-to-cart", ac.AddToCart)            // 添加商品到购物车
	engine.GET("/cart-items", ac.GetCartItems)           // 获取购物车商品
	engine.DELETE("/cart-items/:aid", ac.RemoveFromCart) // 删除购物车商品
	engine.POST("/checkout", ac.Checkout)                // 结算
	//交易记录模块
	engine.GET("/user-art-stats", ac.GetUserArtStats)      // 获取获取用户交易记录
	engine.GET("/purchased-orders", ac.GetPurchasedOrders) // 获取购买订单
	engine.GET("/sold-orders", ac.GetSoldOrders)           // 获取售出订单
	engine.GET("/published-art", ac.GetPublishedArt)       // 获取已发布艺术品
}

// GetArtPieces 处理获取艺术品列表请求
func (ac *ArtCopyProController) GetArtItem(c *gin.Context) {
	artPieces, err := service.GetArtItemService() // 调用服务层获取艺术品
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 获取失败
		return
	}
	c.JSON(http.StatusOK, artPieces) // 返回艺术品列表
}

// AddArtPiece 处理添加艺术品请求
func (ac *ArtCopyProController) AddArtPiece(c *gin.Context) {
	// 获取文本字段值
	title := c.PostForm("title")
	description := c.PostForm("description")
	tag := c.PostForm("tag")
	priceStr := c.PostForm("price")

	// 验证必填项
	if title == "" || description == "" {
		fmt.Println("输入无效: 标题或描述不能为空")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // 输入无效
		return
	}

	// 从请求头获取 token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		fmt.Println("未登录")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	// 验证 token 并获取当前用户 ID
	currentUserID, _, err := utils.ValidateToken(tokenString) // 获取用户 ID
	if err != nil {
		fmt.Println("token无效:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	// 处理文件上传
	file, err := c.FormFile("artImage")
	if err != nil {
		fmt.Println("文件上传失败:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"}) // 文件上传失败
		return
	}

	// 获取当前时间戳并格式化为字符串，用于文件名
	timestamp := time.Now().Format("20060102150405")
	newFileName := fmt.Sprintf("%s_%s", timestamp, file.Filename) // 使用时间戳和原文件名

	// 保存文件到指定目录
	filePath := filepath.Join("D:\\Merkle Tree\\Go\\GOitem\\src\\ac\\web\\project\\src\\upload-img", newFileName) // 假设 uploads 是存储文件的目录
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"}) // 保存文件失败
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Println("价格格式无效，无法转换为数字:", err)
		return
	}

	// 构建 ArtPiece 结构体
	artPiece := model.ArtPiece{
		Title:        title,
		Description:  description,
		Tag:          tag,
		ArtImage:     newFileName, // 保存图片名称
		ArtPrice:     price,
		CreatorID:    currentUserID, // 设置创作者 ID
		ContractHash: "",            // 合约标识先留空
		UserID:       currentUserID, // 设置权利持有者 ID
		IsSold:       0,
	}

	aid, err := service.AddArtPieceService(artPiece) // 调用服务层添加艺术品
	if err != nil {
		fmt.Println("添加失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 添加失败
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Art piece added successfully", "aid": aid}) // 返回成功消息
}

// GetArtProByID 处理根据 ID 获取艺术品请求
func (ac *ArtCopyProController) GetArtProByID(c *gin.Context) {
	id := c.Param("id") // 获取 URL 中的 ID 参数
	intID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("string转int失败:", err)
		return
	}
	uintID := uint(intID)
	artPro, err := service.GetArtProByIDService(uintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 获取失败
		return
	}

	c.JSON(http.StatusOK, artPro) // 返回艺术品数据
}

// AddToCart 处理添加商品到购物车请求
func (ac *ArtCopyProController) AddToCart(c *gin.Context) {
	// 从请求头获取 token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		fmt.Println("未登录")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	// 验证 token 并获取当前用户 ID
	currentUserID, _, err := utils.ValidateToken(tokenString) // 获取用户 ID
	if err != nil {
		fmt.Println("token无效:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}
	artIDStr := c.PostForm("art_id") // 从请求中获取艺术品 ID

	// 验证艺术品 ID
	if artIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	artID, err := strconv.ParseUint(artIDStr, 10, 32) // 转换为 uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid art ID"})
		return
	}

	// 获取艺术品信息以验证 CreatorID
	artPiece, err := service.GetArtPieceByIDService(uint(artID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取艺术品信息"})
		return
	}

	// 检查当前用户是否是艺术品的创作者
	if artPiece.CreatorID == currentUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "您不能购买自己发布的艺术商品"})
		return
	}

	// 调用服务层添加到购物车
	err = service.AddToCartService(currentUserID, uint(artID))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "该商品已在购物车中，请勿重复添加"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "已加入购物车"})
}

// GetCartItems 处理获取购物车商品请求
func (ac *ArtCopyProController) GetCartItems(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	userID, _, err := utils.ValidateToken(tokenString) // 获取用户 ID
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	artPieces, err := service.GetCartItemsService(userID) // 调用服务层获取艺术品
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 获取失败
		return
	}

	c.JSON(http.StatusOK, artPieces) // 返回艺术品数据
}

// RemoveFromCart 处理删除购物车商品请求
func (ac *ArtCopyProController) RemoveFromCart(c *gin.Context) {
	// 从请求头获取 token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		fmt.Println("未登录")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}
	// 验证 token 并获取当前用户 ID
	userID, _, err := utils.ValidateToken(tokenString) // 获取用户 ID
	if err != nil {
		fmt.Println("token无效:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	artIDStr := c.Param("aid")                        // 获取 URL 中的艺术品 ID
	artID, err := strconv.ParseUint(artIDStr, 10, 32) // 转换为 uint
	if err != nil {
		fmt.Println("Invalid art ID:", artIDStr, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid art ID"})
		return
	}

	// 调用服务层删除购物车商品
	err = service.RemoveFromCartService(uint(artID), userID)
	if err != nil {
		fmt.Println("删除失败：", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // 删除失败
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "商品已从购物车中删除"})
}

// Checkout 处理购物车结算请求
func (ac *ArtCopyProController) Checkout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	// 验证 token 并获取当前用户 ID
	userID, _, err := utils.ValidateToken(tokenString)
	if err != nil {
		fmt.Println("token无效", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	// 解析请求体
	var checkoutReq req.CheckoutRequest
	if err := c.ShouldBindJSON(&checkoutReq); err != nil {
		fmt.Println("解析请求数据失败", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 获取艺术品信息
	artPiece, err := service.GetArtPieceByIDService(checkoutReq.ArtworkId)
	if err != nil {
		fmt.Println("无法获取艺术品信息", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取艺术品信息"})
		return
	}

	// 创建交易记录
	transaction := model.Transaction{
		ArtID:           checkoutReq.ArtworkId,
		BuyerID:         userID,
		SellerID:        artPiece.UserID,
		Price:           artPiece.ArtPrice,
		IsCompleted:     1,
		CreatedAt:       time.Now(),
		TransactionHash: checkoutReq.TransactionHash, // 添加交易哈希
	}

	// 添加交易记录到数据库
	newTransaction, err := service.AddTransaction(transaction)
	if err != nil {
		fmt.Println("添加交易记录失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加交易记录失败"})
		return
	}

	// 创建买家订单记录
	orderBuyer := model.Order{
		Tid:         newTransaction.Tid, // 关联交易 ID
		UserID:      userID,
		ArtID:       checkoutReq.ArtworkId,
		OrderStatus: "bought",
		CreateTime:  time.Now(), // 当前时间
	}

	// 创建卖家订单记录
	orderSeller := model.Order{
		Tid:         newTransaction.Tid, // 关联交易 ID
		UserID:      artPiece.UserID,
		ArtID:       checkoutReq.ArtworkId,
		OrderStatus: "sold",
		CreateTime:  time.Now(), // 当前时间
	}

	// 添加订单记录到数据库
	if err := service.AddOrder(orderBuyer); err != nil {
		fmt.Println("添加买家订单记录失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加买家订单记录失败"})
		return
	}

	// 添加订单记录到数据库
	if err := service.AddOrder(orderSeller); err != nil {
		fmt.Println("添加卖家订单记录失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加卖家订单记录失败"})
		return
	}

	// 更新艺术品信息
	artPiece.UserID = userID // 更新版权持有者为买家
	artPiece.IsSold = 1      // 更新状态为已售出
	if err := service.UpdateArtPiece(artPiece); err != nil {
		fmt.Println("更新艺术品状态失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新艺术品状态失败"})
		return
	}

	// 删除购物车中的商品
	if err := service.RemoveCartPieceService(checkoutReq.ArtworkId); err != nil {
		fmt.Println("删除购物车商品失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除购物车商品失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "购买成功"})
}

// GetUserArtStats 获取用户交易记录
func (ac *ArtCopyProController) GetUserArtStats(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	// 验证 token 并获取当前用户 ID
	userID, _, err := utils.ValidateToken(tokenString) // 获取用户 ID
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	// 查询已发布艺术品数量
	publishedCount, err := service.GetPublishedArtCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取已发布艺术品数量失败"})
		return
	}

	// 查询已购入艺术品数量
	purchasedCount, err := service.GetPurchasedArtCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取已购入艺术品数量失败"})
		return
	}

	// 查询已售出艺术品数量
	soldCount, err := service.GetSoldArtCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取已售出艺术品数量失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"published": publishedCount,
		"purchased": purchasedCount,
		"sold":      soldCount,
	})
}

// GetPurchasedOrders 获取当前用户的购买订单
func (ac *ArtCopyProController) GetPurchasedOrders(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	userID, _, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	transactions, err := service.GetTransactionsByBuyerID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取购买订单失败"})
		return
	}

	var orderDetails []model.OrderDetail
	for _, transaction := range transactions {
		sellerUsername, _ := service.GetUsernameByUserID(transaction.SellerID)
		artPiece, _ := service.GetArtPieceByID(transaction.ArtID)

		orderDetail := model.OrderDetail{
			SellerUsername:  sellerUsername,
			ArtTitle:        artPiece.Title,
			ArtImage:        artPiece.ArtImage,
			Price:           transaction.Price,
			TransactionHash: transaction.TransactionHash,
			CreateTime:      transaction.CreatedAt,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	c.JSON(http.StatusOK, orderDetails)
}

// GetSoldOrders 获取当前用户的售出订单
func (ac *ArtCopyProController) GetSoldOrders(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	userID, _, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	transactions, err := service.GetTransactionsBySellerID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取售出订单失败"})
		return
	}

	var orderDetails []model.OrderDetail
	for _, transaction := range transactions {
		buyerUsername, _ := service.GetUsernameByUserID(transaction.BuyerID)
		artPiece, _ := service.GetArtPieceByID(transaction.ArtID)

		orderDetail := model.OrderDetail{
			BuyerUsername:   buyerUsername,
			ArtTitle:        artPiece.Title,
			ArtImage:        artPiece.ArtImage,
			Price:           transaction.Price,
			TransactionHash: transaction.TransactionHash,
			CreateTime:      transaction.CreatedAt,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	c.JSON(http.StatusOK, orderDetails)
}

// GetPublishedArt 获取当前用户发布的艺术品
func (ac *ArtCopyProController) GetPublishedArt(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	userID, _, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	artPieces, err := service.GetArtPiecesByCreatorID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取已发布艺术品失败"})
		return
	}

	c.JSON(http.StatusOK, artPieces)
}
