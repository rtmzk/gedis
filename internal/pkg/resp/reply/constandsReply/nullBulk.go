package constandsReply

type NullBulkReply struct {
}

var nullBulkBytes = []byte("$-1\r\n")

func (n *NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

var nullBulkReply = new(NullBulkReply)

func NewNullBulkReply() *NullBulkReply {
	return nullBulkReply
}
