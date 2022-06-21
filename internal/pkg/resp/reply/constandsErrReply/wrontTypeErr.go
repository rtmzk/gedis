package constandsErrReply

type WrongTypeErrReply struct {
}

var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")
var wrongTypeErrReply = &WrongTypeErrReply{}

func (w *WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key holding the wrong kind of value"
}

func (w *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

func NewWrongTypeErrReply() *WrongTypeErrReply {
	return wrongTypeErrReply
}
