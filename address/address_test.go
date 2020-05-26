package address

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

type fixture struct {
	adr string
	hex string
}

var fixtures = []fixture{
	{"umi1u3dam33jaf64z4s008g7su62j4za72ljqff9dthsataq8k806nfsgrhdhg", "55a9e45bddc632ea7551560f79d1e8734a9545df2bf2025256aef0eafa03d8efd4d3"},
	{"genesis1nls5a34h9gza3f0p6lwmwhfnq7wrlsdfdj4tcqpu2aaeq22gjaus0vt2zv", "00009fe14ec6b72a05d8a5e1d7ddb75d33079c3fc1a96caabc003c577b9029489779"},
	{"zzz1dagmqppr0kz2y3wwh0zq4k4jezz57r2yp0m7jygxdcwc70l239fsxzulxv", "6b5a6f51b004237d84a245cebbc40adab2c8854f0d440bf7e911066e1d8f3fea8953"},
}

func TestFromBech32(t *testing.T) {
	for _, f := range fixtures {
		b1, _ := hex.DecodeString(f.hex)
		b2 := FromBech32(f.adr).ToBytes()

		if !bytes.Equal(b1, b2) {
			t.Error("For", f.adr, "expected", f.hex, "got", hex.EncodeToString(b2))
		}
	}
}

func TestToBech32(t *testing.T) {
	for _, f := range fixtures {
		b, _ := hex.DecodeString(f.hex)
		adr := FromBytes(b).ToBech32()

		if strings.Compare(f.adr, adr) != 0 {
			t.Error("For", hex.EncodeToString(b), "expected", f.adr, "got", adr)
		}
	}
}
