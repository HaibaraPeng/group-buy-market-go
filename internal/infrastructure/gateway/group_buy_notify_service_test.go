package gateway

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBuyNotifyService_GroupBuyNotify(t *testing.T) {
	// 创建一个模拟的HTTP服务器用于测试
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法
		assert.Equal(t, "POST", r.Method)
		// 验证内容类型
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// 读取请求体
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		expectedBody := "{\"teamId\":\"57199993\",\"outTradeNoList\":[\"038426231487\",\"652896391719\",\"619401409195\"]}"
		assert.JSONEq(t, expectedBody, string(body))

		// 返回成功响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))
	defer server.Close()

	// 创建服务实例
	service := NewGroupBuyNotifyService()

	// 准备测试数据
	notifyRequestDTOJSON := "{\"teamId\":\"57199993\",\"outTradeNoList\":[\"038426231487\",\"652896391719\",\"619401409195\"]}"

	// 执行测试
	response, err := service.GroupBuyNotify(context.Background(), server.URL, notifyRequestDTOJSON)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, "success", response)
}

func TestGroupBuyNotifyService_GroupBuyNotify_Error(t *testing.T) {
	// 创建一个返回错误状态码的模拟服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}))
	defer server.Close()

	// 创建服务实例
	service := NewGroupBuyNotifyService()

	// 准备测试数据
	notifyRequestDTOJSON := "{\"teamId\":\"57199993\",\"outTradeNoList\":[\"038426231487\",\"652896391719\",\"619401409195\"]}"

	// 执行测试
	response, err := service.GroupBuyNotify(context.Background(), server.URL, notifyRequestDTOJSON)

	// 验证结果 - 应该返回错误
	assert.Error(t, err)
	assert.Equal(t, "", response)
	assert.Contains(t, err.Error(), "HTTP request failed with status code: 500")
}

func TestGroupBuyNotifyService_GroupBuyNotify_NetworkError(t *testing.T) {
	// 创建服务实例
	service := NewGroupBuyNotifyService()

	// 使用无效的URL测试网络错误
	notifyRequestDTOJSON := "{\"teamId\":\"57199993\",\"outTradeNoList\":[\"038426231487\",\"652896391719\",\"619401409195\"]}"

	// 执行测试
	response, err := service.GroupBuyNotify(context.Background(), "http://invalid-url-that-does-not-exist-and-will-timeout", notifyRequestDTOJSON)

	// 验证结果 - 应该返回网络错误
	assert.Error(t, err)
	assert.Equal(t, "", response)
	assert.Contains(t, err.Error(), "拼团回调 HTTP 接口服务异常")
}

func TestGroupBuyNotifyService_GroupBuyNotify_InvalidJSON(t *testing.T) {
	// 创建一个模拟的HTTP服务器用于测试
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 返回成功响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))
	defer server.Close()

	// 创建服务实例
	service := NewGroupBuyNotifyService()

	// 准备无效的JSON测试数据
	notifyRequestDTOJSON := "{\"invalid\": json}"

	// 执行测试
	response, err := service.GroupBuyNotify(context.Background(), server.URL, notifyRequestDTOJSON)

	// 验证结果 - 请求应该成功发送，即使JSON无效（因为只是作为字符串传递）
	assert.NoError(t, err)
	assert.Equal(t, "success", response)
}

func TestGroupBuyNotifyService_GroupBuyNotify_WithDifferentJSONFormats(t *testing.T) {
	// 创建一个模拟的HTTP服务器用于测试
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法
		assert.Equal(t, "POST", r.Method)
		// 验证内容类型
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// 读取请求体
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		expectedBody := "{\"teamId\":\"57199993\",\"outTradeNoList\":\"038426231487,652896391719,619401409195\"}"
		assert.JSONEq(t, expectedBody, string(body))

		// 返回成功响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("different_json_format_success"))
	}))
	defer server.Close()

	// 创建服务实例
	service := NewGroupBuyNotifyService()

	// 准备测试数据 - 使用Java测试中的另一种JSON格式
	notifyRequestDTOJSON := "{\"teamId\":\"57199993\",\"outTradeNoList\":\"038426231487,652896391719,619401409195\"}"

	// 执行测试
	response, err := service.GroupBuyNotify(context.Background(), server.URL, notifyRequestDTOJSON)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, "different_json_format_success", response)
}
