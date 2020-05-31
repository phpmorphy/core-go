package key

import (
	"testing"
)

func TestIntMinTableDriven(t *testing.T) {
	seed := make([]byte, 64)
	sec := NewSecretKey(seed)
	pub := sec.PublicKey()

	msg := make([]byte, 32)
	sig := sec.Sign(msg)

	if pub.VerifySignature(msg, sig) {
		t.Error("Expected", "true", "got", "false")
	}
}
