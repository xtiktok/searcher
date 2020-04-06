package str

import (
	"crypto/md5"
	"hash/crc32"
	"sync"
)

type StringStore struct {
	sync.Map
}

var store [16384]*StringStore

func init() {
	for i := range store {
		store[i] = &StringStore{}
	}
}

func Set(key string, value interface{}) {
	i := CurrentStorage(key)
	store[i].Store(key, value)
}

func Get(key string) (interface{}, bool) {
	i := CurrentStorage(key)
	return store[i].Load(key)
}

func Del(key string) {
	i := CurrentStorage(key)
	store[i].Delete(key)
}

func CurrentStorage(key string) int {
	var item string
	if len(key) < 5 {
		item = key + string([]byte{0, 0, 0, 0, 0}[0:5-len(key)])
	} else {
		item = key[0:5]
	}
	c := md5.New()
	c.Write([]byte(item))
	res := crc32.ChecksumIEEE([]byte(item))
	return int(res % 16384)
}
