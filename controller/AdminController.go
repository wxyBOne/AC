package controller

import (
	"ac/req"
	"ac/service"
	"ac/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminController struct{}

// AdminRouterController 设置管理员相关的路由
func (ac AdminController) AdminRouterController(engine *gin.Engine) {
	engine.POST("/admin/login", AdminLoginHandle) // 管理员登录
	//用户信息管理模块
	engine.GET("/admin/users", GetAllUsersHandle)          // 获取所有用户
	engine.PUT("/admin/users/:uid", UpdateUserAdminHandle) // 更新用户信息
	engine.DELETE("/admin/users/:uid", DeleteUserHandle)   // 删除用户
	engine.GET("/admin/users/search", SearchUsersHandle)   // 搜索用户
	//艺术品认证确权模块
	engine.GET("/admin/pending-artworks", GetPendingArtworks) // 获取审核艺术品
	engine.POST("/admin/review-artwork", ReviewArtwork)       // 审核艺术品
	//数据监测模块
	engine.GET("/admin/transaction/count", GetTransactionCountHandle)     // 获取交易总量
	engine.GET("/admin/transaction/summary", GetTransactionSummaryHandle) // 获取交易总额和近期交易
}

// AdminLoginHandle 处理管理员登录请求
func AdminLoginHandle(c *gin.Context) {
	var loginReq req.AdminLoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据无效"})
		return
	}

	// 检查输入是否为空
	if loginReq.Adminname == "" || loginReq.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名和密码不能为空"})
		return
	}

	adminID, err := service.LoginAdminService(loginReq.Adminname, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 生成管理员token
	token, err := utils.GenerateAdminToken(adminID, loginReq.Adminname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "登录成功",
		"adminName": loginReq.Adminname,
		"adminID":   adminID,
		"token":     token,
	})
}

// GetAllUsersHandle 获取所有用户
func GetAllUsersHandle(c *gin.Context) {
	// 验证管理员token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	users, err := service.GetAllUsersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUserHandle 更新用户信息
func UpdateUserAdminHandle(c *gin.Context) {
	// 验证管理员token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	uidStr := c.Param("uid")
	uid, err := strconv.ParseUint(uidStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	var updateReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据无效"})
		return
	}

	err = service.UpdateUserByAdminService(uint(uid), updateReq.Username, updateReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}

// DeleteUserHandle 删除用户
func DeleteUserHandle(c *gin.Context) {
	// 验证管理员token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	uidStr := c.Param("uid")
	uid, err := strconv.ParseUint(uidStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	err = service.DeleteUserService(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

// SearchUsersHandle 搜索用户
func SearchUsersHandle(c *gin.Context) {
	// 验证管理员token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	query := c.Query("q")
	users, err := service.SearchUsersService(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "搜索用户失败"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetPendingArtworks 获取待审核艺术品列表
func GetPendingArtworks(c *gin.Context) {
	artworks, err := service.GetPendingArtworksService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取待审核艺术品失败"})
		return
	}
	c.JSON(http.StatusOK, artworks)
}

// ReviewArtwork 处理艺术品审核
func ReviewArtwork(c *gin.Context) {
	artIDStr := c.PostForm("art_id")
	statusStr := c.PostForm("status")
	auditorIDStr := c.PostForm("auditor_id")
	contractHash := c.PostForm("contract_hash")

	artID, _ := strconv.ParseUint(artIDStr, 10, 32)
	status, _ := strconv.Atoi(statusStr)
	auditorID, _ := strconv.ParseUint(auditorIDStr, 10, 32)

	err := service.ReviewArtworkService(uint(artID), uint(auditorID), status, contractHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "审核操作失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "审核完成"})
}

// GetTransactionCountHandle 获取交易总量
func GetTransactionCountHandle(c *gin.Context) {
	// 验证管理员token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	count, err := service.GetTransactionCountService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取交易总量失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

// GetTransactionSummaryHandle 获取交易总额和近期交易
func GetTransactionSummaryHandle(c *gin.Context) {
	// 验证管理员token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	totalAmount, err := service.GetTransactionTotalAmountService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取交易总额失败"})
		return
	}

	recentTransactions, err := service.GetRecentTransactionsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取近期交易失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalAmount":        totalAmount,
		"recentTransactions": recentTransactions,
	})
}
