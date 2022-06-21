package constandsErrReply

type SyntaxErrReply struct{}

var syntaxErrBytes = []byte("-Err syntax error\r\n")
var syntaxErrReply = &SyntaxErrReply{}

func (s *SyntaxErrReply) Error() string {
	return "Err syntax error"
}

func (s *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

func NewSyntaxErrReply() *SyntaxErrReply {
	return syntaxErrReply
}
