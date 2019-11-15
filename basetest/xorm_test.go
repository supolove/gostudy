package basetest

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"testing"
	"xorm.io/core"
)

type AccountXone struct {
	Id      int64  `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

type AccountXtow struct {
	Id      int64  `xorm:"not null pk autoincr INT(11)"`
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

var x *xorm.Engine

// 增
func addXorm() {
	_, err := x.Insert(
		&AccountXone{
			Name:    "xiaoqiang",
			Balance: 12,
		},
		&AccountXone{
			Name:    "xiaoming",
			Balance: 25,
		},
		&AccountXone{
			Name:    "haier",
			Balance: 8,
		},
		&AccountXone{
			Name:    "fengqingyang",
			Balance: 99,
		},
		&AccountXone{
			Name:    "shudaqiang",
			Balance: 36,
		},
	)

	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}

// 删
func deleteXorm() {
	_, err := x.Delete(&AccountXone{Id: 1})

	if err != nil {
		panic(err)
	}
}

// 查
func selectXorm1() {
	a := &AccountXone{Id: 1}
	b, err := x.Get(a)

	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

// 查2
func selectXorm2() {
	a := &AccountXone{}
	b, err := x.Id(1).Get(a)

	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

// 查3
func selectXorm3() {
	a := &AccountXone{}
	b, err := x.Where("Id=?", 1).Get(a)

	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

// 查批量
func seletXorm4() {
	as := []AccountXone{}
	err := x.Desc("balance").Find(&as)
	if err != nil {
		panic(err)
	}
	fmt.Println(as)
}

// 查5,多值查询
func seletXorm5() {
	a := []AccountXone{}
	err := x.In("Id", 1, 2, 3).Find(&a)

	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

// Cols 特定字段查询
func seletXorm6() {
	a := []AccountXone{}
	err := x.Cols("Id", "Name").Find(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

// 更新数据
// 默认开始了乐观锁
func updateXorm1() {

	// 必须要有Version
	a := AccountXone{
		Id:      2,
		Name:    "xiaoming2",
		Balance: 0,
		Version: 2,
	}

	//count,err := x.Id(1).Update(&a)
	count, err := x.ID(2).Update(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

// 更新数据
func updateXorm2() {
	affected, err := x.Exec("update account set name = ? where id = ?", "xiaoqiang1", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(affected)
}

// Cols 更新数据
func updateXorm3() {
	a := AccountXone{
		Name:    "haier1",
		Version: 2,
	}
	count, err := x.Id(3).Cols("name").Update(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

// 事物
func shiwuXorm1() {
	session := x.NewSession()
	defer session.Close()

	err := session.Begin()

	///// 插入数据 /////
	a := AccountXone{
		Name:    "linhuchong",
		Balance: 96,
	}
	var cout int64
	cout, err = session.Insert(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(cout)
	///// 插入数据 /////

	err = session.Commit()
	if err != nil {
		panic(err)
	}

	// 回滚
	/*
		if err != nil {
			session.Rollback()
		}
	*/
}

func TestXorm(t *testing.T) {
	var err error
	x, err = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	// 结构体名称和表名映射
	// 一样名字
	//x.SetMapper(core.SameMapper{})
	//
	//x.SetMapper(core.SnakeMapper{})
	// 映射带前缀
	x.SetMapper(core.PrefixMapper{
		Mapper: core.SnakeMapper{},
		Prefix: "test_",
	})

	if err = x.Sync2(new(AccountXone)); err != nil {
		panic(err)
	}

	//gonicNames := []string{"SSL"}
	//for _, name := range gonicNames {
	//	core.LintGonicMapper[name] = true
	//}

	fmt.Println(x.TableName(new(AccountXone)))

	//shiwuXorm1()
	//updateXorm3()
	//seletXorm4()
}
