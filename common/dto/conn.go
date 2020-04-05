package dto

import "net"

type TsConn struct {
	Conn           net.Conn        // 链接
	Auth           bool            // 是否检验
	Count          map[uint16]uint // 计数器
	SendPackNum    int64           // 发送字节数
	ReceivePackNum int64           // 接受字节数
}
