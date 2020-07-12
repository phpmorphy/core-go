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
	"encoding/base64"
	"testing"

	"github.com/umi-top/umi-core/key"
)

func TestSecKeyConstructor(t *testing.T) {
	exp, _ := base64.StdEncoding.DecodeString(
		"u1mzvCnmyIbgs8RNM9GGGHOWcBdMvD7GIKC0m9zTFcaGXaAPQMbuPdZ1oAnTCfR/1rHTyC3J5n7x+dlFimHM8w==")
	act := key.NewSecretKey(exp).ToBytes()

	if !bytes.Equal(exp, act) {
		t.Fatalf("Expected: %x, got: %x", exp, act)
	}
}

func TestSecKeyToPubKey(t *testing.T) {
	sec, _ := base64.StdEncoding.DecodeString(
		"u1mzvCnmyIbgs8RNM9GGGHOWcBdMvD7GIKC0m9zTFcaGXaAPQMbuPdZ1oAnTCfR/1rHTyC3J5n7x+dlFimHM8w==")
	exp, _ := base64.StdEncoding.DecodeString("hl2gD0DG7j3WdaAJ0wn0f9ax08gtyeZ+8fnZRYphzPM=")
	act := key.NewSecretKey(sec).PublicKey().ToBytes()

	if !bytes.Equal(exp, act) {
		t.Fatalf("Expected: %x, got: %x", exp, act)
	}
}

func TestSecKeySign(t *testing.T) {
	cases := []struct {
		desc string
		key  string
		msg  string
		sig  string
	}{
		{
			desc: "1st",
			key:  "u1mzvCnmyIbgs8RNM9GGGHOWcBdMvD7GIKC0m9zTFcaGXaAPQMbuPdZ1oAnTCfR/1rHTyC3J5n7x+dlFimHM8w==",
			msg:  "9tJbOqCDGGU0E4F0hYQR88MExTleIverV4iYgQs1bzn+gKmf7HMO3A==",
			sig:  "5a+mePEJlbUzTrqM5uxtVklI4KK+wtxBgkt4jiregPLmqasQ+4kTMu2KQfAJd7IlFYZqH2yM6lZDufXVY6ooBQ==",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			sec, _ := base64.StdEncoding.DecodeString(tc.key)
			msg, _ := base64.StdEncoding.DecodeString(tc.msg)
			sig, _ := base64.StdEncoding.DecodeString(tc.sig)

			act := key.NewSecretKey(sec).Sign(msg)

			if !bytes.Equal(sig, act) {
				t.Fatalf("Expected: %v, got: %v", tc.sig, act)
			}
		})
	}
}
