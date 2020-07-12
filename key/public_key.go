// Copyright (c) 2020 UMI
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package key

import (
	"crypto/ed25519"
)

type PublicKey struct {
	key ed25519.PublicKey
}

func NewPublicKey(b []byte) *PublicKey {
	key := make(ed25519.PublicKey, ed25519.PublicKeySize)
	copy(key, b)

	return &PublicKey{key: key}
}

func (p *PublicKey) PublicKey() *PublicKey {
	return p
}

func (p *PublicKey) VerifySignature(sig []byte, msg []byte) bool {
	return ed25519.Verify(p.key, msg, sig)
}

func (p *PublicKey) ToBytes() []byte {
	b := make([]byte, ed25519.PublicKeySize)
	copy(b, p.key)

	return b
}
