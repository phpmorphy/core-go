package transaction

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"github.com/umi-top/umi-core/address"
	"github.com/umi-top/umi-core/key"
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

type Transaction interface {
	Hash() []byte

	Version() uint8
	SetVersion(uint8) Transaction

	Sender() address.Address
	SetSender(address2 address.Address) Transaction

	Recipient() address.Address
	SetRecipient(address2 address.Address) Transaction

	Value() uint64
	SetValue(uint64) Transaction

	Nonce() uint64
	SetNonce(uint64) Transaction

	Signature() []byte
	SetSignature([]byte) Transaction

	Verify() error
	Sign(key.SecretKey) Transaction

	ToBytes() []byte
}

type transaction []byte

func NewTransaction() Transaction {
	tx := make(transaction, Length)
	tx.SetVersion(Basic)
	return tx
}

func FromBytes(b []byte) Transaction {
	t := make(transaction, Length)
	copy(t, b)
	return t
}

func (t transaction) Hash() []byte {
	h := sha256.New()
	h.Write(t)
	return h.Sum(nil)
}

func (t transaction) Version() uint8 {
	return t[0]
}

func (t transaction) SetVersion(ver uint8) Transaction {
	t[0] = ver
	return t
}

func (t transaction) Sender() address.Address {
	return address.FromBytes(t[1:35])
}

func (t transaction) SetSender(adr address.Address) Transaction {
	copy(t[1:35], adr.ToBytes())
	return t
}

func (t transaction) Recipient() address.Address {
	return address.FromBytes(t[35:69])
}

func (t transaction) SetRecipient(adr address.Address) Transaction {
	copy(t[35:69], adr.ToBytes())
	return t
}

func (t transaction) Value() uint64 {
	return binary.BigEndian.Uint64(t[69:77])
}

func (t transaction) SetValue(v uint64) Transaction {
	binary.BigEndian.PutUint64(t[69:77], v)
	return t
}

func (t transaction) Nonce() uint64 {
	return binary.BigEndian.Uint64(t[77:85])
}

func (t transaction) SetNonce(v uint64) Transaction {
	binary.BigEndian.PutUint64(t[77:85], v)
	return t
}

func (t transaction) Signature() []byte {
	s := make([]byte, 64)
	copy(s, t[85:150])
	return s
}

func (t transaction) SetSignature(s []byte) Transaction {
	copy(t[85:150], s)
	return t
}

func (t transaction) Verify() error {
	if !t.Sender().PublicKey().VerifySignature(t[0:85], t[85:149]) {
		return errors.New("invalid signature")
	}
	return nil
}

func (t transaction) Sign(key key.SecretKey) Transaction {
	sig := key.Sign(t[0:85])
	copy(t[85:], sig)
	return t
}

func (t transaction) ToBytes() []byte {
	b := make([]byte, Length)
	copy(b, t)
	return b
}
