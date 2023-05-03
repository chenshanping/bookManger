package service

import (
	"bookManage/global"
	"bookManage/model"
)

var UserService = userService{}

type userService struct {
}

// GetAllUser 获取用户列表
func (usrv userService) GetAllUser(pageSize int, offsetval int) (userList []*model.User, total int64) {
	global.DB.Model(userList).Count(&total).Limit(pageSize).Offset(offsetval).Find(&userList)
	return
}

// CreateUser 添加用户
func (usrv userService) CreateUser(user *model.User) (err error) {
	if err := global.DB.Create(user).Error; err != nil {
		return err
	}
	return
}

// FindByUsername 根据账号查找用户
func (usrv userService) FindByUsername(username string) (user model.User) {
	global.DB.Where("username=?", username).First(&user)
	return
}
func (usrv userService) GetUserId(name string) uint {
	username := usrv.FindByUsername(name)
	return username.ID

}
func (usrv userService) IDExist(id string) bool {
	var user model.User
	global.DB.Where("id=?", id).Find(&user)
	if user.ID == 0 {
		return false
	}
	return true

}
func (usrv userService) DeleteUser(id string) {
	var user model.User
	// 执行软删除
	global.DB.Where("id=?", id).Delete(&user)
}

func (usrv userService) UpdateUser(id string, user model.User) {
	global.DB.Where("id=?", id).Updates(&user)
}

func (usrv userService) UserDetail(name string) (user []model.User) {
	global.DB.Where("username=?", name).Find(&user)
	return user

}
