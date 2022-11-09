package util

import "time"

// user info tbale
type UserJWTInfo struct {
	Uid      string
	UserName string
	// UserPhone int64
	// CreatedAt time.Time
}

// user info tbale
type UserInfo struct {
	Uid      int
	UserName string
	Identity string
	// UserPhone int64
	AvatarUrl string
	CreatedAt time.Time
}
