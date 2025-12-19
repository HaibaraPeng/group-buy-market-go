package repository

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"group-buy-market-go/internal/domain/tag/model"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/po"
)

type TagRepository struct {
	log                *log.Helper
	crowdTagsDAO       dao.CrowdTagsDAO
	crowdTagsDetailDAO dao.CrowdTagsDetailDAO
	crowdTagsJobDAO    dao.CrowdTagsJobDAO
	redisClient        *redis.Client
}

// NewTagRepository creates a new tag repository
func NewTagRepository(
	logger log.Logger,
	crowdTagsDAO dao.CrowdTagsDAO,
	crowdTagsDetailDAO dao.CrowdTagsDetailDAO,
	crowdTagsJobDAO dao.CrowdTagsJobDAO,
	redisClient *redis.Client,
) *TagRepository {
	return &TagRepository{
		log:                log.NewHelper(logger),
		crowdTagsDAO:       crowdTagsDAO,
		crowdTagsDetailDAO: crowdTagsDetailDAO,
		crowdTagsJobDAO:    crowdTagsJobDAO,
		redisClient:        redisClient,
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
	index := r.getIndexFromUserId(userId)

	// Set bit in Redis bitmap
	bitKey := "tag:" + tagId
	err = r.redisClient.SetBit(ctx, bitKey, index, 1).Err()
	if err != nil {
		// Log error but don't fail the operation
		// In a production environment, you might want to use a proper logger
		r.log.Errorf("failed to set bit in Redis: %v", err)
	}

	return nil
}

// getIndexFromUserId generates an index from user ID
// This is a simple implementation - in production you might want to use a more sophisticated hashing
func (r *TagRepository) getIndexFromUserId(userId string) int64 {
	var hash int64
	for _, char := range userId {
		hash = (hash*31 + int64(char)) & 0x7FFFFFFF // Keep it positive
	}
	return hash
}

// UpdateCrowdTagsStatistics updates crowd tags statistics
func (r *TagRepository) UpdateCrowdTagsStatistics(ctx context.Context, tagId string, count int) error {
	crowdTagsReq := &po.CrowdTags{
		TagId:      tagId,
		Statistics: count,
	}

	return r.crowdTagsDAO.UpdateCrowdTagsStatistics(ctx, crowdTagsReq)
}
