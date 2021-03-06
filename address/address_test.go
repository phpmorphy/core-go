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

package address_test

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"github.com/umi-top/umi-core/address"
	"github.com/umi-top/umi-core/key"
)

func TestVersion(t *testing.T) {
	exp := address.Genesis
	act := address.NewAddress().SetVersion(exp).Version()

	if exp != act {
		t.Fatalf("Expected: %d, got: %d", exp, act)
	}
}

func TestPrefix(t *testing.T) {
	exp := "zzz"
	act := address.NewAddress().SetPrefix(exp).Prefix()

	if exp != act {
		t.Fatalf("Expected: %s, got: %s", exp, act)
	}
}

func TestBech32(t *testing.T) {
	exp := "umi1u3dam33jaf64z4s008g7su62j4za72ljqff9dthsataq8k806nfsgrhdhg"
	act := address.FromBech32(exp).ToBech32()

	if exp != act {
		t.Fatalf("Expected: %s, got: %s", exp, act)
	}
}

func TestKey(t *testing.T) {
	rnd := make([]byte, ed25519.PublicKeySize)
	_, _ = rand.Read(rnd)

	pub := key.NewPublicKey(rnd)

	exp := pub.ToBytes()
	act := address.FromKey(pub).PublicKey().ToBytes()

	if !bytes.Equal(exp, act) {
		t.Fatalf("Expected: %x, got: %x", exp, act)
	}
}

func TestBytes(t *testing.T) {
	exp := make([]byte, address.Length)
	_, _ = rand.Read(exp)
	act := address.FromBytes(exp).ToBytes()

	if !bytes.Equal(exp, act) {
		t.Fatalf("Expected: %x, got: %x", exp, act)
	}
}
