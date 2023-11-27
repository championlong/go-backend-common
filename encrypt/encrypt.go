package encrypt

type Method interface {
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, signature []byte) error
}
