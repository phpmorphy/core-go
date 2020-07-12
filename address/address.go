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

package address

import (
	"encoding/binary"

	"github.com/umi-top/umi-core/key"
	"github.com/umi-top/umi-core/util"
	"github.com/umi-top/umi-core/util/bech32"
)

const (
	Length  int    = 34
	Genesis uint16 = 0
	Umi     uint16 = 21929
)

type Address struct {
	Bytes []byte
}

func NewAddress() *Address {
	a := &Address{Bytes: make([]byte, Length)}
	a.SetVersion(Umi)

	return a
}

func FromBech32(s string) *Address {
	b, err := bech32.Decode(s)
	if err != nil {
		return nil
	}

	return &Address{Bytes: b}
}

func FromBytes(b []byte) *Address {
	a := &Address{Bytes: make([]byte, Length)}
	copy(a.Bytes, b)

	return a
}

func FromKey(key key.Key) *Address {
	a := &Address{Bytes: make([]byte, Length)}
	a.SetVersion(Umi)
	a.SetPublicKey(key.PublicKey())

	return a
}

func (a *Address) Prefix() string {
	return util.VersionToPrefix(a.Version())
}

func (a *Address) SetPrefix(p string) *Address {
	binary.BigEndian.PutUint16(a.Bytes[0:2], util.PrefixToVersion(p))
	return a
}

func (a *Address) PublicKey() *key.PublicKey {
	return key.NewPublicKey(a.Bytes[2:34])
}

func (a *Address) SetPublicKey(p *key.PublicKey) *Address {
	copy(a.Bytes[2:], p.ToBytes())
	return a
}

func (a *Address) Version() uint16 {
	return binary.BigEndian.Uint16(a.Bytes[0:2])
}

func (a *Address) SetVersion(v uint16) *Address {
	binary.BigEndian.PutUint16(a.Bytes[0:2], v)
	return a
}

func (a *Address) ToBech32() string {
	return bech32.Encode(a.Bytes)
}

func (a *Address) ToBytes() []byte {
	b := make([]byte, len(a.Bytes))
	copy(b, a.Bytes)

	return b
}
