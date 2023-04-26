package mysql

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"gorm.io/gorm"
)

func (m *MySQL) GetAllUser() []UserInfo {
	var user []UserInfo
	m.Find(&user)
	return user
}

func (m *MySQL) UserInfo(username string) UserInfo {
	var user UserInfo
	m.Where("user_name = ?", username).First(&user)
	return user
}

func (m *MySQL) FindUserPassword(ruser string) (user UserInfo, err error) {
	err = m.Where("user_name = ?", ruser).Find(&user).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		logger.Error("Mysql Find user password error:", err)
	}
	return
}

func (m *MySQL) FindEmail(email string) bool {
	var user UserInfo
	if m.Where("email = ?", email).First(&user).RowsAffected == 0 {
		return false
	}
	return true
}

func (m *MySQL) FindUser(ruser string) bool {
	var user UserInfo
	if m.Where("user_name = ?", ruser).First(&user).RowsAffected == 0 {
		return false
	}
	return true
}
func (m *MySQL) UpdateUserPassword(username string, password string) bool {
	var user *UserInfo
	if err := m.Model(user).Where("user_name = ?", username).Update("password", password).Error; err != nil {
		logger.Error("Mysql update user password error: ", err)
		return false
	}
	return true
}
func (m *MySQL) UpdateUserType(username string, usertype string) bool {
	var user *UserInfo
	if err := m.Model(user).Where("user_name = ?", username).Update("user_type", usertype).Error; err != nil {
		logger.Error("Mysql update user type error: ", err)
		return false
	}
	return true
}
func (m *MySQL) ExistUserId(id int) bool {
	var user *UserInfo
	if err := m.Model(user).Where("id = ?", id).First(&user).Error; err != nil {
		logger.Error("Mysql exist user id error: ", err)
		return false
	}
	return true
}
func (m *MySQL) ExistUserName(username string) bool {
	var user *UserInfo
	if err := m.Model(user).Where("user_name = ?", username).First(&user).Error; err != nil {
		logger.Error("Mysql exist user name error: ", err)
		return false
	}
	return true
}
func (m *MySQL) GetUserType(userid int) string {
	var user *UserInfo
	m.Where("id = ?", userid).First(&user)
	return user.UserType
}
func (m *MySQL) DelUser(userid int) bool {
	var user *UserInfo
	if err := m.Model(user).Where("id = ?", userid).Delete(&user).Error; err != nil {
		logger.Error("Mysql del user error:", err)
		return false
	}
	return true
}
func (m *MySQL) CreatUser(nick_name, email, password string) bool {
	var usertype string
	if nick_name == "iguidao" {
		usertype = model.USERTYPEADMIN
	} else {
		usertype = model.USERTYPEVISITOR
	}
	if result := m.Create(&UserInfo{
		UserName: nick_name,
		Email:    email,
		Password: password,
		UserType: usertype,
		Enable:   true,
	}); result.Error != nil {
		logger.Error("Mysql create mysql user fails", result.Error)
		return false
	}
	return true
}
