package requests

import (
	"github.com/thedevsaddam/govalidator"
	"sync/pkg/model/casbin_rule"
)

func ValidateRuleEditForm(data casbin_rule.CasbinRule) map[string][]string {
	rules := govalidator.MapData{
		"ptype": []string{
			"required",
			"alpha_num",
			"between:1,10",
		},
		"v0": []string{
			"required",
			"alpha_num",
			"between:1,30",
		},
		"v1": []string{
			"required",
			"alpha_num",
			"between:1,30",
		},
		"v2": []string{
			"alpha_num",
			"between:1,30",
		},
		"v3": []string{
			"alpha_num",
			"between:1,30",
		},
		"v4": []string{
			"alpha_num",
			"between:1,30",
		},
		"v5": []string{
			"alpha_num",
			"between:1,30",
		},
	}

	messages := govalidator.MapData{
		"ptype": []string{
			"required：规则类型为必填",
			"alpha_num:只允许数字和英文",
			"between:用户名在1到10个字符之间",
			"not_exists：用户名已经存在",
		},
		"v0": []string{
			"required：规则类型为必填",
			"alpha_num:只允许数字和英文",
			"between:用户名在1到10个字符之间",
			"not_exists：用户名已经存在",
		},
		"v1": []string{
			"required：规则类型为必填",
			"alpha_num:只允许数字和英文",
			"between:用户名在1到10个字符之间",
			"not_exists：用户名已经存在",
		},
		"v2": []string{
			"required：规则类型为必填",
			"alpha_num:只允许数字和英文",
			"between:用户名在1到100个字符之间",
			"not_exists：用户名已经存在",
		},
		"v3": []string{
			"required：规则类型为必填",
			"alpha_num:只允许数字和英文",
			"between:用户名在1到100个字符之间",
		},
		"v4": []string{
			"alpha_num:只允许数字和英文",
			"between:用户名在1到100个字符之间",
		},
		"v5": []string{
			"alpha_num:只允许数字和英文",
			"between:用户名在1到100个字符之间",
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
