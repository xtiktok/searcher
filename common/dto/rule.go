package dto

type Rule struct {
	Argc         int          `about:"参数个数: -1 不限制 0 不需要额外参数  1 一个参数...."`
	OddEvenCheck int          `about:"参数奇偶个数校验: 0 不校验  1 只能有奇数个参数  2 只能由偶数个参数 "`
	Max          int          `about:"最大参数个数  0 不限制，  n 限制为n"`
	Min          int          `about:"最少参数限制 0 不限制， n限制为n"`
	TypeCheck    map[int]uint `about:"特殊参数类型校验，第几个应该是什么类型"`
}
