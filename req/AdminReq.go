package req

// AdminLoginReq 管理员登录请求对象
type AdminLoginReq struct {
	Adminname string `json:"adminname" binding:"required"` // 管理员用户名
	Password  string `json:"password" binding:"required"`  // 密码
}
