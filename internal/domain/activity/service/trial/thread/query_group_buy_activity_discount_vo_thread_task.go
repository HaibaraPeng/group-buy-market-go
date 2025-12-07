package thread

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// QueryGroupBuyActivityDiscountVOThreadTask 查询营销配置任务
// 对应Java中的Callable<GroupBuyActivityDiscountVO>接口实现
type QueryGroupBuyActivityDiscountVOThreadTask struct {
	source             string
	channel            string
	activityRepository repository.ActivityRepository
}

// NewQueryGroupBuyActivityDiscountVOThreadTask 创建查询营销配置任务实例
func NewQueryGroupBuyActivityDiscountVOThreadTask(
	source string,
	channel string,
	activityRepository repository.ActivityRepository,
) *QueryGroupBuyActivityDiscountVOThreadTask {
	return &QueryGroupBuyActivityDiscountVOThreadTask{
		source:             source,
		channel:            channel,
		activityRepository: activityRepository,
	}
}

// Call 执行查询任务，相当于Java中的call()方法
func (t *QueryGroupBuyActivityDiscountVOThreadTask) Call() (*model.GroupBuyActivityDiscountVO, error) {
	return t.activityRepository.QueryGroupBuyActivityDiscountVO(t.source, t.channel)
}

// AsyncCall 异步执行查询任务
func (t *QueryGroupBuyActivityDiscountVOThreadTask) AsyncCall() <-chan struct {
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
		result, err := t.Call()
		resultChan <- struct {
			Result *model.GroupBuyActivityDiscountVO
			Error  error
		}{Result: result, Error: err}
		close(resultChan)
	}()

	return resultChan
}
