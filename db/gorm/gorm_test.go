package mgorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

func Test_gorm(t *testing.T) {
	db, err := gorm.Open("mysql", "user:123456@/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	//自动检查 Product 结构是否变化，变化则进行迁移
	db.AutoMigrate(&Product{})

	// 增
	//db.Create(&Product{Code: "L1214", Price: 1004})

	//// 查
	//var product Product
	//db.First(&product, 1) // 找到id为1的产品

	//db.First(&product, "code = ?", "L1212") // 找出 code 为 l1212 的产品
	//
	//// 改 - 更新产品的价格为 2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删 - 删除产品
	//p := Product{
	//	Code: "L1212",
	//	Price: 1000,
	//}
	//p.ID = 4
	// 根据主键删除，如果主键没值将删除所有记录，这点有点坑
	//d := db.Delete(&p)
	//fmt.Println(d.RowsAffected)

}
