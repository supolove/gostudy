package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"testing"
)

func TestName(t *testing.T) {

	e := casbin.NewEnforcer("./acl_config.ini", "./acl_police.csv")

	sub := "bb"
	obj := "data1"
	act := "read"

	ok := e.Enforce(sub, obj, act)

	if ok == true {
		fmt.Println("通过")
	} else {
		fmt.Println("未通过")
	}
}
