package model2

// XxxResponse 是对应Java版本XxxResponse类的Go实现
type XxxResponse struct {
	Age string
}

// NewXxxResponse 创建一个新的XxxResponse实例
func NewXxxResponse(age string) *XxxResponse {
	return &XxxResponse{
		Age: age,
	}
}

// GetAge 返回年龄
func (xr *XxxResponse) GetAge() string {
	return xr.Age
}
