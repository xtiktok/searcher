package consts

import "searcher/common/dto"

const (
	UnSupport = 0
	ShutDown  = 1
	Restart   = 2
	Auth      = 5001

	StringSet    = 6001
	StringGet    = 6002
	StringSetEx  = 6003
	StringSetNx  = 6004
	StringDel    = 6005
	StringExpire = 6006
	StringTtl    = 6007

	HashHSet    = 6101
	HashHGet    = 6102
	HashHGetAll = 6103
	HashHDel    = 6104

	NotifyWatch = 7001
	KeysAll     = 10000
)

var Command = map[string]uint16{
	"shutdown": ShutDown,
	"restart":  Restart,
	"auth":     Auth,
	"set":      StringSet,
	"get":      StringGet,
	"setex":    StringSetEx,
	"setnx":    StringSetNx,
	"del":      StringDel,
	"expire":   StringExpire,
	"ttl":      StringTtl,

	"hset":    HashHSet,
	"hget":    HashHGet,
	"hgetall": HashHGetAll,
	"hdel":    HashHDel,
	"watch":   NotifyWatch,
	"keys":    KeysAll,
}

var CommandRule = map[string]*dto.Rule{
	"auth": {
		Argc:         1,
		OddEvenCheck: 0,
	},
	"set": {
		Argc:         2,
		OddEvenCheck: 0,
	},
	"get": {
		Argc:         1,
		OddEvenCheck: 0,
	},
	"setex": {
		Argc:         3,
		OddEvenCheck: 0,
		TypeCheck: map[int]uint{
			1: RuleTypeUInt,
		},
	},
	"setnx": {
		Argc:         2,
		OddEvenCheck: 0,
	},
	"del": {
		Argc:         1,
		OddEvenCheck: 0,
	},
	"expire": {
		Argc:         1,
		OddEvenCheck: 0,
		TypeCheck: map[int]uint{
			1: RuleTypeUInt,
		},
	},
	"ttl": {
		Argc:         1,
		OddEvenCheck: 0,
		TypeCheck:    nil,
	},

	"hset": {
		Argc:         -1,
		OddEvenCheck: 1,
		Min:          3,
	},
	"hget": {
		Argc:         2,
		OddEvenCheck: 0,
	},
	"hgetall": {
		Argc:         -1,
		OddEvenCheck: 0,
		Min:          2,
	},
	"hdel": {
		Argc:         -1,
		OddEvenCheck: 0,
		Min:          2,
	},
	"watch": {
		Argc:         -1,
		OddEvenCheck: 0,
		Max:          3,
		Min:          2,
	},
	"keys": {
		Argc:         1,
		OddEvenCheck: 0,
	},
}
