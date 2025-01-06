package db

import (
	errors2 "errors"
	"gorm.io/gorm"
	"loan-server/common/errors"
	"loan-server/model"
)

func (m *MyDb) FindCacheByKey(cacheKey string) (*model.Cache, error) {
	cache := model.Cache{}
	tx := m.Db.Where(&model.Cache{CacheKey: cacheKey}).First(&cache)
	if tx.Error != nil {
		if errors2.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &cache, nil
}

func (m *MyDb) AddCache(tx *gorm.DB, cache *model.Cache) (bool, error) {
	if cacheByKey, err := m.FindCacheByKey(cache.CacheKey); err != nil {
		return false, err
	} else {
		if cacheByKey != nil {
			return false, errors.New(errors.ParameterError)
		}
	}

	tx = tx.Create(&cache)
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}

func (m *MyDb) UpdateCache(tx *gorm.DB, cache *model.Cache) (bool, error) {
	if cacheByKey, err := m.FindCacheByKey(cache.CacheKey); err != nil {
		return false, err
	} else {
		if cacheByKey == nil {
			return false, errors.New(errors.ParameterError)
		}
	}

	tx = tx.Model(cache).Updates(map[string]interface{}{
		"cache_value": cache.CacheValue,
		"expired":     cache.Expired,
	})
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
