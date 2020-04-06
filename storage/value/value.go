package value

import "sync"

type Value struct {
	sync.RWMutex
	Type uint8
	Ttl  int64
	Val  interface{}
}
