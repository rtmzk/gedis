package constandsErrReply

type ProtocolErrReply struct {
	Msg string
}

func (p *ProtocolErrReply) Error() string {
	return "ERR Protocol error: '" + p.Msg
}
func (p *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error: '" + p.Msg + "'\r\n")
}
