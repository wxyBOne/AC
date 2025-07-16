package model

// User 代表用户模型
type User struct {
	Uid        uint   `gorm:"primaryKey"` // 用户ID
	Username   string `gorm:"unique"`     // 用户名
	Userpass   string // 用户密码
	PrivateKey string // 私钥
	PublicKey  string // 公钥
}
