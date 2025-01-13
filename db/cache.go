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
			_, err := m.AddCache(cacheKey, "")
			return nil, err
		}
		return nil, tx.Error
	}
	return &cache, nil
}

func (m *MyDb) AddCache(cacheKey string, cacheValue string) (bool, error) {
	cache := model.Cache{}
	tx := m.Db.Where(&model.Cache{CacheKey: cacheKey}).First(&cache)
	if tx.Error != nil {
		if errors2.Is(tx.Error, gorm.ErrRecordNotFound) {
			tx = m.Db.Create(&model.Cache{
				CacheKey:   cacheKey,
				CacheValue: cacheValue,
			})
			if tx.Error != nil {
				return false, tx.Error
			}
			return true, nil
		}
		return false, tx.Error
	}
	return false, nil
}

func (m *MyDb) UpdateCache(cacheKey string, cacheValue string) (bool, error) {
	if cacheByKey, err := m.FindCacheByKey(cacheKey); err != nil {
		return false, err
	} else {
		if cacheByKey == nil {
			return false, errors.New(errors.ParameterError)
		}
	}

	tx := m.Db.Where(&model.Cache{
		CacheKey: cacheKey,
	}).Updates(model.Cache{
		CacheValue: cacheValue,
	})
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
