package helper

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:@/grpc-ecom?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
