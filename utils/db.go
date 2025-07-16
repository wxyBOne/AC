package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// Init 初始化数据库连接
func Init(dsn string) error {
	// 使用gorm打开数据库连接
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
