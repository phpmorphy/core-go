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

package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"

	"github.com/umi-top/umi-core/address"
	"github.com/umi-top/umi-core/key"
	"github.com/umi-top/umi-core/util"
)

const Length = 150
const (
	Genesis = iota
	Basic
	CreateSmartContract
	UpdateSmartContract
	UpdateProfitAddress
	UpdateFeeAddress
	CreateTransitAddress
	DeleteTransitAddress
)

var (
	ErrInvalidVersion       = errors.New("transaction: invalid version")
	ErrInvalidValue         = errors.New("transaction: invalid value")
	ErrInvalidRecipient     = errors.New("transaction: invalid recipient")
	ErrInvalidPrefix        = errors.New("transaction: invalid prefix")
	ErrInvalidFeePercent    = errors.New("transaction: invalid fee percent")
	ErrInvalidProfitPercent = errors.New("transaction: invalid profit percent")
	ErrInvalidSignature     = errors.New("transaction: invalid signature")
)

type Transaction struct {
	Bytes []byte
}

func NewTransaction() *Transaction {
	tx := &Transaction{Bytes: make([]byte, Length)}
	tx.SetVersion(Basic)

	return tx
}

func FromBytes(b []byte) *Transaction {
	t := &Transaction{Bytes: make([]byte, Length)}
	copy(t.Bytes, b)

	return t
}

func (t *Transaction) FeePercent() uint16 {
	return binary.BigEndian.Uint16(t.Bytes[39:41])
}

func (t *Transaction) Hash() []byte {
	h := sha256.New()
	_, _ = h.Write(t.Bytes)

	return h.Sum(nil)
}

func (t *Transaction) Name() string {
	return ""
}

func (t *Transaction) SetName(n string) *Transaction {
	return t
}

func (t *Transaction) Nonce() uint64 {
	return binary.BigEndian.Uint64(t.Bytes[77:85])
}

func (t *Transaction) SetNonce(v uint64) *Transaction {
	binary.BigEndian.PutUint64(t.Bytes[77:85], v)
	return t
}

func (t *Transaction) Prefix() string {
	return util.VersionToPrefix(binary.BigEndian.Uint16(t.Bytes[35:37]))
}

func (t *Transaction) SetPrefix(p string) *Transaction {
	binary.BigEndian.PutUint16(t.Bytes[35:37], util.PrefixToVersion(p))
	return t
}

func (t *Transaction) ProfitPercent() uint16 {
	return binary.BigEndian.Uint16(t.Bytes[37:39])
}

func (t *Transaction) SetProfitPercent(v uint16) *Transaction {
	binary.BigEndian.PutUint16(t.Bytes[37:39], v)
	return t
}

func (t *Transaction) Recipient() *address.Address {
	return address.FromBytes(t.Bytes[35:69])
}

func (t *Transaction) SetRecipient(a *address.Address) *Transaction {
	copy(t.Bytes[35:69], a.ToBytes())
	return t
}

func (t *Transaction) Sender() *address.Address {
	return address.FromBytes(t.Bytes[1:35])
}

func (t *Transaction) SetSender(a *address.Address) *Transaction {
	copy(t.Bytes[1:35], a.ToBytes())
	return t
}

func (t *Transaction) Signature() []byte {
	s := make([]byte, 64)
	copy(s, t.Bytes[85:150])

	return s
}

func (t *Transaction) SetSignature(s []byte) *Transaction {
	copy(t.Bytes[85:150], s)
	return t
}

func (t *Transaction) Version() uint8 {
	return t.Bytes[0]
}

func (t *Transaction) SetVersion(v uint8) *Transaction {
	t.Bytes[0] = v
	return t
}

func (t *Transaction) Value() uint64 {
	return binary.BigEndian.Uint64(t.Bytes[69:77])
}

func (t *Transaction) SetValue(v uint64) *Transaction {
	binary.BigEndian.PutUint64(t.Bytes[69:77], v)
	return t
}

func (t *Transaction) Sign(k key.SecretKey) *Transaction {
	sig := k.Sign(t.Bytes[0:85])
	copy(t.Bytes[85:], sig)

	return t
}

func (t *Transaction) ToBytes() []byte {
	b := make([]byte, Length)
	copy(b, t.Bytes)

	return b
}

func (t *Transaction) Verify() error {
	if t.Version() == 0 {
		return ErrInvalidVersion
	}

	if t.Version() == 1 {
		if t.Value() > 90_071_992_547_409_91 {
			return ErrInvalidValue
		}

		if bytes.Equal(t.Sender().Bytes, t.Recipient().Bytes) {
			return ErrInvalidRecipient
		}
	}

	if t.Version() == 2 || t.Version() == 3 {
		if t.Prefix() == "umi" {
			return ErrInvalidPrefix
		}

		if t.ProfitPercent() > 500 || t.ProfitPercent() < 100 {
			return ErrInvalidProfitPercent
		}

		if t.FeePercent() > 2000 {
			return ErrInvalidFeePercent
		}
	}

	if !t.Sender().PublicKey().VerifySignature(t.Bytes[0:85], t.Bytes[85:149]) {
		return ErrInvalidSignature
	}

	return nil
}
