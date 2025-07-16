package controller

import (
	"ac/dao"
	"ac/req"
	"ac/service"
	"ac/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// UserRouterController 设置用户相关的路由
func (uc UserController) UserRouterController(engine *gin.Engine) {
	//用户注册模块
	engine.POST("/register", RegisterUserHandle) // 注册用户
	//用户登录模块
	engine.POST("/login", LoginUserHandle)              // 登录用户
	engine.POST("/reset-password", ResetPasswordHandle) // 重置密码
	//用户信息管理模块
	engine.POST("/update-user", UpdateUserHandle) // 修改个人信息
}

// RegisterUserHandle 处理用户注册请求
func RegisterUserHandle(c *gin.Context) {
	var registerReq req.RegisterReq
	err := c.ShouldBindJSON(&registerReq) // 绑定请求数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // 输入无效
		return
	}

	// 检查密码长度
	if len(registerReq.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"}) // 密码长度检查
		return
	}

	privateKey := utils.CustomGenerateKey()                                                                                       // 生成私钥
	_, err = service.RegisterUserService(registerReq.Username, registerReq.Password, privateKey.Privatekey, privateKey.Publickey) // 注册用户
	if err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"}) // 用户名已存在
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"}) // 注册失败
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful", "privateKey": privateKey.Privatekey}) // 返回私钥
}

// LoginUserHandle 处理用户登录请求
func LoginUserHandle(c *gin.Context) {
	var loginReq req.LoginReq
	err := c.ShouldBindJSON(&loginReq) // 绑定请求数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // 输入无效
		return
	}

	// 检查输入是否为空
	if loginReq.Username == "" || loginReq.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password cannot be empty"}) // 输入框不可为空
		return
	}

	_, privateKey, err := service.LoginUserService(loginReq.Username, loginReq.Password) // 登录用户
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // 用户未注册或密码错误
		return
	}

	// 生成 JWT
	userID, err := dao.GetUserIDByUsername(loginReq.Username)
	token, err := utils.GenerateToken(userID, loginReq.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"}) // 生成 token 失败
		return
	}

	// 返回用户信息，包括用户名、私钥和 token
	c.JSON(http.StatusOK, gin.H{
		"username":   loginReq.Username, // 确保返回用户名
		"privateKey": privateKey,        // 返回私钥
		"token":      token,             // 返回 JWT
	})
}

// ResetPasswordHandle 处理重置密码请求
func ResetPasswordHandle(c *gin.Context) {
	var resetReq req.ResetPasswordReq
	err := c.ShouldBindJSON(&resetReq) // 绑定请求数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // 输入无效
		return
	}

	// 查密码长度
	if len(resetReq.NewPassword) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"}) // 密码长度检查
		return
	}

	// 检查输入是否为空
	if resetReq.PrivateKey == "" || resetReq.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Private key and new password cannot be empty"}) // 输入框不可为空
		return
	}

	err = service.ResetPasswordService(resetReq.PrivateKey, resetReq.NewPassword) // 重置密码
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password reset failed"}) // 重置失败
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"}) // 返回成功响应
}

// UpdateUserHandle 处理用户信息更新请求
func UpdateUserHandle(c *gin.Context) {
	var updateReq req.UpdateUserReq
	err := c.ShouldBindJSON(&updateReq) // 绑定请求数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"}) // 输入无效
		return
	}

	// 输入验证
	if updateReq.Username == "" || updateReq.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password cannot be empty"}) // 输入框不可为空
		return
	}

	if len(updateReq.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"}) // 密码长度检查
		return
	}

	// 从请求头获取 token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	// 验证 token 并获取当前用户名
	_, currentUsername, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token无效"})
		return
	}

	// 更新用户信息逻辑
	status, err := service.UpdateUserService(currentUsername, updateReq.Username, updateReq.Password)
	if err != nil {
		if status == 1 {
			c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"}) // 用户名已存在
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"}) // 更新失败
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update successful"}) // 返回成功消息
}
