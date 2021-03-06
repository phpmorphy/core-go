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

package util

import (
	"strings"
)

func PrefixToVersion(p string) uint16 {
	if strings.Compare(p, "genesis") == 0 {
		return 0
	}

	return (uint16(p[0]-96) << 10) + (uint16(p[1]-96) << 5) + uint16(p[2]-96)
}

func VersionToPrefix(v uint16) string {
	if v == 0 {
		return "genesis"
	}

	p := make([]byte, 3)
	p[0] = uint8(v&0x7C00>>10) + 96
	p[1] = uint8(v&0x03E0>>5) + 96
	p[2] = uint8(v&0x001F) + 96

	return string(p)
}
