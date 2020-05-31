package key

import (
	"crypto/ed25519"
)

type SecretKey struct {
	key ed25519.PrivateKey
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (sec *SecretKey) PublicKey() *PublicKey {
	return &PublicKey{key: sec.key.Public().(ed25519.PublicKey)}
}

func (sec *SecretKey) Sign(msg []byte) []byte {
	return ed25519.Sign(sec.key, msg)
}

func (sec *SecretKey) ToBytes() []byte {
	b := make([]byte, ed25519.PrivateKeySize)
	copy(b, sec.key)
	return b
}

func (pub *PublicKey) VerifySignature(msg []byte, sig []byte) bool {
	return ed25519.Verify(pub.key, msg, sig)
}

func (pub *PublicKey) ToBytes() []byte {
	b := make([]byte, ed25519.PublicKeySize)
	copy(b, pub.key)
	return b
}

func NewPublicKey(b []byte) *PublicKey {
	key := make(ed25519.PublicKey, ed25519.PublicKeySize)
	copy(key, b)
	return &PublicKey{key: key}
}

func NewSecretKey(b []byte) *SecretKey {
	key := make(ed25519.PrivateKey, ed25519.PrivateKeySize)
	copy(key, b)
	return &SecretKey{key: key}
}
