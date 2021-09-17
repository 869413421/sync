package role_service

import (
	"fmt"
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

func GetAllPermission(roleId uint64) (permissionList []permission.Permission, err error) {
	e := enforcer.Enforcer
	rolePermission := e.GetPermissionsForUser(types.UInt64ToString(roleId))

	err = model.DB.Model(P)Where("(url,method) in", rolePermission).Joins("join permissions on ").Find(&permissionList).Error
	fmt.Println(permissionList)
	fmt.Println(err)
	return
}
