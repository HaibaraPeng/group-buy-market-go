package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/common/consts"
	"group-buy-market-go/internal/common/utils"
	"group-buy-market-go/internal/domain/activity/biz/trial/core"
	"group-buy-market-go/internal/domain/activity/biz/trial/factory"
	"group-buy-market-go/internal/domain/activity/biz/trial/node"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"time"
)

// IndexService 营销首页服务
// 提供对外的营销首页服务接口
type IndexService struct {
	v1.UnimplementedIndexHTTPServer
	strategyFactory    *factory.DefaultActivityStrategyFactory
	activityRepository *repository.ActivityRepository
}

// NewIndexService 创建营销首页服务实例
func NewIndexService(rootNode *node.RootNode, activityRepository *repository.ActivityRepository) *IndexService {
	// 构建策略树：根节点 -> 开关节点 -> 营销节点 -> 结束节点
	strategyFactory := factory.NewDefaultActivityStrategyFactory(rootNode)

	return &IndexService{
		strategyFactory:    strategyFactory,
		activityRepository: activityRepository,
	}
}

// QueryGroupBuyMarketConfig 查询拼团营销配置
// 对应Java中的queryGroupBuyMarketConfig方法
func (s *IndexService) QueryGroupBuyMarketConfig(ctx context.Context, req *v1.QueryGroupBuyMarketConfigRequest) (*v1.QueryGroupBuyMarketConfigReply, error) {
	log.Infof("查询拼团营销配置开始 userId:%s goodsId:%s", req.GetUserId(), req.GetGoodsId())

	// 参数校验
	if req.GetUserId() == "" || req.GetSource() == "" || req.GetChannel() == "" || req.GetGoodsId() == "" {
		return nil, fmt.Errorf("%s: %s", string(consts.ILLEGAL_PARAMETER), consts.ILLEGAL_PARAMETER.GetErrorMessage())
	}

	// Create market product entity
	marketProduct := &model.MarketProductEntity{
		UserId:  req.UserId,
		GoodsId: req.GoodsId,
		Source:  req.Source,
		Channel: req.Channel,
	}

	// 1. 营销优惠试算
	trialBalanceEntity, err := s.MarketTrial(ctx, marketProduct)
	if err != nil {
		return nil, err
	}

	groupBuyActivityDiscountVO := trialBalanceEntity.GroupBuyActivityDiscountVO
	activityId := groupBuyActivityDiscountVO.ActivityId

	// 2. 查询拼团组队
	userGroupBuyOrderDetailEntities, err := s.QueryInProgressUserGroupBuyOrderDetailList(ctx, activityId, req.UserId, 1, 2)
	if err != nil {
		return nil, err
	}

	// 3. 统计拼团数据
	teamStatisticVO, err := s.QueryTeamStatisticByActivityId(ctx, activityId)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	goods := &v1.Goods{
		GoodsId:        trialBalanceEntity.GoodsId,
		OriginalPrice:  trialBalanceEntity.OriginalPrice,
		DeductionPrice: trialBalanceEntity.DeductionPrice,
		PayPrice:       trialBalanceEntity.PayPrice,
	}

	teams := make([]*v1.Team, 0)
	if userGroupBuyOrderDetailEntities != nil && len(userGroupBuyOrderDetailEntities) > 0 {
		for _, userGroupBuyOrderDetailEntity := range userGroupBuyOrderDetailEntities {
			team := &v1.Team{
				UserId:             userGroupBuyOrderDetailEntity.UserId,
				TeamId:             userGroupBuyOrderDetailEntity.TeamId,
				ActivityId:         userGroupBuyOrderDetailEntity.ActivityId,
				TargetCount:        int32(userGroupBuyOrderDetailEntity.TargetCount),
				CompleteCount:      int32(userGroupBuyOrderDetailEntity.CompleteCount),
				LockCount:          int32(userGroupBuyOrderDetailEntity.LockCount),
				ValidStartTime:     userGroupBuyOrderDetailEntity.ValidStartTime.Unix(),
				ValidEndTime:       userGroupBuyOrderDetailEntity.ValidEndTime.Unix(),
				ValidTimeCountdown: utils.DifferenceDateTime2Str(time.Now(), userGroupBuyOrderDetailEntity.ValidEndTime),
				OutTradeNo:         userGroupBuyOrderDetailEntity.OutTradeNo,
			}
			teams = append(teams, team)
		}
	}

	teamStatistic := &v1.TeamStatistic{
		AllTeamCount:         int32(teamStatisticVO.AllTeamCount),
		AllTeamCompleteCount: int32(teamStatisticVO.AllTeamCompleteCount),
		AllTeamUserCount:     int32(teamStatisticVO.AllTeamUserCount),
	}

	reply := &v1.QueryGroupBuyMarketConfigReply{
		Goods:         goods,
		TeamList:      teams,
		TeamStatistic: teamStatistic,
	}

	log.Infof("查询拼团营销配置完成 userId:%s goodsId:%s", req.GetUserId(), req.GetGoodsId())

	return reply, nil
}

// IndexMarketTrial 营销首页试算
// 对应Java中的IndexMarketTrial方法
func (s *IndexService) MarketTrial(ctx context.Context, marketProduct *model.MarketProductEntity) (*model.TrialBalanceEntity, error) {
	log.Infof("营销首页试算 marketProduct:%v", marketProduct)

	// 获取策略处理器
	strategyHandler := s.strategyFactory.StrategyHandler()

	// 创建动态上下文
	dynamicContext := &core.DynamicContext{}

	// 应用策略处理器
	trialBalanceEntity, err := strategyHandler.Apply(ctx, marketProduct, dynamicContext)
	if err != nil {
		return nil, err
	}
	return trialBalanceEntity, nil
}

// QueryInProgressUserGroupBuyOrderDetailList 查询进行中的拼团订单详情列表
// 对应Java中的queryInProgressUserGroupBuyOrderDetailList方法
func (s *IndexService) QueryInProgressUserGroupBuyOrderDetailList(ctx context.Context, activityId int64, userId string, ownerCount int, randomCount int) ([]*model.UserGroupBuyOrderDetailEntity, error) {
	unionAllList := make([]*model.UserGroupBuyOrderDetailEntity, 0)

	// 查询个人拼团数据
	if ownerCount != 0 {
		ownerList, err := s.activityRepository.QueryInProgressUserGroupBuyOrderDetailListByOwner(ctx, activityId, userId, ownerCount)
		if err != nil {
			return nil, err
		}
		if ownerList != nil && len(ownerList) > 0 {
			unionAllList = append(unionAllList, ownerList...)
		}
	}

	// 查询其他非个人拼团
	if randomCount != 0 {
		randomList, err := s.activityRepository.QueryInProgressUserGroupBuyOrderDetailListByRandom(ctx, activityId, userId, randomCount)
		if err != nil {
			return nil, err
		}
		if randomList != nil && len(randomList) > 0 {
			unionAllList = append(unionAllList, randomList...)
		}
	}

	return unionAllList, nil
}

// QueryTeamStatisticByActivityId 根据活动ID查询团队统计信息
// 对应Java中的queryTeamStatisticByActivityId方法
func (s *IndexService) QueryTeamStatisticByActivityId(ctx context.Context, activityId int64) (*model.TeamStatisticVO, error) {
	return s.activityRepository.QueryTeamStatisticByActivityId(ctx, activityId)
}
