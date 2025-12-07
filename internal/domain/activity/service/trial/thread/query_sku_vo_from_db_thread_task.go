package thread

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// QuerySkuVOFromDBThreadTask 查询商品信息任务
// 对应Java中的Callable<SkuVO>接口实现
type QuerySkuVOFromDBThreadTask struct {
	goodsId            string
	activityRepository repository.ActivityRepository
}

// NewQuerySkuVOFromDBThreadTask 创建查询商品信息任务实例
func NewQuerySkuVOFromDBThreadTask(
	goodsId string,
	activityRepository repository.ActivityRepository,
) *QuerySkuVOFromDBThreadTask {
	return &QuerySkuVOFromDBThreadTask{
		goodsId:            goodsId,
		activityRepository: activityRepository,
	}
}

// Call 执行查询任务，相当于Java中的call()方法
func (t *QuerySkuVOFromDBThreadTask) Call() (*model.SkuVO, error) {
	return t.activityRepository.QuerySkuByGoodsId(t.goodsId)
}
