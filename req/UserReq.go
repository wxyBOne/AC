package req

// 注册请求对象
type RegisterReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 登录请求对象
type LoginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 重置密码请求对象(登录注册页)
type ResetPasswordReq struct {
	PrivateKey  string `json:"privateKey"`  // 私钥
	NewPassword string `json:"newPassword"` // 新密码
}

// 修改密码请求对象(用户中心)
type UpdateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
