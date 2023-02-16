package mysql

// add cfg
func (m *MySQL) AddCfg(name, value, note string) (int, bool) {
	addcfginfo := &Rconfig{
		Name:  name,
		Value: value,
		Note:  note,
	}
	result := m.Create(&addcfginfo)
	if result.Error != nil {
		return 0, false
	}
	return addcfginfo.ID, true
}

// update cfg
func (m *MySQL) UpdateCfg(name, value string) bool {
	var cfg Rconfig
	result := m.Model(&cfg).Where("name = ?", name).Update("value", value)
	return result.Error == nil
}

// check cfg
func (m *MySQL) ExistCfg(name string) bool {
	var cfg *Rconfig
	if err := m.Model(cfg).Where("name = ?", name).First(&cfg).Error; err != nil {
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
func (m *MySQL) GetOneCfg(name string) Rconfig {
	var cfg Rconfig
	m.Where("name = ?", name).First(&cfg)
	return cfg
}

// get one cfg value
func (m *MySQL) GetOneCfgValue(name string) string {
	var cfg Rconfig
	m.Where("name = ?", name).First(&cfg)
	return cfg.Value
}
