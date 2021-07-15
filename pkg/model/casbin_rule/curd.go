package casbin_rule

import (
	"github.com/gin-gonic/gin"
	"sync/pkg/model"
	"sync/pkg/pagination"
)

// Pagination 获取所有用户
func Pagination(ctx *gin.Context, where map[string]interface{}, perPage int) (rules []CasbinRule, viewData pagination.ViewData, err error) {
	//1.初始化分页实例
	db := model.DB.Model(CasbinRule{}).Order("created_at desc")
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	_pager := pagination.New(ctx, db, "/user", perPage)

	// 2. 获取视图数据
	viewData = _pager.Paging()

	// 3. 获取数据
	_pager.Results(&rules)

	return rules, viewData, nil
}
