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

package util_test

import (
	"strings"
	"testing"

	"github.com/umi-top/umi-core/util"
)

type fixture struct {
	hrp string
	ver uint16
}

var fixtures = []fixture{
	{"genesis", 0},
	{"aaa", 1057},
	{"zzz", 27482},
	{"umi", 21929},
}

func TestPrefixToVersion(t *testing.T) {
	for _, f := range fixtures {
		ver := util.PrefixToVersion(f.hrp)
		if f.ver != ver {
			t.Error("For", f.hrp, "expected", f.ver, "got", ver)
		}
	}
}

func TestVersionToPrefix(t *testing.T) {
	for _, f := range fixtures {
		hrp := util.VersionToPrefix(f.ver)
		if strings.Compare(f.hrp, hrp) != 0 {
			t.Error("For", f.ver, "expected", f.hrp, "got", hrp)
		}
	}
}
