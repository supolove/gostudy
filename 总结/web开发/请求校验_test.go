package web开发

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func emailFormatValid(email string) bool {
	return true
}

func createUser() {

}

// json,validate标签
// 一般通过反射获取
// vt.Field(i).Tag.Get("validate")
// 使用Go内置的Parser对源代码进行扫描
type RegisterReq struct {
	Username       string `json:"username"         validate:"gt=0"`
	PasswordNew    string `json:"password_new"     validate:"gt=0"`
	PasswordRepeat string `json:"password_repeat"  validate:"eqfield=PasswordNew"`
	Email          string `json:"email"            validate:"email"`
}

func register(req RegisterReq) error {
	if len(req.Username) > 0 {
		if len(req.PasswordNew) > 0 && len(req.PasswordRepeat) > 0 {
			if req.PasswordNew == req.PasswordRepeat {
				if emailFormatValid(req.Email) {
					createUser()
					return nil
				}
				return errors.New("invalid email")
			} else {
				return errors.New("password and reinput must be the same")
			}
		} else {
			return errors.New("password and password reinput must be longer than 0")
		}
	} else {
		return errors.New("length of username cannot be 0")
	}
}

// 优化
func register2(req RegisterReq) error {
	if len(req.Username) == 0 {
		return errors.New("length of username cannot be 0")
	}

	if len(req.PasswordNew) == 0 || len(req.PasswordRepeat) == 0 {
		return errors.New("password and password reinput must be longer than 0")
	}

	if req.PasswordNew != req.PasswordRepeat {
		return errors.New("password and reinput must be the same")
	}

	if !emailFormatValid(req.Email) {
		return errors.New("invalid email")
	}

	createUser()
	return nil
}

// 用validator解放体力劳动
// 引入库https://github.com/go-playground/validator
func TestRequestVerify(t *testing.T) {

	req := http.Request{}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		createUser()
	}

}

//  validator 树
//原理很简单，就是用反射对结构体进行树形遍历。
type Nested struct {
	Email string `validate:"email"`
}
type T struct {
	Age    int `validate:"eq=10"`
	Nested Nested
}

func validateEmail(input string) bool {
	if pass, _ := regexp.MatchString(
		`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, input,
	); pass {
		return true
	}
	return false
}

func validate(v interface{}) (bool, string) {
	validateResult := true
	errmsg := "success"
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		tagContent := vt.Field(i).Tag.Get("validate")
		k := fieldVal.Kind()

		switch k {
		case reflect.Int:
			val := fieldVal.Int()
			tagValStr := strings.Split(tagContent, "=")
			tagVal, _ := strconv.ParseInt(tagValStr[1], 10, 64)
			if val != tagVal {
				errmsg = "validate int failed, tag is: " + strconv.FormatInt(
					tagVal, 10,
				)
				validateResult = false
			}
		case reflect.String:
			val := fieldVal.String()
			tagValStr := tagContent
			switch tagValStr {
			case "email":
				nestedResult := validateEmail(val)
				if nestedResult == false {
					errmsg = "validate mail failed, field val is: " + val
					validateResult = false
				}
			}
		case reflect.Struct:
			// 如果有内嵌的 struct，那么深度优先遍历
			// 就是一个递归过程
			valInter := fieldVal.Interface()
			nestedResult, msg := validate(valInter)
			if nestedResult == false {
				validateResult = false
				errmsg = msg
			}
		}
	}
	return validateResult, errmsg
}

func TestRequestVerify2(t *testing.T) {

	var a = T{Age: 11, Nested: Nested{Email: "abc@abc.com"}}

	validateResult, errmsg := validate(a)
	fmt.Println(validateResult, errmsg)

}
