package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/infrastructure/dcc"
)

// DccService DCC服务
// 提供对外的DCC服务接口
type DccService struct {
	v1.UnimplementedDccHTTPServer
	log *log.Helper
	dcc *dcc.DCC
}

// NewDccService 创建标签服务实例
func NewDccService(logger log.Logger, dcc *dcc.DCC) *DccService {
	return &DccService{
		log: log.NewHelper(logger),
		dcc: dcc,
	}
}

// UpdateConfig 更新配置
func (s *DccService) UpdateConfig(ctx context.Context, req *v1.UpdateConfigRequest) (*v1.UpdateConfigReply, error) {
	s.log.Infof("DCC 动态配置值变更 key:%s value:%s", req.Key, req.Value)

	// 发布配置变更消息到Redis
	err := s.dcc.PublishConfigChange(ctx, req.Key, req.Value)
	if err != nil {
		s.log.Errorf("DCC 动态配置值变更失败 key:%s value:%s err:%v", req.Key, req.Value, err)
		return &v1.UpdateConfigReply{Success: false}, nil
	}

	return &v1.UpdateConfigReply{Success: true}, nil
}
