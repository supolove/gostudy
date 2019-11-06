package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main2() {
	read, _ := ioutil.ReadFile("resources/config.json")
	var f interface{}
	err := json.Unmarshal(read, &f)
	if err != nil {
		fmt.Println(err)
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int")
		case []interface{}:
			fmt.Println(k, "is an array")
		default:
			fmt.Println(k, "is unknown")

		}
		//result := fmt.Sprintf("%s:%v",k,v)
		//fmt.Println(result)
	}

	//decoder := json.NewDecoder(f)
	//
	//for t,e:= decoder.Token(); e == nil ; t,e = decoder.Token(){
	//	fmt.Println(t)
	//}
}
