package role_service

import (
	"sync/pkg/enforcer"
	"sync/pkg/logger"
	"sync/pkg/model"
	"sync/pkg/model/permission"
	"sync/pkg/types"
)

// AddPermissionsByRole 为角色添加权限
func AddPermissionsByRole(roleId uint64, ids []interface{}) (err error) {
	//1.删除现有权限
	e := enforcer.Enforcer
	id := types.UInt64ToString(roleId)
	_, err = e.DeletePermissionsForUser(id)
	if err != nil {
		logger.Danger(err, "AddPermissions DeletePermissionsForUser Error")
		return
	}

	//2.查找所有需添加权限
	var permissionIds []uint64
	for _, permissionId := range ids {
		permissionIds = append(permissionIds, uint64(permissionId.(float64)))
	}
	var permissionList []permission.Permission
	model.DB.Find(&permissionList, permissionIds)

	//3.添加权限
	for _, val := range permissionList {
		_, err = e.AddPermissionForUser(id, val.Url, val.Method)
		if err != nil {
			logger.Danger(err, "AddPermissions AddPermissionForUser Error")
			return
		}
	}

	return nil
}

// GetAllPermission 获取角色所有权限
func GetAllPermission(roleId uint64) (permissionList []permission.Permission, err error) {
	err = model.DB.Model(&permission.Permission{}).Joins("JOIN casbin_rule ON url=v1 AND method=V2").Where("v0 = ?",types.UInt64ToString(roleId)).Find(&permissionList).Error
	return
}
