package constandsReply

type PongReply struct {
}

var pongbytes = []byte("+PONG\r\n")

func (p *PongReply) ToBytes() []byte {
	return pongbytes
}

func NewPongReply() *PongReply {
	return &PongReply{}
}
