package model

import (
	"gorm.io/gorm"
	"time"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 10:42
 *
 * Description: 用户
 **/

type YanwuUser struct {
	// --- 声明该struct是gorm模型
	gorm.Model
	ID        uint      `gorm:"primaryKey"`             // 通过primaryKey指定主键
	Name      string    `gorm:"size:32;index:idx_name"` // 用户名，通过size指定长度，通过index指定一个名为idx_name的索引
	Email     string    `gorm:"uniqueIndex:un_email"`   // 用户邮箱，通过index指定一个名为un_email的唯一索引
	Age       uint64    `gorm:"check:age > 100"`        // 年龄
	Password  string    `gorm:"size:32;"`               // 用户密码
	CreatedAt time.Time `gorm:""`                       // 创建时间，只读
	UpdatedAt time.Time `gorm:""`                       // 更新时间，允许读和更新
}

// TableName 设置表名，默认是结构体的名的复数形式
func (YanwuUser) TableName() string {
	return "YANWU_USER"
}
