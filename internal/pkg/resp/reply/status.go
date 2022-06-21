package reply

type StatusReply struct {
	Status string
}

func (s *StatusReply) ToBytes() []byte {
	return []byte(CommonStringPrefix + s.Status + CRLF)
}

func NewStatusReply(status string) *StatusReply {
	return &StatusReply{
		Status: status,
	}
}
