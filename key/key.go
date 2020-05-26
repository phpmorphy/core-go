package key

type Key interface {
	ToBytes() []byte
}

type PublicKey interface {
	Key
	VerifySignature(msg []byte, sig []byte) bool
}

type SecretKey interface {
	Key
	PublicKey() PublicKey
	Sign(msg []byte) []byte
}
