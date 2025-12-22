package model

// GroupBuyOrderAggregate 拼团订单聚合根
type GroupBuyOrderAggregate struct {
	UserEntity        *UserEntity        `json:"userEntity"`
	PayActivityEntity *PayActivityEntity `json:"payActivityEntity"`
	PayDiscountEntity *PayDiscountEntity `json:"payDiscountEntity"`
}
