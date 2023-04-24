package mysql

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"gorm.io/gorm"
)

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

func (m *MySQL) CreatUser(nick_name, email, usertype, password string) bool {
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
