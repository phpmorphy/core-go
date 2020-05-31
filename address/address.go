package address

import (
	"encoding/binary"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/umi-top/umi-core/key"
	"github.com/umi-top/umi-core/util"
)

const Length = 34

type Address struct {
	Bytes []byte
}

func NewAddress() *Address {
	adr := &Address{Bytes: make([]byte, Length)}
	adr.SetVersion(21929)
	return adr
}

func FromBytes(b []byte) *Address {
	adr := &Address{Bytes: make([]byte, Length)}
	copy(adr.Bytes, b)
	return adr
}

func FromBech32(b string) *Address {
	hrp, wrd, _ := bech32.Decode(b)
	pub, _ := bech32.ConvertBits(wrd, 5, 8, false)

	ver := util.PrefixToVersion(hrp)

	adr := &Address{Bytes: make([]byte, Length)}
	adr.Bytes[0] = uint8(ver >> 8)
	adr.Bytes[1] = uint8(ver & 0xFF)
	copy(adr.Bytes[2:], pub)

	return adr
}

func FromPubKey(key key.PublicKey) *Address {
	adr := &Address{Bytes: make([]byte, Length)}
	adr.SetVersion(21929)
	copy(adr.Bytes[2:], key.ToBytes())
	return adr
}

func (a *Address) Version() uint16 {
	return binary.BigEndian.Uint16(a.Bytes[0:2])
}

func (a *Address) Prefix() string {
	return util.VersionToPrefix(a.Version())
}

func (a *Address) SetVersion(ver uint16) *Address {
	binary.BigEndian.PutUint16(a.Bytes[0:2], ver)
	return a
}

func (a *Address) PublicKey() *key.PublicKey {
	return key.NewPublicKey(a.Bytes[2:34])
}

func (a *Address) ToBytes() []byte {
	b := make([]byte, len(a.Bytes))
	copy(b, a.Bytes)
	return b
}

func (a *Address) ToBech32() string {
	ver := uint16(a.Bytes[1]) + uint16(a.Bytes[0])<<8
	wrd, _ := bech32.ConvertBits(a.Bytes[2:], 8, 5, true)
	adr, _ := bech32.Encode(util.VersionToPrefix(ver), wrd)
	return adr
}
