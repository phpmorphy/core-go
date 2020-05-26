package address

import (
	"encoding/binary"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/umi-top/umi-core/key"
	"github.com/umi-top/umi-core/util"
)

const Length = 34

type address []byte

type Address interface {
	Version() uint16
	SetVersion(ver uint16) Address

	PublicKey() key.PublicKey

	ToBech32() string

	ToBytes() []byte
}

func NewAddress() Address {
	adr := make(address, Length)
	adr.SetVersion(21929)
	return adr
}

func FromBytes(b []byte) Address {
	adr := make(address, Length)
	copy(adr[:], b)
	return adr
}

func FromBech32(b string) Address {
	hrp, wrd, _ := bech32.Decode(b)
	pub, _ := bech32.ConvertBits(wrd, 5, 8, false)

	ver := util.PrefixToVersion(hrp)
	adr := make(address, Length)
	adr[0] = uint8(ver >> 8)
	adr[1] = uint8(ver & 0xFF)
	copy(adr[2:], pub)

	return Address(adr)
}

func (a address) Version() uint16 {
	return binary.BigEndian.Uint16(a[0:2])
}

func (a address) SetVersion(ver uint16) Address {
	binary.BigEndian.PutUint16(a[0:2], ver)
	return a
}

func (a address) PublicKey() key.PublicKey {
	return key.NewPublicKey(a[2:34])
}

func (a address) ToBytes() []byte {
	b := make([]byte, len(a))
	copy(b, a[:])
	return b
}

func FromPubKey(key key.PublicKey) Address {
	adr := make(address, Length)
	copy(adr[2:], key.ToBytes())
	return adr.SetVersion(21929)
}

func (a address) ToBech32() string {
	ver := uint16(a[1]) + uint16(a[0])<<8
	wrd, _ := bech32.ConvertBits(a[2:], 8, 5, true)
	adr, _ := bech32.Encode(util.VersionToPrefix(ver), wrd)
	return adr
}
