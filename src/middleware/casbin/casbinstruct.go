package casbin

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	n             = 12
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var src = rand.NewSource(time.Now().UnixNano())

type CasbinRule struct {
	ID        string    `gorm:"type:uuid;primaryKey;type:char(12);column:id"`
	Ptype     string    `gorm:"size:32;uniqueIndex:unique_index"`
	V0        string    `gorm:"size:64;uniqueIndex:unique_index"`
	V1        string    `gorm:"size:512;uniqueIndex:unique_index"`
	V2        string    `gorm:"size:16;uniqueIndex:unique_index"`
	V3        string    `gorm:"size:32;uniqueIndex:unique_index"`
	V4        string    `gorm:"size:32;uniqueIndex:unique_index"`
	V5        string    `gorm:"size:32;uniqueIndex:unique_index"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (c *CasbinRule) BeforeCreate(tx *gorm.DB) error {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	c.ID = string(b)
	return nil
}
