package model

import (
	"errors"
	"searcher/common/consts"
)

type TsResponse []byte

func (ts *TsResponse) Init() {
	version01 := uint8(uint16(consts.ServerVersion) >> 8)
	version02 := uint8(uint16(consts.ServerVersion) & 0x00ff)
	*ts = append(*ts, []byte{0xaa, 0x55, version01, version02}...)
}

//  标记: 0xaa 0x55,   server版本 0x00 0x01 , type  0x01 , length 0x00 0x00 0x00 0x32
func (ts *TsResponse) AddBody(body *ResponseBody) {
	*ts = append(*ts, body.VarType)
	var length [4]byte
	lengthInt := len(body.Body)
	length[0] = uint8((lengthInt & 0xff000000) >> 24)
	length[1] = uint8((lengthInt & 0x00ff0000) >> 16)
	length[2] = uint8((lengthInt & 0x0000ff00) >> 8)
	length[3] = uint8(lengthInt & 0x000000ff)
	*ts = append(*ts, length[:]...)
	if len(body.Body) > 0 {
		*ts = append(*ts, []byte(body.Body)...)
	}
}

type ResponseHeader struct {
	Version    uint16
	Type       uint8
	BodyLength uint32
}

type ResponseBody struct {
	VarType uint8
	Body    string
}

func ParseRespHeader(data []byte) (*ResponseHeader, error) {
	if len(data) != 9 {
		return nil, errors.New("parse error")
	}

	if data[0] != 0xaa || data[1] != 0x55 {
		return nil, errors.New("parse error")
	}
	header := &ResponseHeader{}
	header.Version = (uint16(data[2]) << 8) + uint16(data[3])
	header.Type = data[4]
	if data[4] == consts.NilVar {
		return header, nil
	}

	header.BodyLength = (uint32(data[5]) << 24) + (uint32(data[6]) << 16) + (uint32(data[7]) << 8) + uint32(data[8])
	return header, nil
}

func BuildResponse(body *ResponseBody) *TsResponse {
	resp := &TsResponse{}
	resp.Init()
	resp.AddBody(body)
	return resp
}

func ParseRespBody(t uint8, data []byte) (*ResponseBody, error) {
	body := &ResponseBody{}
	body.VarType = t
	body.Body = string(data)
	return body, nil
}
