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

package bech32_test

import (
	"strings"
	"testing"

	"github.com/umi-top/umi-core/util/bech32"
)

func TestBech32(t *testing.T) {
	tests := []struct {
		str   string
		valid bool
	}{
		{"A12UEL5L", true},
		{"an83characterlonghumanreadablepartthatcontainsthenumber1andtheexcludedcharactersbio1tt5tgs", true},
		{"abcdef1qpzry9x8gf2tvdw0s3jn54khce6mua7lmqqqxw", true},
		{"11qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqc8247j", true},
		{"split1checkupstagehandshakeupstreamerranterredcaperred2y9e3w", true},
		{"split1checkupstagehandshakeupstreamerranterredcaperred2y9e2w", false},                                // invalid checksum
		{"s lit1checkupstagehandshakeupstreamerranterredcaperredp8hs2p", false},                                // invalid character (space) in hrp
		{"spl" + string(127) + "t1checkupstagehandshakeupstreamerranterredcaperred2y9e3w", false},              // invalid character (DEL) in hrp
		{"split1cheo2y9e2w", false},                                                                            // invalid character (o) in data part
		{"split1a2y9w", false},                                                                                 // too short data part
		{"1checkupstagehandshakeupstreamerranterredcaperred2y9e3w", false},                                     // empty hrp
		{"11qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqsqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqc8247j", false}, // too long
	}

	for _, test := range tests {
		str := test.str
		hrp, decoded, err := bech32.DecodeRef(str)
		if !test.valid {
			// Invalid string decoding should result in error.
			if err == nil {
				t.Error("expected decoding to fail for invalid string", test.str)
			}
			continue
		}

		// Valid string decoding should result in no error.
		if err != nil {
			t.Errorf("expected string to be valid bech32: %v", err)
		}

		// Check that it encodes to the same string
		encoded, err := bech32.EncodeRef(hrp, decoded)
		if err != nil {
			t.Errorf("encoding failed: %v", err)
		}

		if encoded != strings.ToLower(str) {
			t.Errorf("expected data to encode to %v, but got %v",
				str, encoded)
		}

		// Flip a bit in the string an make sure it is caught.
		pos := strings.LastIndexAny(str, "1")
		flipped := str[:pos+1] + string((str[pos+1] ^ 1)) + str[pos+2:]
		_, err = bech32.Decode(flipped)
		if err == nil {
			t.Error("expected decoding to fail")
		}
	}
}
