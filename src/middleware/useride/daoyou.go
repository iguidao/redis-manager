package useride

import (
	"log"

	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func Gd_login(username, password string) bool {
	scrypt_password := Get_scrypt(password)
	//mysql_password := mysql.Get_user(phone)
	mysql_password, err := mysql.DB.FindUserPassword(username)
	// log.Println(scrypt_password)
	// log.Println(mysql_password.UserPassword)
	if err != nil {
		log.Println("数据库查询错误: ", err, username)
		return false
	}

	if scrypt_password != mysql_password.Password {
		// log.Println("登陆失败: ", username)
		return false
	} else {
		// log.Println("登陆成功: ", username)
		return true
	}
	//return "ok"

}

// func CacheUserinfo(token string, Phone int64) {
// 	opredis.RegisterAuthRedis(token)
// 	userinfo := mysql.DB.UserInfo(Phone)
// 	userconver := util.UserConverge(userinfo)
// 	jsonBody, _ := json.Marshal(userconver)
// 	opredis.WriteUserRedis(string(jsonBody), userinfo.Base.ID)
// }
// func RefreshUserinfo(token, oldtoken string) {
// 	opredis.RegisterAuthRedis(token)
// 	opredis.DelAuthRedis(oldtoken)
// }

// // 查mysql用户信息存redis
// func GetUserToRedis(id string) util.UserInfo {
// 	userinfo := mysql.DB.UserIdInfo(id)
// 	userconver := util.UserConverge(userinfo)
// 	jsonBody, _ := json.Marshal(userconver)
// 	opredis.WriteUserRedis(string(jsonBody), userinfo.Base.ID)
// 	return userconver
// }
