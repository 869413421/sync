package casbin_rule

var CasbinRuleTable = "casbin_rule"

type CasbinRule struct {
	Ptype string `valid:"ptype" json:"ptype"`
	V0    string `valid:"v0" json:"v_0"`
	V1    string `valid:"v1" json:"v_1"`
	V2    string `valid:"v2" json:"v_2"`
	V3    string `valid:"v3" json:"v_3"`
	V4    string `valid:"v4" json:"v_4"`
	V5    string `valid:"v5" json:"v_5"`
	Name  string `valid:"name" json:"name"`
	DESC  string `valid:"desc" json:"desc"`
}
