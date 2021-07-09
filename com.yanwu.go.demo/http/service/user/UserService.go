package user

import (
	mapper "../../data/mapper"
	model "../../data/model"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/7/9 10:49
 *
 * Description: 用户
 **/

func Create(user *model.YanwuUser) uint {
	return mapper.SaveUser(user)
}
