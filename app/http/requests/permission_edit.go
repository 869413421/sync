package requests

import (
	"github.com/thedevsaddam/govalidator"
	"gorm.io/gorm"
	"sync/pkg/logger"
	"sync/pkg/model/permission"
)

func ValidatePermission(data permission.Permission) map[string][]string {
	rules := govalidator.MapData{
		"url": []string{
			"required",
			"between:2,255",
		},
		"name": []string{
			"required",
			"between:2,255",
		},
		"method": []string{
			"required",
			"in:GET,POST,PUT,DELETE,HEAD,OPTIONS",
			"between:2,255",
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
		"url": []string{
			"required：URL类型为必填",
			"between:URL在2到255个字符之间",
		},
		"name": []string{
			"required：名称为必填",
			"between: 在2到255个字符之间",
			"not_exists：用户名已经存在",
		},
		"method": []string{
			"required：请求类型类型为必填",
			"in:选项错误",
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
	where := map[string]interface{}{
		"url":    data.Url,
		"name":   data.Name,
		"method": data.Method,
	}

	permission, err := permission.GetByWhere(where)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Danger(err, "ValidatePermission Error")
	}

	if permission.ID > 0 && permission.ID != data.ID {
		errs["url"] = append(errs["url"], "权限已经存在")
	}

	return errs
}
