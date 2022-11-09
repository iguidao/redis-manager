package useride

import (
	"bytes"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"
)

// 密码加密
func Get_scrypt(password string) string {
	pw_new := string([]byte(password)[:3])
	user_src := []byte(pw_new)
	other_src := []byte("redis")
	all_src := [][]byte{user_src, other_src}
	salt := bytes.Join(all_src, []byte{})
	dk, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	user_password := base64.StdEncoding.EncodeToString(dk)
	return user_password
}
