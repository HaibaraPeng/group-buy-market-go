package types

// MarketProductEntity 营销商品实体
// 表示参与营销活动的商品信息
type MarketProductEntity struct {
	ID          int64   `json:"id"`          // 商品ID
	Name        string  `json:"name"`        // 商品名称
	Description string  `json:"description"` // 商品描述
	Price       float64 `json:"price"`       // 商品价格
	SkuID       int64   `json:"skuId"`       // SKU ID
	Stock       int32   `json:"stock"`       // 库存数量
}

// TrialBalanceEntity 试算平衡实体
// 表示经过各节点处理后的试算结果
type TrialBalanceEntity struct {
	TotalAmount    float64 `json:"totalAmount"`    // 总金额
	DiscountAmount float64 `json:"discountAmount"` // 折扣金额
	FinalAmount    float64 `json:"finalAmount"`    // 最终金额
	Success        bool    `json:"success"`        // 是否成功
	Message        string  `json:"message"`        // 处理消息
	Code           string  `json:"code"`           // 结果码
}
