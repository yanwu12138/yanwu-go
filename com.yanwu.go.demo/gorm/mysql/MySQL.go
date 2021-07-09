package main

import (
	"../model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/14 15:48
 *
 * Description:
 **/

var myDb *gorm.DB

func init() {
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
	myDb = db
	fmt.Println("database init success:")
}

func main() {
	//testCreateUser()
	testSelectUser()
}

func testCreateUser() {
	fmt.Println("------------- 插入数据 -------------")
	// ----- 插入数据
	user := model.GormUser{Name: "yanwu", Email: "yanwu0527@163.com", Age: 29}
	result := myDb.Create(&user)
	fmt.Println("RowsAffected:", result.RowsAffected, ", userId: ", user.ID)

	fmt.Println("------------- 批量插入 -------------")
	// ----- 批量插入
	users := []model.GormUser{{Name: "yanwu", Email: "yanwu0527@163.com", Age: 29}, {Name: "lotus", Email: "499496273@qq.com", Age: 29}, {Name: "wenxin", Age: 3}, {Name: "wenfu", Age: 2}}
	results := myDb.CreateInBatches(users, len(users))
	for _, user := range users {
		fmt.Println("RowsAffected:", results.RowsAffected, ", userId: ", user.ID)
	}

	fmt.Println("------------- 根据Map插入 -------------")
	userMap := []map[string]interface{}{
		{"Name": "yanwu", "Email": "yanwu0527@163.com", "Age": 29},
		{"Name": "lotus", "Email": "499496273@qq.com", "Age": 29},
		{"Name": "wenxin", "Email": "wenxin@qq.com", "Age": 3},
		{"Name": "wenfu", "Email": "wenfu@qq.com", "Age": 2},
	}
	resultMap := myDb.Model(&model.GormUser{}).CreateInBatches(userMap, len(userMap))
	fmt.Println("RowsAffected:", resultMap.RowsAffected)
}

func testSelectUser() {
	fmt.Println("------------- 获取第一条记录（主键升序） -------------")
	var userFirst model.GormUser
	myDb.First(&userFirst)
	fmt.Printf("userFirst: %+v\n", userFirst)

	fmt.Println("------------- 获取一条记录，没有指定排序字段 -------------")
	var userTake model.GormUser
	myDb.Take(&userTake)
	fmt.Printf("userTake: %+v\n", userTake)

	fmt.Println("------------- 获取最后一条记录（主键降序） -------------")
	var userLast model.GormUser
	myDb.Last(&userLast)
	fmt.Printf("userLast: %+v\n", userLast)

	fmt.Println("------------- 根据主键检索数据 -------------")
	var userById model.GormUser
	myDb.First(&userById, 10)
	fmt.Printf("userById: id: %v > %+v\n", userById.ID, userById)

	fmt.Println("------------- 根据主键集合检索数据 -------------")
	var usersById []model.GormUser
	myDb.Find(&usersById, []uint{12, 15, 17})
	for _, user := range usersById {
		fmt.Printf("userByIds: id: %v > %+v\n", user.ID, user)
	}

	fmt.Println("------------- 查询所有数据 -------------")
	var userAll []model.GormUser
	myDb.Find(&userAll)
	for _, user := range userAll {
		fmt.Printf("userAll: id: %v > %+v\n", user.ID, user)
	}

	fmt.Println("------------- 根据简单条件查询 -------------")
	var usersByName []model.GormUser
	myDb.Where("name = ?", "yanwu").Find(&usersByName)
	for _, user := range usersByName {
		fmt.Printf("usersByName: id: %v > %+v\n", user.ID, user)
	}

	fmt.Println("------------- 根据多个条件查询 -------------")
	var usersByNameAndAge []model.GormUser
	myDb.Where("name = ? and age = ?", "lotus", 29).Find(&usersByNameAndAge)
	for _, user := range usersByNameAndAge {
		fmt.Printf("usersByNameAndAge: id: %v > %+v\n", user.ID, user)
	}

	fmt.Println("------------- 根据Struct进行查询 -------------")
	// 注意 当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条。
	var usersByStruct []model.GormUser
	myDb.Where(&model.GormUser{Name: "wenxin"}).Find(&usersByStruct)
	for _, user := range usersByStruct {
		fmt.Printf("usersByStruct: id: %v > %+v\n", user.ID, user)
	}

	fmt.Println("------------- 根据Map进行查询 -------------")
	var usersByMap []model.GormUser
	myDb.Where(map[string]interface{}{"name": "wenfu"}).Find(&usersByMap)
	for _, user := range usersByMap {
		fmt.Printf("usersByMap: id: %v > %+v\n", user.ID, user)
	}


}
