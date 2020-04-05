package storage

import (
	"searcher/common/consts"
	"searcher/model"
	"searcher/storage/str"
)

func DoAction(action uint16, args []string) (*model.ResponseBody, error) {
	body := &model.ResponseBody{
		VarType: consts.NilVar,
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
	}
	body.VarType = consts.StringVar
	body.Body = "ok"
	return body, nil
}
