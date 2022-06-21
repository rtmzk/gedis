package constandsReply

type EmptyMultiBulkReply struct {
}

var emptyMultiBulkBytes = []byte("*0\r\n")

func (emb *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

var emptyMultiBulkReply = new(EmptyMultiBulkReply)

func NewEmptyMultiBulkReply() *EmptyMultiBulkReply {
	return emptyMultiBulkReply
}
