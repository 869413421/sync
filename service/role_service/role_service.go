package role_service

import (
	"sync/pkg/enforcer"
	"sync/pkg/logger"
	"sync/pkg/types"
)

func AddPermissions(roleId uint64, permissions []interface{}) (err error) {
	e := enforcer.Enforcer
	id := types.UInt64ToString(roleId)
	_, err = e.DeletePermissionsForUser(id)
	if err != nil {
		logger.Danger(err, "AddPermissions DeletePermissionsForUser Error")
		return
	}

	for _, permission := range permissions {
		_, err = e.AddPermissionForUser(id, types.Float64ToString(permission.(float64)))
		if err != nil {
			logger.Danger(err, "AddPermissions AddPermissionForUser Error")
			return
		}
	}

	return nil
}
