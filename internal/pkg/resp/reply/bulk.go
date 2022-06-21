package reply

import (
	"github.com/rtmzk/gedis/internal/pkg/resp/reply/constandsReply"
	"strconv"
)

type BulkReply struct {
	Arg []byte
}

func (b *BulkReply) ToBytes() []byte {
	if len(b.Arg) == 0 {
		return constandsReply.NewNullBulkReply().ToBytes()
	}
	return []byte("$" + strconv.Itoa(len(b.Arg)) + CRLF + string(b.Arg) + CRLF)
}

func NewBulkReply(arg []byte) *BulkReply {
	return &BulkReply{
		Arg: arg,
	}
}
