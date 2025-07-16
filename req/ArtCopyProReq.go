package req

// ArtPieceReq 艺术品请求对象
type ArtPieceReq struct {
	Title       string  `json:"title" binding:"required"`       // 艺术品标题
	Description string  `json:"description" binding:"required"` // 艺术品描述
	Tag         string  `json:"tag"`                            // 艺术品标签
	ArtImage    string  `json:"art_image" binding:"required"`   // 图片
	ArtPrice    float64 `json:"art_price" binding:"required"`   // 艺术品价格
}

type CheckoutRequest struct {
	ArtworkId       uint   `json:"artworkId"`
	TransactionHash string `json:"transactionHash"`
}
