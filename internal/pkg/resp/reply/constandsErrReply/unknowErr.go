package constandsErrReply

type UnknowErrReply struct {
}

var unknowErrReplyBytes = []byte("-Err unknow\r\n")
var unknowErrReply = &UnknowErrReply{}

func (u *UnknowErrReply) Error() string {
	return "-Err unknow"
}

func (u *UnknowErrReply) ToBytes() []byte {
	return unknowErrReplyBytes
}

func NewUnknowReply() *UnknowErrReply {
	return unknowErrReply
}
