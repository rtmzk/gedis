package constandsReply

type OkReply struct {
}

var okBytes = []byte("+OK\r\n")

func (o OkReply) ToBytes() []byte {
	return okBytes
}

var okReply = new(OkReply)

func NewOkReply() *OkReply {
	return okReply
}
