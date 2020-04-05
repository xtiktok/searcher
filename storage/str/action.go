package str

import (
	"errors"
	"searcher/common/consts"
)

var MemMap map[string]string

func init() {
	MemMap = make(map[string]string)
}

func DoSet(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New(consts.ErrorWrongSyntax)
	}
	MemMap[args[0]] = args[1]
	return consts.Ok, nil
}

func DoGet(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New(consts.ErrorWrongSyntax)
	}
	res, ok := MemMap[args[0]]
	if !ok {
		return "", errors.New(consts.ErrorNilReturn)
	}
	return res, nil
}
