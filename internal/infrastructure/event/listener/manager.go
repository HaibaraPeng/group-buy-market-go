package listener

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/infrastructure/data"
)

// EventListenerManager 负责管理所有事件监听器
type EventListenerManager struct {
	teamSuccessListener *TeamSuccessEventListener
	logger              *log.Helper
	data                *data.Data
}

// NewEventListenerManager 创建事件监听器管理器
func NewEventListenerManager(
	teamSuccessListener *TeamSuccessEventListener,
	logger log.Logger,
	data *data.Data,
) *EventListenerManager {
	return &EventListenerManager{
		teamSuccessListener: teamSuccessListener,
		logger:              log.NewHelper(logger),
		data:                data,
	}
}

// Start 启动所有事件监听器
func (m *EventListenerManager) Start(ctx context.Context) error {
	m.logger.Info("正在启动事件监听器...")

	// 启动团队成功事件监听器
	err := m.teamSuccessListener.ListenTeamSuccessEvent(ctx)
	if err != nil {
		return fmt.Errorf("failed to start team success event listener: %w", err)
	}

	m.logger.Info("事件监听器启动成功")
	return nil
}

// Stop 停止所有事件监听器
func (m *EventListenerManager) Stop(ctx context.Context) error {
	m.logger.Info("正在停止事件监听器...")
	return nil
}

// Implement transport.Service interface
func (m *EventListenerManager) Run() error {
	// Since the event listeners are initialized in Start(), we just keep this method for interface compliance
	// Actual event processing happens in background goroutines started by Listen methods
	<-make(chan struct{}) // Block forever, actual work happens in background
	return nil
}
