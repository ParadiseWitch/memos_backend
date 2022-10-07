package main

import (
	"fmt"
	"memos/server/dto"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	url := "root:123456@tcp(maiiiiiid.fun:23306)/memos?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"

	DB, err := gorm.Open("mysql", url)
	if err != nil {
		fmt.Print(err)
	}
	// DB.DB().SetMaxIdleConns(10)
	// DB.SingularTable(true)
	// var u dto.User
	// r := map[string]interface{}{
	// 	// "id":   1,
	// 	"name": "admin",
	// }

	var u dto.User
	rset := DB.Table("user").Where("name = ?", "xxx")
	// count := 0
	// rset.Count(&count)
	if rset.Find(&u).RecordNotFound() {
		print("not found")
		return
	}
	fmt.Print(u)
}
