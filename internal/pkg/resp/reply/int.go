package reply

import "strconv"

type IntReply struct {
	Code int64
}

func (i *IntReply) ToBytes() []byte {
	return []byte(CommonIntPrefix + strconv.FormatInt(i.Code, 10) + CRLF)
}

func NewIntReply(code int64) *IntReply {
	return &IntReply{
		Code: code,
	}
}
