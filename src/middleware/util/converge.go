package util

func ReturnDefaultModel(val string, list map[string]string) string {
	for i, v := range list {
		if val == i {
			return v
		}
	}
	return "自定义配置"
}

// func UserConverge(muserinfo mysql.RdUser) UserInfo {
// 	var userinfo UserInfo
// 	userinfo.Uid = muserinfo.Base.ID
// 	userinfo.UserName = muserinfo.NickName
// 	userinfo.Identity = muserinfo.Identity
// 	// userinfo.Phone = muserinfo.UserPhone
// 	userinfo.AvatarUrl = muserinfo.AvatarUrl
// 	userinfo.CreatedAt = muserinfo.Base.CreatedAt
// 	return userinfo
// }

// func claimsClaimsConverge(mclaims *Claims) UserJWTInfo {
// 	var userinfo UserJWTInfo
// 	userinfo.ID = mclaims.UserId
// 	userinfo.UserName = mclaims.UserName
// 	// userinfo.UserPhone = mclaims.UserPhone
// 	// userinfo.CreatedAt = mclaims.CreaTime
// 	return userinfo
// }
