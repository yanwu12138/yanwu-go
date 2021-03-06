package mapper

import (
	model "../../data/model"
	dataSource "../source"
	"log"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 10:48
 * @see
 * Description:
 **/

func SaveUser(user *model.YanwuUser) uint {
	create := dataSource.DataSource.Create(user)
	if create.Error != nil {
		log.Println("save user error.", create.Error)
		return 0
	}
	return user.ID
}

func AllUser() []model.YanwuUser {
	var result []model.YanwuUser
	dataSource.DataSource.Find(&result)
	return result
}

func GetById(id uint) model.YanwuUser {
	var user model.YanwuUser
	dataSource.DataSource.First(&user, id)
	return user
}
