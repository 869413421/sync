package sync_rule

import "sync/pkg/model"

func GetAll() (rules []SyncRule, err error) {
	err = model.DB.Model(SyncRule{}).Order("created_at desc").Find(&rules).Error
	return
}
