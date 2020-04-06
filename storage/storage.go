package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"searcher/common/consts"
	"searcher/common/utils"
	"searcher/model"
	"searcher/storage/hash"
	"searcher/storage/str"
)

func DoAction(action uint16, args []string) (*model.ResponseBody, error) {
	body := &model.ResponseBody{
		VarType: consts.NilVar,
	}

	opname, ok := consts.CommandMap[action]
	if !ok {
		return nil, errors.New(consts.ErrorInvalidOperate)
	}
	rule := consts.CommandRule[opname]
	if rule == nil {
		return nil, errors.New(consts.ErrorInvalidOperate)
	}

	err := utils.ArgsCheck(args, rule)
	if err != nil {
		return nil, err
	}

	switch action {
	case consts.StringSet:
		resp, err := str.DoSet(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = resp
		return body, nil
	case consts.StringGet:
		res, err := str.DoGet(args)
		if err != nil {
			return body, nil
		}
		body.VarType = consts.StringVar
		body.Body = res
		return body, nil
	case consts.StringSetEx:
		err := str.DoSetEx(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = consts.Ok
		return body, nil
	case consts.StringSetNx:
		err := str.DoSetNx(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = consts.Ok
		return body, nil

	case consts.StringExpire:
		err := str.DoExpire(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = consts.Ok
		return body, nil
	case consts.StringTtl:
		ttl, err := str.DoTtl(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = fmt.Sprintf("%d", ttl)
		return body, nil
	case consts.StringDel:
		err := str.DoDel(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = consts.Ok
		return body, nil
	case consts.StringIncr:
		val, err := str.DoIncr(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = fmt.Sprintf("%d", val)
		return body, nil
	case consts.HashHSet:
		err := hash.DoHSet(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = consts.Ok
		return body, nil
	case consts.HashHDel:
		err := hash.DoHDel(args)
		if err != nil {
			return body, err
		}
		body.VarType = consts.StringVar
		body.Body = consts.Ok
		return body, nil
	case consts.HashHGet:
		res, err := hash.DoHGet(args)
		if err != nil {
			return body, err
		}
		if res == "" {
			return body, nil
		}
		body.VarType = consts.StringVar
		body.Body = res
		return body, nil
	case consts.HashHGetAll:
		res, err := hash.DoHGetAll(args)
		if err != nil {
			return body, err
		}
		if res == nil {
			return body, nil
		}
		b, _ := json.Marshal(res)
		body.VarType = consts.Map
		body.Body = string(b)
		return body, nil
	}

	body.VarType = consts.StringVar
	body.Body = consts.Ok
	return body, nil
}
