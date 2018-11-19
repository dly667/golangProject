package gormConnection

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {


	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//Migrate the schema
	db.AutoMigrate(&Product{})

	//创建
	//db.Create(&Product{Code:"L22",Price:1000})

	//读取
	var product Product
	db.First(&product, 1)
	//db.First(&product,"code=?","L1212")
	//
	fmt.Println(product)
	//
	//更新 - 更新product得price为2000
	db.Model(&product).Update("Price", 2000)

	db.First(&product, 1)
	fmt.Println(product)
	//
	////删除 - 删除product
	//db.Delete(&product)
}
