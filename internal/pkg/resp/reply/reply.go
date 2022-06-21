package reply

import "github.com/rtmzk/gedis/pkg/resp"

const (
	CommonErrorPrefix  = "-"
	CommonStringPrefix = "+"
	CommonIntPrefix    = ":"
	CRLF               = "\r\n"
)

func IsErrReply(reply resp.Reply) bool {
	return reply.ToBytes()[0] == '-'
}
