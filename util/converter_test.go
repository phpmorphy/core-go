package util

import (
	"strings"
	"testing"
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
		ver := PrefixToVersion(f.hrp)
		if f.ver != ver {
			t.Error("For", f.hrp, "expected", f.ver, "got", ver)
		}
	}
}

func TestVersionToPrefix(t *testing.T) {
	for _, f := range fixtures {
		hrp := VersionToPrefix(f.ver)
		if strings.Compare(f.hrp, hrp) != 0 {
			t.Error("For", f.ver, "expected", f.hrp, "got", hrp)
		}
	}
}
