package db

import (
	"loan-server/common/consts"
	"strconv"
)

func (m *MyDb) GetLeoBlockHeight(cacheKey string) (int64, error) {

	cache, err := m.FindCacheByKey(consts.LeoBlockHeightKey)
	if err != nil {
		return 0, err
	}
	blockNum, err := strconv.ParseInt(cache.CacheValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}
