package model

import "time"

// ArtPiece 艺术品模型
type ArtPiece struct {
	Aid          uint    `json:"aid" gorm:"primaryKey"` // 艺术品 ID
	Title        string  `json:"title"`                 // 艺术品标题
	Description  string  `json:"description"`           // 艺术品描述
	Tag          string  `json:"tag"`                   // 艺术品标签
	ArtImage     string  `json:"art_image"`             // 图片
	ArtPrice     float64 `json:"art_price"`             // 艺术品价格
	CreatorID    uint    `json:"creator_id"`            // 创作者 ID
	ContractHash string  `json:"contract_hash"`         // 合约标识
	UserID       uint    `json:"user_id"`               // 权利持有者 ID
	IsSold       int     `json:"is_sold"`               // 售出状态，0 为未售出，1 为售出
}

type ArtItem struct {
	Aid          uint    `gorm:"primaryKey" json:"aid"` // 艺术品 ID
	Title        string  `json:"title"`                 // 艺术品标题
	Description  string  `json:"description"`           // 艺术品描述
	Tag          string  `json:"tag"`                   // 艺术品标签
	ArtImage     string  `json:"art_image"`             // 图片
	ArtPrice     float64 `json:"art_price"`             // 艺术品价格
	CreatorName  string  `json:"creator_name"`          // 创作者名字
	ContractHash string  `json:"contract_hash"`         // 合约标识
}

// ShoppingCart 购物车模型
type ShoppingCart struct {
	Scid   uint `gorm:"primaryKey" json:"scid"` // 购物车记录 ID
	UserID uint `json:"user_id"`                // 买家 ID
	ArtID  uint `json:"art_id"`                 // 艺术品 ID
}

// Transaction 交易记录模型
type Transaction struct {
	Tid             uint      `json:"tid" gorm:"primaryKey"` // 交易 ID
	ArtID           uint      `json:"art_id"`                // 艺术品 ID
	BuyerID         uint      `json:"buyer_id"`              // 买家 ID
	SellerID        uint      `json:"seller_id"`             // 卖家 ID
	Price           float64   `json:"price"`                 // 交易价格
	IsCompleted     int       `json:"is_completed"`          // 是否完成，0 为未完成，1 为完成
	TransactionHash string    `json:"transaction_hash"`      // 交易哈希
	CreatedAt       time.Time `json:"created_at"`            // 创建时间
}

// OrderDetail 订单详细信息
type OrderDetail struct {
	SellerUsername  string    `json:"seller_username"`  // 卖家用户名
	BuyerUsername   string    `json:"buyer_username"`   // 买家用户名
	ArtTitle        string    `json:"art_title"`        // 艺术品标题
	ArtImage        string    `json:"art_image"`        // 艺术品图片
	Price           float64   `json:"price"`            // 价格
	TransactionHash string    `json:"transaction_hash"` // 交易哈希
	CreateTime      time.Time `json:"create_time"`      // 创建时间
}

// Order 订单模型
type Order struct {
	Oid         uint      `json:"oid" gorm:"primaryKey"` // 订单 ID
	Tid         uint      `json:"tid"`                   // 交易 ID
	UserID      uint      `json:"user_id"`               // 用户 ID
	ArtID       uint      `json:"art_id"`                // 艺术品 ID
	OrderStatus string    `json:"order_status"`          // 订单状态
	CreateTime  time.Time `json:"create_time"`           // 创建时间
}
