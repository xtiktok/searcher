package model

import (
	"errors"
	"searcher/common/consts"
	"strings"
)

type TsRequest []byte

func (ts *TsRequest) Init() {
	prefix := uint8(consts.ClientPrefix<<4 + consts.ClientVersion)
	*ts = append(*ts, []byte{0xaa, 0x55, prefix}...)
}

func (ts *TsRequest) AddHeader(t uint16) {
	first := uint8(t >> 8)
	second := uint8(t & 0x00ff)
	*ts = append(*ts, []byte{first, second}...)
}

func (ts *TsRequest) AddArgs(args []string) {
	var length [4]byte
	var buffer []byte
	for _, arg := range args {
		buffer = append(buffer, byte(0x00))
		buffer = append(buffer, []byte(arg)...)
	}
	lengthInt := len(buffer)
	length[0] = uint8((lengthInt & 0xff000000) >> 24)
	length[1] = uint8((lengthInt & 0x00ff0000) >> 16)
	length[2] = uint8((lengthInt & 0x0000ff00) >> 8)
	length[3] = uint8(lengthInt & 0x000000ff)
	*ts = append(*ts, length[:]...)
	if len(buffer) > 0 {
		*ts = append(*ts, buffer...)
	}

}

func BuildRequest(t uint16, args []string) *TsRequest {
	req := &TsRequest{}
	req.Init()
	req.AddHeader(t)
	req.AddArgs(args)
	return req
}

type RequestHeader struct {
	Source     uint8
	Version    uint8
	Type       uint16
	BodyLength uint32
}

func ParseHeader(data []byte) (*RequestHeader, error) {
	if len(data) != 9 {
		return nil, errors.New("parse error")
	}

	if data[0] != 0xaa || data[1] != 0x55 {
		return nil, errors.New("parse error")
	}

	header := &RequestHeader{}
	header.Source = data[2] >> 4
	header.Version = data[2] & 0x0f
	header.Type = (uint16(data[3]) << 8) + uint16(data[4])
	header.BodyLength = (uint32(data[5]) << 24) + (uint32(data[6]) << 16) + (uint32(data[7]) << 8) + uint32(data[8])
	return header, nil
}

func ParseBody(data []byte) []string {
	str := string(data)
	str = strings.Trim(str, string([]byte{0x00}))
	if str == "" {
		return nil
	}
	return strings.Split(str, string([]byte{0x00}))
}
