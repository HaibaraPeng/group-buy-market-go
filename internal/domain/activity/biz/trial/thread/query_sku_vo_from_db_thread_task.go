package thread

import (
	"context"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// QuerySkuVOFromDBThreadTask 查询商品信息任务
// 对应Java中的Callable<SkuVO>接口实现
type QuerySkuVOFromDBThreadTask struct {
	goodsId            string
	activityRepository *repository.ActivityRepository
}

// NewQuerySkuVOFromDBThreadTask 创建查询商品信息任务实例
func NewQuerySkuVOFromDBThreadTask(
	goodsId string,
	activityRepository *repository.ActivityRepository,
) *QuerySkuVOFromDBThreadTask {
	return &QuerySkuVOFromDBThreadTask{
		goodsId:            goodsId,
		activityRepository: activityRepository,
	}
}

// Call 执行查询任务，相当于Java中的call()方法
func (t *QuerySkuVOFromDBThreadTask) Call(ctx context.Context) (*model.SkuVO, error) {
	return t.activityRepository.QuerySkuByGoodsId(ctx, t.goodsId)
}

// AsyncCall 异步执行查询任务
func (t *QuerySkuVOFromDBThreadTask) AsyncCall(ctx context.Context) <-chan struct {
	Result *model.SkuVO
	Error  error
} {
	// 创建通道用于返回结果
	resultChan := make(chan struct {
		Result *model.SkuVO
		Error  error
	}, 1)

	// 启动goroutine异步执行
	go func() {
		result, err := t.Call(ctx)
		resultChan <- struct {
			Result *model.SkuVO
			Error  error
		}{Result: result, Error: err}
		close(resultChan)
	}()

	return resultChan
}
