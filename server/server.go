package main

import (
	"fmt"
	"net"
	"searcher/common/dto"
	"searcher/model"
	"searcher/storage"
)

type Pack struct {
	Tag  string `json:"tag"`  //消息识别标记
	Type int16  `json:"type"` //消息类型
	Msg  string `json:"msg"`  //其他消息
}

var ResMap map[string]string

func init() {
	ResMap = make(map[string]string)
}
func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept error %s", err.Error())
			continue
		}
		tsConn := &dto.TsConn{}
		tsConn.Conn = conn
		go handle(tsConn)
	}
}

// 识别标记+保留4bit  0xaa 0x55  协议版本+请求来源标记各4位 0xb1   操作类型0xc1 0x01   数据长度（最大4g） 0x00 0x00 0x00 0x00  数据 0x0a 0x0b ....
func handle(tsConn *dto.TsConn) {
	conn := tsConn.Conn
	defer conn.Close()
	var heads [9]byte
	for {
		_, err := conn.Read(heads[:])
		if err != nil {
			return
		}
		header, err := model.ParseHeader(heads[:])
		if err != nil {
			_, err = conn.Write([]byte("unknown command"))
			if err != nil {
				return
			}
		}
		var args []string
		if header.BodyLength != 0 {
			data := make([]byte, header.BodyLength)
			_, err = conn.Read(data)
			if err != nil {
				return
			}
			args = model.ParseBody(data)
		}

		fmt.Println(header.Type, args)
		body, _ := storage.DoAction(header.Type, args)
		resp := model.BuildResponse(body)
		_, err = conn.Write(*resp)
		if err != nil {
			return
		}
	}
}
