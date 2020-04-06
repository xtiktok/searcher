package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"searcher/common/config"
	"searcher/common/consts"
	"searcher/common/utils"
	"searcher/model"
	"strings"
)

var res map[string]string

func init() {
	res = make(map[string]string)
}

var name = "ts_client"

func main() {

	conf := config.ClientConfig{}
	conf.Addr = flag.String("a", "127.0.0.1", "server ip address")
	conf.Port = flag.Int("p", 6379, "server  remote port")
	conf.Password = flag.String("P", "", "auth password")

	flag.Parse()

	f := bufio.NewReader(os.Stdin)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *conf.Addr, *conf.Port))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "\033[31m%s\n", err.Error())
		return
	}
	defer conn.Close()
	_, _ = fmt.Fprintf(os.Stdout, "\033[32m%s>\033[33m", name)
	for {
		res, _, _ := f.ReadLine()
		input := string(res)
		resp, err := InputHandle(&conn, input)

		if err != nil && err.Error() == consts.ErrorNilReturn {
			_, _ = fmt.Fprintf(os.Stdout, "\033[37m%s\n", err.Error())
			continue
		}

		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "\033[31m%s\n", err.Error())
			continue
		}
		OutPutPrint(resp)
	}
}

func OutPutPrint(v interface{}) {
	res, ok := v.(string)
	if ok {
		if res == "" {
			_, _ = fmt.Fprintf(os.Stdout, "\033[32m%s>\033[33m", name)
			return
		}
		_, _ = fmt.Fprintf(os.Stdout, "\033[33m%s\n", res)
	}
	resMap, ok := v.(map[string]interface{})
	if ok {
		for key, value := range resMap {
			_, _ = fmt.Fprintf(os.Stdout, "\033[33m%s => %v\n", key, value)
		}
	}
	_, _ = fmt.Fprintf(os.Stdout, "\033[32m%s>\033[33m", name)
}

func InputHandle(conn *net.Conn, argsStr string) (interface{}, error) {
	var heads [9]byte
	if conn == nil {
		return "", errors.New("not connected")
	}
	realConn := *conn
	argsStr = strings.TrimSpace(argsStr)
	if argsStr == "" {
		return "", nil
	}
	spaceRe, _ := regexp.Compile(`\s+`)
	params := spaceRe.Split(argsStr, -1)
	if len(params) == 0 {
		return "", nil
	}
	for i := range params {
		params[i] = strings.Trim(params[i], "\"")
	}

	command := params[0]
	args := params[1:]

	t, ok := consts.Command[command]
	if !ok {
		return "", errors.New("command no found")
	}
	rule, ok := consts.CommandRule[command]
	if ok {
		if err := utils.ArgsCheck(args, rule); err != nil {
			return "", err
		}
	}
	req := model.BuildRequest(t, args)
	_, err := realConn.Write(*req)
	if err != nil {
		return "", err
	}

	_, err = realConn.Read(heads[:])
	if err != nil {
		return "", err
	}
	header, err := model.ParseRespHeader(heads[:])
	if err != nil {
		return "", err
	}
	if header.Type == consts.NilVar {
		return "", errors.New(consts.ErrorNilReturn)
	}
	data := make([]byte, header.BodyLength)
	_, err = realConn.Read(data)
	body, err := model.ParseRespBody(header.Type, data)
	if err != nil {
		return "", err
	}
	p := strings.TrimSpace(body.Body)
	if body.VarType == consts.StringVar {
		return p, nil
	}
	if body.VarType == consts.Map {
		resp := make(map[string]interface{})
		err = json.Unmarshal([]byte(p), &resp)
		if err != nil {
			return p, nil
		}
		return resp, nil
	}
	return p, nil
}
