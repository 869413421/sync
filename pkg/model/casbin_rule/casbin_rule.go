package casbin_rule

var CasbinRuleTable = "casbin_rule"

type CasbinRule struct {
	ID    uint64 `valid:"id"`
	Ptype string `valid:"ptype"`
	V0    string `valid:"v0"`
	V1    string `valid:"v1"`
	V2    string `valid:"v2"`
	V3    string `valid:"v3"`
	V4    string `valid:"v4"`
	V5    string `valid:"v5"`
}
