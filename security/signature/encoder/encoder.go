package encoder

type Encoder interface {
	Encode() string
	Init([]byte,string)
}
