package mgorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) BeforeCreate() {
	fmt.Println("钩子方法 before")
}

func (p *Product) AfterCreate() {
	fmt.Println("钩子方法 after")
}

//func (p *Product)Create() {
//	fmt.Println("钩子方法 create")
//}
//
//func (p *Product)Save() {
//	fmt.Println("钩子方法 save")
//}
//
//func (p *Product)Update() {
//	fmt.Println("钩子方法 Update")
//}
//
//func (p *Product)Delete() {
//	fmt.Println("钩子方法 Delete")
//}
//
//func (p *Product)Find() {
//	fmt.Println("钩子方法 Find")
//}
