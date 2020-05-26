package key

import (
	"crypto/ed25519"
)

type secKeyEd25519 struct {
	key ed25519.PrivateKey
}

type pubKeyEd25519 struct {
	key ed25519.PublicKey
}

func (sec secKeyEd25519) PublicKey() PublicKey {
	return sec.key.Public().(PublicKey)
}

func (sec secKeyEd25519) Sign(msg []byte) []byte {
	return ed25519.Sign(sec.key, msg)
}

func (sec secKeyEd25519) ToBytes() []byte {
	b := make([]byte, ed25519.PrivateKeySize)
	copy(b, sec.key)
	return b
}

func (pub pubKeyEd25519) VerifySignature(msg []byte, sig []byte) bool {
	return ed25519.Verify(pub.key, msg, sig)
}

func (pub pubKeyEd25519) ToBytes() []byte {
	b := make([]byte, ed25519.PublicKeySize)
	copy(b, pub.key)
	return b
}

func NewPublicKey(b []byte) PublicKey {
	key := make(ed25519.PublicKey, ed25519.PublicKeySize)
	copy(key, b)
	return pubKeyEd25519{key: key}
}

func NewSecretKey(b []byte) SecretKey {
	key := make(ed25519.PrivateKey, ed25519.PrivateKeySize)
	copy(b, key)
	return secKeyEd25519{key: key}
}
