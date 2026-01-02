package repository

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/common/utils"
	"group-buy-market-go/internal/domain/tag/model"
	"group-buy-market-go/internal/infrastructure"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/po"
)

type TagRepository struct {
	log                *log.Helper
	crowdTagsDAO       dao.CrowdTagsDAO
	crowdTagsDetailDAO dao.CrowdTagsDetailDAO
	crowdTagsJobDAO    dao.CrowdTagsJobDAO
	data               infrastructure.Data
}

// NewTagRepository creates a new tag repository
func NewTagRepository(
	logger log.Logger,
	crowdTagsDAO dao.CrowdTagsDAO,
	crowdTagsDetailDAO dao.CrowdTagsDetailDAO,
	crowdTagsJobDAO dao.CrowdTagsJobDAO,
	data infrastructure.Data,
) *TagRepository {
	return &TagRepository{
		log:                log.NewHelper(logger),
		crowdTagsDAO:       crowdTagsDAO,
		crowdTagsDetailDAO: crowdTagsDetailDAO,
		crowdTagsJobDAO:    crowdTagsJobDAO,
		data:               data,
	}
}

// QueryCrowdTagsJobEntity queries crowd tags job entity by tag id and batch id
func (r *TagRepository) QueryCrowdTagsJobEntity(ctx context.Context, tagId string, batchId string) (*model.CrowdTagsJobEntity, error) {
	crowdTagsJobReq := &po.CrowdTagsJob{
		TagId:   tagId,
		BatchId: batchId,
	}

	crowdTagsJobRes, err := r.crowdTagsJobDAO.QueryCrowdTagsJob(ctx, crowdTagsJobReq)
	if err != nil {
		return nil, err
	}

	if crowdTagsJobRes == nil {
		return nil, nil
	}

	entity := &model.CrowdTagsJobEntity{
		TagType:       crowdTagsJobRes.TagType,
		TagRule:       crowdTagsJobRes.TagRule,
		StatStartTime: crowdTagsJobRes.StatStartTime,
		StatEndTime:   crowdTagsJobRes.StatEndTime,
	}

	return entity, nil
}

// AddCrowdTagsUserId adds a user ID to crowd tags
func (r *TagRepository) AddCrowdTagsUserId(ctx context.Context, tagId string, userId string) error {
	crowdTagsDetail := &po.CrowdTagsDetail{
		TagId:  tagId,
		UserId: userId,
	}

	err := r.crowdTagsDetailDAO.AddCrowdTagsUserId(ctx, crowdTagsDetail)
	if err != nil {
		r.log.Errorf("failed to add crowd tags user id: %v", err)
		return err
	}

	// Implement Redis operations
	// Get index from user ID (simple hash in this case)
	index := utils.GetIndexFromUserId(userId)

	// Set bit in Redis bitmap
	bitKey := "tag:" + tagId
	err = r.data.Rdb.SetBit(ctx, bitKey, index, 1).Err()
	if err != nil {
		// Log error but don't fail the operation
		// In a production environment, you might want to use a proper logger
		r.log.Errorf("failed to set bit in Redis: %v", err)
	}

	return nil
}

// UpdateCrowdTagsStatistics updates crowd tags statistics
func (r *TagRepository) UpdateCrowdTagsStatistics(ctx context.Context, tagId string, count int) error {
	crowdTagsReq := &po.CrowdTags{
		TagId:      tagId,
		Statistics: count,
	}

	return r.crowdTagsDAO.UpdateCrowdTagsStatistics(ctx, crowdTagsReq)
}
