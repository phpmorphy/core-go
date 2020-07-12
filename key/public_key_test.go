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

package key_test

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"testing"

	"github.com/umi-top/umi-core/key"
)

func TestPublicKey(t *testing.T) {
	rnd := make([]byte, ed25519.PublicKeySize)
	_, _ = rand.Read(rnd)

	pub := key.NewPublicKey(rnd)

	exp := pub.ToBytes()
	act := pub.PublicKey().ToBytes()

	if !bytes.Equal(exp, act) {
		t.Fatalf("Expected: %x, got: %x", exp, act)
	}
}

func TestPubKeyVerifySignature(t *testing.T) {
	cases := []struct {
		desc string
		key  string
		msg  string
		sig  string
		exp  bool
	}{
		{
			desc: "valid",
			key:  "oD7CzMxo3UYjXg/URrZPluOSOjAbzYVxIXDyONlR5pI=",
			msg:  "K7B3Y9MILKseAlBDkjuwjc48NT3vMWrhixyh7diP8O8B",
			sig:  "7mVMWMzqHgy+I9GSlS0XFAXV1IjGmeZhlDQOBMjrwua7EULygNIKgkiQ2h6kSeDq76tBomoaPbc8faFYwNO0Dg==",
			exp:  true,
		},
		{
			desc: "invalid",
			key:  "MUAmRXK6+YHhASTdWN7Xx2keYPG1V+VoVIXN3RNIBSE=",
			msg:  "UGffQxqOxfMcTcWRVaRklCS/MNme5j2IzUh0J8ksbPTd",
			sig:  "kQ7z0+PDJBaQeihqd0hForqdBTVr8mrAO0Sg6RWMi3EbFSdHMVVicqSZVthcr+gjpnjjdOiKbxembcCoXAieCQ==",
			exp:  false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			pub, _ := base64.StdEncoding.DecodeString(tc.key)
			msg, _ := base64.StdEncoding.DecodeString(tc.msg)
			sig, _ := base64.StdEncoding.DecodeString(tc.sig)

			act := key.NewPublicKey(pub).VerifySignature(sig, msg)

			if act != tc.exp {
				t.Fatalf("Expected: %v, got: %v", tc.exp, act)
			}
		})
	}
}
