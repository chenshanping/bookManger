package dao

import (
	"bookManage/global"
	"bookManage/models"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
)

//type UserDao interface {
//	PassWordCrypto(pwd string) string
//	CreateUser()
//	GetUserData() models.User
//	GetUserid() uint
//	GetUserList() []models.User
//	DeleteUser()
//	SetUserToken()
//}

func PassWordCrypto(pwd string) string {
	b := md5.Sum([]byte(pwd)) // 加密数据
	return fmt.Sprintf("%x", b)
}

func CreateUser(p *models.User) {
	global.DB.Create(p)
}
func GetUserData(name string) models.User {
	user := models.User{}
	global.DB.Where("username=?", name).First(&user)
	return user
}
func GetUserid(name string) uint {

	//user := models.User{}
	//global.DB.Where("username=?", name).First(&user)
	if GetUserData(name).ID == 0 {
		return 0
	}
	return GetUserData(name).ID
}

func GetUserList() []models.User {
	var users []models.User
	global.DB.Find(&users)
	return users

}

func DeleteUser(value string) {
	var user models.User
	//global.DB.Model(&user).Delete(value)
	global.DB.Where("username=?", value).Delete(&user)
}

func SetUserToken(p *models.User) {
	token := uuid.New().String()
	global.DB.Model(p).Update("token", token)
}
