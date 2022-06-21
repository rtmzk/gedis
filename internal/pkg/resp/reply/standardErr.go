package reply

type StandardErrReply struct {
	Status string
}

func (s *StandardErrReply) Error() string {
	return s.Status
}

func (s *StandardErrReply) ToBytes() []byte {
	return []byte(CommonErrorPrefix + s.Status + CRLF)
}

func NewStandardErrReply(status string) *StandardErrReply {
	return &StandardErrReply{
		Status: status,
	}
}
