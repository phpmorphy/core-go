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
