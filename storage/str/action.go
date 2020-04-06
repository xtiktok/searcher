package str

import (
	"errors"
	"fmt"
	"searcher/common/consts"
	value2 "searcher/storage/value"
	"strconv"
	"time"
)

func DoSet(args []string) (string, error) {
	value := &value2.Value{
		Type: consts.StringType,
		Ttl:  0,
		Val:  args[1],
	}
	Set(args[0], value)
	return consts.Ok, nil
}

func DoGet(args []string) (string, error) {
	value, ok := Get(args[0])
	if !ok {
		return "", errors.New(consts.ErrorNilReturn)
	}
	if value == nil {
		return "", errors.New(consts.ErrorNilReturn)
	}

	resp, ok := value.(*value2.Value)
	if !ok {
		return "", errors.New(consts.ErrorUnCorrectType)
	}
	if resp.Ttl != 0 && resp.Ttl < time.Now().Unix() {
		// 并发操作未考虑
		Del(args[0])
		return "", errors.New(consts.ErrorNilReturn)
	} else {

	}
	return resp.Val.(string), nil
}

func DoExpire(args []string) error {
	value, ok := Get(args[0])
	if !ok {
		return errors.New(consts.ErrorNilReturn)
	}
	if value == nil {
		return errors.New(consts.ErrorNilReturn)
	}
	resp, ok := value.(*value2.Value)
	if !ok {
		return errors.New(consts.ErrorUnCorrectType)
	}

	ttl, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return errors.New(consts.ErrorWrongSyntax)
	}

	resp.Ttl = time.Now().Unix() + ttl
	return nil
}

func DoTtl(args []string) (int64, error) {
	value, ok := Get(args[0])
	if !ok {
		return 0, errors.New(consts.ErrorNilReturn)
	}
	if value == nil {
		return 0, errors.New(consts.ErrorNilReturn)
	}
	resp, ok := value.(*value2.Value)
	if !ok {
		return 0, errors.New(consts.ErrorWrongSyntax)
	}
	ttl := resp.Ttl - time.Now().Unix()
	if ttl <= 0 {
		ttl = -1
	}
	return ttl, nil
}

func DoSetEx(args []string) error {
	ttl, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return errors.New(consts.ErrorWrongSyntax)
	}

	value := &value2.Value{
		Type: consts.StringType,
		Ttl:  time.Now().Unix() + ttl,
		Val:  args[2],
	}
	Set(args[0], value)
	return nil
}

func DoSetNx(args []string) error {
	_, err := DoGet(args[0:1])
	if err == nil {
		return nil
	}

	if err.Error() != consts.ErrorNilReturn {
		return err
	}

	value := &value2.Value{
		Type: consts.StringType,
		Ttl:  0,
		Val:  args[1],
	}
	Set(args[0], value)
	return nil
}

func DoIncr(args []string) (int64, error) {
	val := int64(1)
	if len(args) == 2 {
		val, _ = strconv.ParseInt(args[1], 10, 64)
	}
	value, ok := Get(args[0])
	valueNew := &value2.Value{
		Type: consts.StringType,
		Ttl:  0,
		Val:  fmt.Sprintf("%d", val),
	}
	if !ok || value == nil {
		Set(args[0], valueNew)
		return val, nil
	}
	resp, ok := value.(*value2.Value)
	if !ok || (resp.Ttl != 0 && resp.Ttl < time.Now().Unix()) {
		Set(args[0], valueNew)
		return val, nil
	}
	varInt, err := strconv.ParseInt(resp.Val.(string), 10, 64)
	if err != nil {
		return 0, errors.New(consts.ErrorInvalidOperate)
	}
	resp.Val = fmt.Sprintf("%d", varInt+val)
	return varInt + val, nil
}

func DoDel(args []string) error {
	Del(args[0])
	return nil
}
