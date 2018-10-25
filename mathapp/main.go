package main

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
	//fmt.Printf("Hello world. Sqrt(2) = %v\n",mymath.Sqrt(2))
	//fmt.Print(mymath.Test())
	//mymath.Sum()
	//mymath.ShowYangHuiTriangle(10)
	//var a,b = GoBase.Function1()
	//fmt.Println(a,b)
	//GoBase.Test001()
	//GoBase.Test04Main()

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
