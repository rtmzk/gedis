package constandsReply

type EmptyReply struct {
}

var emptyBytes = []byte("$0\r\n")

func (e *EmptyReply) ToBytes() []byte {
	return emptyBytes
}

var emptyReply = new(EmptyReply)

func NewEmptyReply() *EmptyReply {
	return emptyReply
}
