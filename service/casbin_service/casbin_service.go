package casbin_service

import (
	. "sync/pkg/model/casbin_rule"
)

// GetPermissionTree 获取权限树
func GetPermissionTree() []interface{} {
	//1.获取权限列表
	where := make(map[string]interface{})
	where["ptype"] = "p"
	where["v0"] = ""
	rules := GetList(where)

	//2.构建权限树
	tree := getTreeByList(0, rules)

	//3.返回数据
	return tree
}

// getTreeByList 根据list构建权限树
func getTreeByList(parentId uint64, rules []CasbinRule) []interface{} {
	//1.创建一个切片
	var data []interface{}

	//2.递归构建新数据
	for _, rule := range rules {
		if rule.ParentId == parentId {
			item := make(map[string]interface{})
			item["value"] = rule.ID
			item["label"] = rule.Name
			item["children"] = getTreeByList(rule.ID, rules)
			data = append(data, item)
		}
	}

	//3.返回数据
	return data
}
