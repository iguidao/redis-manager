package opredis

type QueryResult struct {
	Ttl   string      `json:"ttl"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
	Len   int         `json:"len"`
	// Debug string        `json:"debug"`
}
