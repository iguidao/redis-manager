package opredis

import "time"

type QueryResult struct {
	Ttl   time.Duration `json:"ttl"`
	Type  string        `json:"type"`
	Value interface{}   `json:"value"`
}
