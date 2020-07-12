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

type SecretKey struct {
	key ed25519.PrivateKey
}

func NewSecretKey(b []byte) *SecretKey {
	key := make(ed25519.PrivateKey, ed25519.PrivateKeySize)
	copy(key, b)

	return &SecretKey{key: key}
}

func (s *SecretKey) PublicKey() *PublicKey {
	return &PublicKey{key: s.key.Public().(ed25519.PublicKey)}
}

func (s *SecretKey) Sign(msg []byte) []byte {
	return ed25519.Sign(s.key, msg)
}

func (s *SecretKey) ToBytes() []byte {
	b := make([]byte, ed25519.PrivateKeySize)
	copy(b, s.key)

	return b
}
