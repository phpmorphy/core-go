package block

import (
	"crypto/sha256"
	"encoding/binary"
	"github.com/umi-top/umi-core/key"
	"github.com/umi-top/umi-core/transaction"
)

const HeaderLength = 167

type Block struct {
	Bytes []byte
}

func NewBlock() *Block {
	b := &Block{Bytes: make([]byte, HeaderLength)}
	b.SetVersion(1)
	return b
}

func FromBytes(b []byte) *Block {
	q := &Block{Bytes: make([]byte, len(b))}
	copy(q.Bytes, b)
	return q
}

func (b *Block) ToBytes() []byte {
	c := make([]byte, len(b.Bytes))
	copy(c, b.Bytes)
	return c
}

func (b *Block) Hash() []byte {
	h := sha256.New()
	h.Write(b.Bytes[0:HeaderLength])
	return h.Sum(nil)
}

func (b *Block) Version() uint8 {
	return b.Bytes[0]
}

func (b *Block) SetVersion(ver uint8) *Block {
	b.Bytes[0] = ver
	return b
}

func (b *Block) PreviousBlockHash() []byte {
	h := make([]byte, 32)
	copy(h, b.Bytes[1:33])
	return h
}

func (b *Block) SetPreviousBlockHash(h []byte) *Block {
	copy(b.Bytes[1:33], h)
	return b
}

func (b *Block) MerkleRootHash() []byte {
	h := make([]byte, 32)
	copy(h, b.Bytes[33:65])
	return h
}

func (b *Block) SetMerkleRootHash(h []byte) *Block {
	copy(b.Bytes[33:65], h)
	return b
}

func (b *Block) Timestamp() uint32 {
	return binary.BigEndian.Uint32(b.Bytes[65:69])
}

func (b *Block) SetTimestamp(t uint32) {
	binary.BigEndian.PutUint32(b.Bytes[65:69], t)
}

func (b *Block) TxCount() uint16 {
	return binary.BigEndian.Uint16(b.Bytes[69:71])
}

func (b *Block) setTxCount(n uint16) *Block {
	binary.BigEndian.PutUint16(b.Bytes[69:71], n)
	return b
}

func (b *Block) PublicKey() *key.PublicKey {
	return key.NewPublicKey(b.Bytes[71:103])
}

func (b *Block) SetPublicKey(k key.PublicKey) *Block {
	copy(b.Bytes[71:103], k.ToBytes())
	return b
}

func (b *Block) Signature() []byte {
	s := make([]byte, 64)
	copy(s, b.Bytes[103:167])
	return s
}

func (b *Block) Sign(k key.SecretKey) *Block {
	copy(b.Bytes[103:167], k.Sign(b.Bytes[0:103]))
	return b
}

func (b *Block) Transaction(idx uint16) *transaction.Transaction {
	offset := idx*transaction.Length + HeaderLength
	return transaction.FromBytes(b.Bytes[offset : offset+transaction.Length])
}

func (b *Block) AppendTransaction(t *transaction.Transaction) *Block {
	b.setTxCount(b.TxCount() + 1)
	b.Bytes = append(b.Bytes, t.ToBytes()...)
	return b
}
