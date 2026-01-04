package gateway

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"group-buy-market-go/internal/common/consts"
)

// GroupBuyNotifyService 拼团回调服务
type GroupBuyNotifyService struct {
	httpClient *http.Client
}

// NewGroupBuyNotifyService 创建拼团回调服务实例
func NewGroupBuyNotifyService() *GroupBuyNotifyService {
	return &GroupBuyNotifyService{
		httpClient: &http.Client{},
	}
}

// GroupBuyNotify 拼团回调方法
func (g *GroupBuyNotifyService) GroupBuyNotify(ctx context.Context, apiUrl string, notifyRequestDTOJSON string) (string, error) {
	// 1. 构建参数
	body := bytes.NewBufferString(notifyRequestDTOJSON)
	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, body)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// 2. 调用接口
	resp, err := g.httpClient.Do(req)
	if err != nil {
		return "", exception.NewAppExceptionWithMessage(
			consts.HTTP_EXCEPTION,
			fmt.Sprintf("拼团回调 HTTP 接口服务异常 %s: %v", apiUrl, err),
		)
	}
	defer resp.Body.Close()

	// 3. 返回结果
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", exception.NewAppExceptionWithMessage(
			consts.HTTP_EXCEPTION,
			fmt.Sprintf("HTTP request failed with status code: %d", resp.StatusCode),
		)
	}

	return string(result), nil
}
