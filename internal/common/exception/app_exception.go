package exception

import (
	"fmt"
	"group-buy-market-go/internal/common/consts"
)

// AppException 应用异常结构体
type AppException struct {
	Code    consts.ResponseCode
	Message string
}

// Error 实现error接口
func (e *AppException) Error() string {
	return fmt.Sprintf("错误代码: %s, 错误信息: %s", e.Code, e.Message)
}

// NewAppException 创建新的应用异常
func NewAppException(code consts.ResponseCode) *AppException {
	return &AppException{
		Code:    code,
		Message: code.GetErrorMessage(),
	}
}

// NewAppExceptionWithMessage 创建带自定义消息的应用异常
func NewAppExceptionWithMessage(code consts.ResponseCode, message string) *AppException {
	return &AppException{
		Code:    code,
		Message: message,
	}
}
