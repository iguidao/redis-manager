package mysql

import "github.com/iguidao/redis-manager/src/middleware/logger"

func (m *MySQL) GetAllHistory() []OpHistory {
	var ophistory []OpHistory
	m.Find(&ophistory)
	return ophistory
}

// add cluster
func (m *MySQL) AddHistory(userid int, opinfo, opparams string) (int, bool) {
	addcluster := &OpHistory{
		UserId:   userid,
		OpInfo:   opinfo,
		OpParams: opparams,
	}
	result := m.Create(&addcluster)
	if result.Error != nil {
		logger.Error("Mysql add history error:", result.Error)
		return 0, false
	}
	return addcluster.ID, true
	// return gdarticle.ID.String(), true
}
