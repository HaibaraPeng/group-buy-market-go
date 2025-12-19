package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// TagService 标签服务
// 提供对外的标签服务接口
type TagService struct {
	v1.UnimplementedTagHTTPServer
	log           *log.Helper
	tagRepository *repository.TagRepository
}

// NewTagService 创建标签服务实例
func NewTagService(logger log.Logger, tagRepository *repository.TagRepository) *TagService {
	return &TagService{
		log:           log.NewHelper(logger),
		tagRepository: tagRepository,
	}
}

// ExecTagBatchJob 执行标签批处理任务
func (s *TagService) ExecTagBatchJob(ctx context.Context, tagId string, batchId string) error {
	s.log.Infof("人群标签批次任务 tagId:%s batchId:%s", tagId, batchId)

	// 1. 查询批次任务
	crowdTagsJobEntity, err := s.tagRepository.QueryCrowdTagsJobEntity(ctx, tagId, batchId)
	if err != nil {
		s.log.Errorf("查询批次任务失败: %v", err)
		return err
	}

	// 检查任务是否存在
	if crowdTagsJobEntity == nil {
		s.log.Warnf("未找到标签批次任务 tagId:%s batchId:%s", tagId, batchId)
		return nil
	}

	// 2. 采集用户数据 - 这部分需要采集用户的消费类数据，后续有用户发起拼单后再处理。

	// 3. 数据写入记录
	// 模拟用户列表
	userIdList := []string{"xiaofuge", "liergou"}

	// 4. 一般人群标签的处理在公司中，会有专门的数据数仓团队通过脚本方式写入到数据库，就不用这样一个个或者批次来写。
	for _, userId := range userIdList {
		err := s.tagRepository.AddCrowdTagsUserId(ctx, tagId, userId)
		if err != nil {
			s.log.Errorf("添加用户标签失败 tagId:%s userId:%s error:%v", tagId, userId, err)
			// 根据业务需求决定是否继续处理其他用户或直接返回错误
			// 这里我们记录错误但继续处理其他用户
		}
	}

	// 5. 更新人群标签统计量
	err = s.tagRepository.UpdateCrowdTagsStatistics(ctx, tagId, len(userIdList))
	if err != nil {
		s.log.Errorf("更新人群标签统计量失败: %v", err)
		return err
	}

	return nil
}
