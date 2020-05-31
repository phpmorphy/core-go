package block

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"github.com/umi-top/umi-core/transaction"
	"math/rand"
	"strings"
	"testing"
)

var tx0 []byte
var tx1 []byte
var blk []byte
var hshz []byte
var prev []byte
var merk []byte
var pubk []byte
var sign []byte

var timez uint32 = 1590492060

func init() {
	tx0, _ = base64.StdEncoding.DecodeString("AVWpGa01ZBN38MPUOD55BNcloJpJema+xYsGMlX594oWXzhrWgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/////////////////////w08w/24Q/3uOCzXicm57KtJNAbHVwzWxZhM7+09FmvMvHEjGshdu5LJr87h8F4apBnnrKe/SWu7R6SmMQrdVwwA")
	tx1, _ = base64.StdEncoding.DecodeString("AGtaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVqRmtNWQTd/DD1Dg+eQTXJaCaSXpmvsWLBjJV+feKFl84f/////////9//////////5kH/KVsF+iXju9URnu6ULJtPRTScGPlJui+9RqMhpVc1tt4G9suxok/OkA54Dd8qTDkSTKb6T8h5GuFrI87VA4A")
	blk, _ = base64.StdEncoding.DecodeString("AT60JvnMp7hgDKgK0vMUybIrOYquF1sWaWDT25p+Ae8kBDK4ic1/JAgO5v6E2JMUauCSrDoNZUNZJ0cETHzHWi1ezPucAAIZrTVkE3fww9Q4PnkE1yWgmkl6Zr7FiwYyVfn3ihZfOKF8cN8DaKNYYfOTxdaAydWaL17RKwvcgFBW6fjIgJX01w3lt8NbAsiBKn/LZWL2DkXPnGCEqngbYtKqfojC1w0BVakZrTVkE3fww9Q4PnkE1yWgmkl6Zr7FiwYyVfn3ihZfOGtaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/////////////////////DTzD/bhD/e44LNeJybnsq0k0BsdXDNbFmEzv7T0Wa8y8cSMayF27ksmvzuHwXhqkGeesp79Ja7tHpKYxCt1XDAAAa1oAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFWpGa01ZBN38MPUOD55BNcloJpJema+xYsGMlX594oWXzh//////////3//////////mQf8pWwX6JeO71RGe7pQsm09FNJwY+Um6L71GoyGlVzW23gb2y7GiT86QDngN3ypMORJMpvpPyHka4WsjztUDgA=")
	hshz, _ = hex.DecodeString("d5365021b10c2aa0314ec4a1e72820389b01bf9aa4630492a65ce53b16319239")
	prev, _ = hex.DecodeString("3eb426f9cca7b8600ca80ad2f314c9b22b398aae175b166960d3db9a7e01ef24")
	merk, _ = hex.DecodeString("0432b889cd7f24080ee6fe84d893146ae092ac3a0d6543592747044c7cc75a2d")
	pubk, _ = hex.DecodeString("19ad35641377f0c3d4383e7904d725a09a497a66bec58b063255f9f78a165f38")
	sign, _ = hex.DecodeString("a17c70df0368a35861f393c5d680c9d59a2f5ed12b0bdc805056e9f8c88095f4d70de5b7c35b02c8812a7fcb6562f60e45cf9c6084aa781b62d2aa7e88c2d70d")
}

func TestGetHash(t *testing.T) {
	hsh := FromBytes(blk).Hash()
	if !bytes.Equal(hshz, hsh) {
		t.Error("Expected", hex.EncodeToString(hshz), "got", hex.EncodeToString(hsh))
	}
}

func TestGetVersion(t *testing.T) {
	ver := FromBytes(blk).Version()
	if ver != 1 {
		t.Error("Expected", 1, "got", ver)
	}
}

func TestGetPreviousBlockHash(t *testing.T) {
	prv := FromBytes(blk).PreviousBlockHash()
	if !bytes.Equal(prev, prv) {
		t.Error("Expected", hex.EncodeToString(prev), "got", hex.EncodeToString(prv))
	}
}

func TestGetMerkleRootHash(t *testing.T) {
	mrk := FromBytes(blk).MerkleRootHash()
	if !bytes.Equal(merk, mrk) {
		t.Error("Expected", hex.EncodeToString(merk), "got", hex.EncodeToString(mrk))
	}
}

func TestGetTimestamp(t *testing.T) {
	tm := FromBytes(blk).Timestamp()
	if timez != tm {
		t.Error("Expected", timez, "got", tm)
	}
}

func TestGetTxCount(t *testing.T) {
	cnt := FromBytes(blk).TxCount()
	if cnt != 2 {
		t.Error("Expected", 2, "got", cnt)
	}
}

func TestGetPublicKey(t *testing.T) {
	pk := FromBytes(blk).PublicKey().ToBytes()
	if !bytes.Equal(pubk, pk) {
		t.Error("Expected", hex.EncodeToString(pubk), "got", hex.EncodeToString(pk))
	}
}

func TestGetSignature(t *testing.T) {
	sg := FromBytes(blk).Signature()
	if !bytes.Equal(sign, sg) {
		t.Error("Expected", hex.EncodeToString(sign), "got", hex.EncodeToString(sg))
	}
}

func TestGetTransaction(t *testing.T) {
	t0 := FromBytes(blk).Transaction(0).ToBytes()
	if !bytes.Equal(tx0, t0) {
		t.Error("Expected", hex.EncodeToString(tx0), "got", hex.EncodeToString(t0))
	}

	t1 := FromBytes(blk).Transaction(1).ToBytes()
	if !bytes.Equal(tx1, t1) {
		t.Error("Expected", hex.EncodeToString(tx1), "got", hex.EncodeToString(t1))
	}
}

func TestAppendTransaction(t *testing.T) {
	rn := make([]byte, 150)
	rand.Read(rn)
	tx0 := transaction.FromBytes(rn)
	rand.Read(rn)
	tx1 := transaction.FromBytes(rn)
	rand.Read(rn)
	tx2 := transaction.FromBytes(rn)

	bl := NewBlock()
	bl.AppendTransaction(tx0)
	bl.AppendTransaction(tx1)
	bl.AppendTransaction(tx2)

	cnt := bl.TxCount()
	if cnt != 3 {
		t.Error("Expected 3 got", cnt)
	}

	ln := len(bl.ToBytes())
	if ln != 617 {
		t.Error("Expected 617 got", ln)
	}

	xx1 := bl.Transaction(1).Bytes

	if !bytes.Equal(xx1, tx1.Bytes) {
		t.Error("Expected", hex.EncodeToString(tx1.Bytes), "got", hex.EncodeToString(xx1))
	}
}

func TestMerkel(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())

	rn := make([]byte, 150)
	//rand.Read(rn)
	tx0 := transaction.FromBytes(rn)
	//rand.Read(rn)
	tx1 := transaction.FromBytes(rn)
	//rand.Read(rn)
	tx2 := transaction.FromBytes(rn)

	bl := NewBlock()
	bl.AppendTransaction(tx0)
	bl.AppendTransaction(tx1)
	bl.AppendTransaction(tx2)
	//bl.AppendTransaction(tx2)
	bl.SetMerkleRootHash(bl.CalculateMerkleRoot())

	if strings.Compare("fd6e1bd8870d9c5c408aaafd8a33235a4e03d74a67c7a46b91af8de493dc6aaf", string(bl.MerkleRootHash())) == 0 {
		t.Error("Expected", "fd6e1bd8870d9c5c408aaafd8a33235a4e03d74a67c7a46b91af8de493dc6aaf", "got", hex.EncodeToString(bl.MerkleRootHash()))
	}
}
