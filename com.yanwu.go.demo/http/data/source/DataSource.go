package source

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 16:22
 *
 * Description:
 **/

var DataSource *gorm.DB

func InitDataSource() {
	// ----- 初始化数据库连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:yanwu12138@tcp(192.168.56.150:3306)/go_db_test?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                          // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	DataSource = db
	fmt.Println("database init success:")
}
