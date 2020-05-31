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

func (t *Transaction) Hash() []byte {
	h := sha256.New()
	h.Write(t.Bytes)
	return h.Sum(nil)
}

func (t *Transaction) Version() uint8 {
	return t.Bytes[0]
}

func (t *Transaction) SetVersion(ver uint8) *Transaction {
	t.Bytes[0] = ver
	return t
}

func (t *Transaction) Sender() *address.Address {
	return address.FromBytes(t.Bytes[1:35])
}

func (t *Transaction) SetSender(adr *address.Address) *Transaction {
	copy(t.Bytes[1:35], adr.ToBytes())
	return t
}

func (t *Transaction) Recipient() *address.Address {
	return address.FromBytes(t.Bytes[35:69])
}

func (t *Transaction) SetRecipient(adr *address.Address) *Transaction {
	copy(t.Bytes[35:69], adr.ToBytes())
	return t
}

func (t *Transaction) Value() uint64 {
	return binary.BigEndian.Uint64(t.Bytes[69:77])
}

func (t *Transaction) Prefix() string {
	return util.VersionToPrefix(binary.BigEndian.Uint16(t.Bytes[35:37]))
}

func (t *Transaction) SetValue(v uint64) *Transaction {
	binary.BigEndian.PutUint64(t.Bytes[69:77], v)
	return t
}

func (t *Transaction) Nonce() uint64 {
	return binary.BigEndian.Uint64(t.Bytes[77:85])
}

func (t *Transaction) SetNonce(v uint64) *Transaction {
	binary.BigEndian.PutUint64(t.Bytes[77:85], v)
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

func (t *Transaction) Verify() error {
	if t.Version() == 0 {
		return errors.New("invalid version")
	}

	if t.Version() == 1 {
		if t.Value() > 90_071_992_547_409_91 {
			return errors.New("invalid value")
		}

		if bytes.Equal(t.Sender().Bytes, t.Recipient().Bytes) {
			return errors.New("invalid recipient")
		}
	}

	if t.Version() == 2 || t.Version() == 3 {
		if t.Prefix() == "umi" {
			return errors.New("invalid prefix")
		}

		if t.ProfitPercent() > 500 || t.ProfitPercent() < 100 {
			return errors.New("invalid profit percent")
		}

		if t.FeePercent() > 2000 {
			return errors.New("invalid fee percent")
		}
	}

	if !t.Sender().PublicKey().VerifySignature(t.Bytes[0:85], t.Bytes[85:149]) {
		return errors.New("invalid signature")
	}
	return nil
}

func (t *Transaction) ProfitPercent() uint16 {
	return binary.BigEndian.Uint16(t.Bytes[37:39])
}

func (t *Transaction) FeePercent() uint16 {
	return binary.BigEndian.Uint16(t.Bytes[39:41])
}

func (t *Transaction) Sign(key key.SecretKey) *Transaction {
	sig := key.Sign(t.Bytes[0:85])
	copy(t.Bytes[85:], sig)
	return t
}

func (t *Transaction) ToBytes() []byte {
	b := make([]byte, Length)
	copy(b, t.Bytes)
	return b
}
