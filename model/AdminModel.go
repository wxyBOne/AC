package model

// Admin 代表管理员模型
type Admin struct {
	AdminID   uint   `gorm:"column:admin_id;primaryKey"` // 管理员ID
	Adminname string `gorm:"column:adminname;unique"`    // 管理员用户名
	Adminpass string `gorm:"column:adminpass"`           // 管理员密码
}

// ReviewRecord 审核记录模型
type ReviewRecord struct {
	Rid       uint `gorm:"column:rid;primaryKey"` // 审核记录ID
	ArtID     uint `gorm:"column:art_id"`         // 艺术品ID
	AuditorID uint `gorm:"column:auditor_id"`     // 审核员ID
	Status    int  `gorm:"column:status"`         // 审核状态：0-不通过，1-通过
}

// ArtReviewItem 用于前端展示的艺术品审核信息
type ArtReviewItem struct {
	Aid         uint    `json:"aid"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Tag         string  `json:"tag"`
	ArtImage    string  `json:"art_image"`
	ArtPrice    float64 `json:"art_price"`
	CreatorName string  `json:"creator_name"`
	Status      int     `json:"status"` // 审核状态
}
