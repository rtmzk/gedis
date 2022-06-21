package constandsReply

type NoReply struct {
}

var noReplyBytes = []byte("")

func (n NoReply) ToBytes() []byte {
	return noReplyBytes
}

var noReply = new(NoReply)

func NewNoReply() *NoReply {
	return noReply
}
