package requests

import (
	"github.com/thedevsaddam/govalidator"
	"sync/pkg/model/role"
	"sync/pkg/types"
)

func ValidateRole(data role.Role) map[string][]string {
	rules := govalidator.MapData{
		"name": []string{
			"required",
			"alpha_num",
			"between:2,255",
			"not_exists:roles,name," + types.UInt64ToString(data.ID),
		},
		"desc": []string{
			"required",
			"between:2,255",
		},
		"order": []string{
			"numeric",
			"min:0",
		},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required：名称为必填",
			"between: 在2到255个字符之间",
			"alpha_num:必须是英文",
			"not_exists：角色已经存在",
		},
		"desc": []string{
			"required：规则类型为必填",
			"between:简介在2到255个字符之间",
		},
		"order": []string{
			"required：规则类型为必填",
			"numeric:排序只允许为数字",
		},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	errs := govalidator.New(opts).ValidateStruct()

	return errs
}
