package model

import "time"

type ClusterNodeTables struct {
	CreateTime time.Time
	CluserId   int
	NodeId     string
	Address    string
	Flags      string
	LinkState  string
	RunStatus  bool
	SlotRange  string
	SlotNumber int
	Children   []*ClusterNodeTables
}
