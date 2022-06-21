package constandsErrReply

type ArgNumErrReply struct {
	Cmd string
}

func (a *ArgNumErrReply) Error() string {
	return "Err wrojng number of arguments for '" + a.Cmd + "' command\r\n"
}

func (a *ArgNumErrReply) ToBytes() []byte {
	return []byte("-Err wrojng number of arguments for '" + a.Cmd + "' command\r\n")
}

func NewArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}
