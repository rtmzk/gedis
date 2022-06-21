package reply

import (
	"bytes"
	"github.com/rtmzk/gedis/internal/pkg/resp/reply/constandsReply"
	"strconv"
)

type MultiBulkReply struct {
	Arg [][]byte
}

func (m *MultiBulkReply) ToBytes() []byte {
	argLen := len(m.Arg)
	var buf = bytes.Buffer{}
	buf.WriteString("*" + strconv.Itoa(argLen) + CRLF)
	for _, elem := range m.Arg {
		if elem == nil {
			buf.WriteString(string(constandsReply.NewNullBulkReply().ToBytes()))
			continue
		}
		buf.WriteString("$" + strconv.Itoa(len(elem)) + CRLF + string(elem) + CRLF)
	}
	return buf.Bytes()
}

func NewMultiBulkReply(arg [][]byte) *MultiBulkReply {
	return &MultiBulkReply{
		arg,
	}
}
