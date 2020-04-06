package hash

import (
	"errors"
	"searcher/common/consts"
	"searcher/storage/str"
	"searcher/storage/value"
)

func DoHSet(args []string) error {
	item, ok := str.Get(args[0])
	if !ok {
		val := &value.Value{
			Type: consts.HashType,
			Ttl:  0,
			Val:  &str.StringStore{},
		}
		str.Set(args[0], val)
		item = val
	}
	val, ok := item.(*value.Value)
	if !ok || val.Type != consts.HashType {
		return errors.New(consts.ErrorInvalidOperate)
	}
	store, ok := val.Val.(*str.StringStore)
	if !ok {
		return errors.New(consts.ErrorInternalError)
	}

	for i := 0; i < len(args)/2; i++ {
		store.Store(args[2*i+1], args[2*i+2])
	}

	return nil
}

func DoHDel(args []string) error {
	item, ok := str.Get(args[0])
	if !ok {
		return nil
	}
	val, ok := item.(*value.Value)
	if !ok {
		return nil
	}

	if val.Type != consts.HashType {
		return errors.New(consts.ErrorInvalidOperate)
	}

	store, ok := val.Val.(*str.StringStore)
	if !ok {
		return errors.New(consts.ErrorInternalError)
	}

	for i := 1; i < len(args); i++ {
		store.Delete(args[i])
	}
	return nil
}

func DoHGet(args []string) (string, error) {
	item, ok := str.Get(args[0])
	if !ok {
		return "", nil
	}
	val, ok := item.(*value.Value)
	if !ok {
		return "", nil
	}

	if val.Type != consts.HashType {
		return "", errors.New(consts.ErrorInvalidOperate)
	}

	store, ok := val.Val.(*str.StringStore)
	if !ok {
		return "", errors.New(consts.ErrorInternalError)
	}
	res, ok := store.Load(args[1])
	if !ok {
		return "", nil
	}
	return res.(string), nil
}

func DoHGetAll(args []string) (map[string]interface{}, error) {
	resp := make(map[string]interface{})
	item, ok := str.Get(args[0])
	if !ok {
		return nil, nil
	}
	val, ok := item.(*value.Value)
	if !ok {
		return nil, nil
	}

	if val.Type != consts.HashType {
		return nil, errors.New(consts.ErrorInvalidOperate)
	}

	store, ok := val.Val.(*str.StringStore)
	if !ok {
		return nil, errors.New(consts.ErrorInternalError)
	}
	store.Range(func(key, value interface{}) bool {
		resp[key.(string)] = value
		return true
	})
	return resp, nil
}
