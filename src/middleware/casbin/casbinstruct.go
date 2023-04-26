package casbin

type CasbinRule struct {
	ID    int    `gorm:"primary_key"`
	Ptype string `gorm:"size:32;uniqueIndex:unique_index"`
	V0    string `gorm:"size:64;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:16;uniqueIndex:unique_index"`
	V3    string `gorm:"size:32;uniqueIndex:unique_index"`
	V4    string `gorm:"size:32;uniqueIndex:unique_index"`
	V5    string `gorm:"size:32;uniqueIndex:unique_index"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
