package casbin_service

import (
	"fmt"
	"sync/pkg/model/casbin_rule"
)

func GetPerssionTree()  {
	where :=make(map[string]interface{})
	where["ptype"]="p"
	where["v0"]=""
	rules:=casbin_rule.GetList(where)
	fmt.Println(rules)
}

