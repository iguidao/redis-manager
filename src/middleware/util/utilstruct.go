package util

import "time"

// user info tbale
type UserJWTInfo struct {
	Uid      int
	UserName string
	UserType string
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
