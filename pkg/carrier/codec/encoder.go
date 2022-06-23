package codec

type Encoder interface {
	Encode(e any) error
}
