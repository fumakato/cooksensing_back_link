package database

import "myapp/model"

// データ全件取得
func FindAllFeatureDataLinks() ([]model.FeatureDataLink, error) {
	var list []model.FeatureDataLink
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
