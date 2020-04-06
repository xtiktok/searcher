package utils

import (
	"errors"
	"searcher/common/consts"
	"searcher/common/dto"
	"strconv"
)

func ArgsCheck(args []string, rule *dto.Rule) error {
	if rule.Argc != -1 && len(args) != rule.Argc {
		return errors.New("wrong syntax")
	}

	if rule.Max != 0 && len(args) > rule.Max {
		return errors.New("wrong syntax")
	}

	if rule.Min != 0 && len(args) < rule.Min {
		return errors.New("wrong syntax")
	}

	if rule.OddEvenCheck == 1 && len(args)%2 == 0 {
		return errors.New("wrong syntax")
	}

	if rule.OddEvenCheck == 2 && len(args)%2 == 1 {
		return errors.New("wrong syntax")
	}

	for i, t := range rule.TypeCheck {
		if len(args) <= i {
			return nil
		}
		switch t {
		case consts.RuleTypeInt:
			_, err := strconv.Atoi(args[i])
			if err != nil {
				return errors.New("wrong syntax")
			}
			break
		case consts.RuleTypeUInt:
			_, err := strconv.ParseUint(args[i], 10, 32)
			if err != nil {
				return errors.New("wrong syntax")
			}
			break
		case consts.RuleTypeInt64:
			_, err := strconv.ParseInt(args[i], 10, 64)
			if err != nil {
				return errors.New("wrong syntax")
			}
			break
		case consts.RuleTypeUInt64:
			_, err := strconv.ParseUint(args[i], 10, 64)
			if err != nil {
				return errors.New("wrong syntax")
			}
			break
		case consts.RuleTypeFloat64:
			_, err := strconv.ParseFloat(args[i], 64)
			if err != nil {
				return errors.New("wrong syntax")
			}
			break
		case consts.RuleTypeFloat:
			_, err := strconv.ParseFloat(args[i], 32)
			if err != nil {
				return errors.New("wrong syntax")
			}
			break
		}
	}
	return nil
}
