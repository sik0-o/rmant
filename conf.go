package rmant

import "time"

type Conf struct {
	Prefix    string
	MarkTTL   time.Duration
	MarkValue interface{}
}
