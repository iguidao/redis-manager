package mysql

import "github.com/iguidao/redis-manager/src/middleware/logger"

// add codis

func (m *MySQL) AddCodis(curl, cname string) (int, bool) {
	addcodisinfo := &CodisInfo{
		Curl:  curl,
		Cname: cname,
	}
	result := m.Create(&addcodisinfo)
	if result.Error != nil {
		logger.Error("Mysql add Codis error:", result.Error)
		return 0, false
	}
	return addcodisinfo.ID, true
	// return gdarticle.ID.String(), true
}

// get codis
func (m *MySQL) GetAllCodis() []CodisInfo {
	var clusters []CodisInfo
	m.Find(&clusters)
	return clusters
}

func (m *MySQL) GetCodisNumber() int64 {
	var clusters []CodisInfo
	var count int64
	m.Find(&clusters).Count(&count)
	return count
}
