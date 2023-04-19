package mysql

import (
	"log"

	"github.com/iguidao/redis-manager/src/middleware/logger"
)

// add cfg
func (m *MySQL) AddCfg(name, key, value string) (int, bool) {
	addcfginfo := &Rconfig{
		Key:   key,
		Value: value,
		Name:  name,
	}
	result := m.Create(&addcfginfo)
	if result.Error != nil {
		return 0, false
	}
	return addcfginfo.ID, true
}
func (m *MySQL) DelCfg(key string) bool {
	var cfg *Rconfig
	if err := m.Model(cfg).Where("`key` = ?", key).Delete(&cfg).Error; err != nil {
		logger.Error(err)
		return false
	}
	return true
}

// update cfg
func (m *MySQL) UpdateCfg(key, value string) bool {
	var cfg Rconfig
	result := m.Model(&cfg).Where("`key` = ?", key).Update("value", value)
	log.Println("result.Error: ", result.Error)
	return result.Error == nil
}

// check cfg
func (m *MySQL) ExistCfg(key string) bool {
	var cfg *Rconfig
	if err := m.Model(cfg).Where("`key` = ?", key).First(&cfg).Error; err != nil {
		return false
	}
	return true
}

// get all cfg
func (m *MySQL) GetAllCfg() []Rconfig {
	var cfg []Rconfig
	m.Find(&cfg)
	return cfg
}

// get one cfg
func (m *MySQL) GetOneCfg(key string) Rconfig {
	var cfg Rconfig
	m.Where("`key` = ?", key).First(&cfg)
	return cfg
}

// get one cfg value
func (m *MySQL) GetOneCfgValue(key string) string {
	var cfg Rconfig
	m.Where("`key` = ?", key).First(&cfg)
	return cfg.Value
}
