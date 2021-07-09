package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/14 14:52
 *
 * Description: 用户
 **/

type GormUser struct {
	gorm.Model                  // 声明该struct是gorm模型
	ID           uint           `gorm:"primaryKey"` // 通过primaryKey指定主键
	Name         string         `gorm:"size:32;index:idx_name"` // 用户名，通过size指定长度，通过index指定一个名为idx_name的索引
	Email        string         `gorm:"uniqueIndex:un_email"` // 用户邮箱，通过index指定一个名为un_email的唯一索引
	Age          uint8          `gorm:"check:age > 100"` // 年龄
	Birthday     *time.Time     // 出生日期
	MemberNumber sql.NullString // 成员数量
	ActivatedAt  sql.NullTime   `gorm:"->;<-:create"` // 激活时间，允许读和创建，但不允许更改
	CreatedAt    time.Time      `gorm:"->"`           // 创建时间，只读
	UpdatedAt    time.Time      `gorm:"->;<-:update"` // 更新时间，允许读和更新
}

// TableName 设置表名，默认是结构体的名的复数形式
func (GormUser) TableName() string {
	return "GORM_USER"
}
