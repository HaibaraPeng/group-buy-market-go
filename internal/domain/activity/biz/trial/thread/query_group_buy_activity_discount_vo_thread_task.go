package thread

import (
	"context"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// QueryGroupBuyActivityDiscountVOThreadTask 查询营销配置任务
// 对应Java中的Callable<GroupBuyActivityDiscountVO>接口实现
type QueryGroupBuyActivityDiscountVOThreadTask struct {
	source             string
	channel            string
	goodsId            string
	activityRepository *repository.ActivityRepository
}

// NewQueryGroupBuyActivityDiscountVOThreadTask 创建查询营销配置任务实例
func NewQueryGroupBuyActivityDiscountVOThreadTask(
	source string,
	channel string,
	goodsId string,
	activityRepository *repository.ActivityRepository,
) *QueryGroupBuyActivityDiscountVOThreadTask {
	return &QueryGroupBuyActivityDiscountVOThreadTask{
		source:             source,
		channel:            channel,
		goodsId:            goodsId,
		activityRepository: activityRepository,
	}
}

// Call 执行查询任务，相当于Java中的call()方法
func (t *QueryGroupBuyActivityDiscountVOThreadTask) Call(ctx context.Context) (*model.GroupBuyActivityDiscountVO, error) {
	// 查询渠道商品活动配置关联配置
	scSkuActivityVO, err := t.activityRepository.QuerySCSkuActivityBySCGoodsId(ctx, t.source, t.channel, t.goodsId)
	if err != nil {
		return nil, err
	}

	if scSkuActivityVO == nil {
		return nil, nil
	}

	// 查询活动配置
	return t.activityRepository.QueryGroupBuyActivityDiscountVO(ctx, scSkuActivityVO.ActivityId)
}

// AsyncCall 异步执行查询任务
func (t *QueryGroupBuyActivityDiscountVOThreadTask) AsyncCall(ctx context.Context) <-chan struct {
	Result *model.GroupBuyActivityDiscountVO
	Error  error
} {
	// 创建通道用于返回结果
	resultChan := make(chan struct {
		Result *model.GroupBuyActivityDiscountVO
		Error  error
	}, 1)

	// 启动goroutine异步执行
	go func() {
		result, err := t.Call(ctx)
		resultChan <- struct {
			Result *model.GroupBuyActivityDiscountVO
			Error  error
		}{Result: result, Error: err}
		close(resultChan)
	}()

	return resultChan
}
