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

package transaction_test

//import (
//	"bytes"
//	"crypto/rand"
//	"encoding/base64"
//	"encoding/binary"
//	"encoding/hex"
//	"github.com/umi-top/umi-core/address"
//	"github.com/umi-top/umi-core/transaction"
//	"testing"
//)
//
//var tx []byte
//var sndr []byte
//var rcpt []byte
//var sign []byte
//var hash []byte
//var value uint64 = 18446744073709551615
//var nonce uint64 = 12345678987654321
//
//func init() {
//	tx, _ = base64.StdEncoding.DecodeString(
//		"AVWpQLu8E0IuDJhUqUG/HFPaTVOjZOi8AGivE7Zm4i92yvBrWgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//////////8AK" +
//			"9xUYpH0sdBzcy2UImcJ/bTHNLM14hjOyfREoMTctoxSgOn/vg1fFugBPYMLRSSOikeZNI3YY8IO+uHnXCazueC90YjrlA4A")
//	sndr, _ = hex.DecodeString("55a940bbbc13422e0c9854a941bf1c53da4d53a364e8bc0068af13b666e22f76caf0")
//	rcpt, _ = hex.DecodeString("6b5a0000000000000000000000000000000000000000000000000000000000000000")
//	sign, _ = hex.DecodeString("d073732d94226709fdb4c734b335e218cec9f444a0c4dcb68c5280e9ffbe0d5f16e8013d" +
//		"830b45248e8a4799348dd863c20efae1e75c26b3b9e0bdd188eb940e")
//	hash, _ = hex.DecodeString("d4ea3e4de848e55161ac31a43a64e42743387d02a25c0b9111c7d4efed0790c3")
//}
//
//func TestGetHash(t *testing.T) {
//	hsh := transaction.FromBytes(tx).Hash()
//	if !bytes.Equal(hash, hsh) {
//		t.Error("Expected", hex.EncodeToString(hash), "got", hex.EncodeToString(hsh))
//	}
//}
//
//func TestGetVersion(t *testing.T) {
//	ver := transaction.FromBytes(tx).Version()
//	if ver != 1 {
//		t.Error("Expected", 1, "got", ver)
//	}
//}
//
//func TestSetVersion(t *testing.T) {
//	tx := transaction.NewTransaction()
//	if tx.Version() != transaction.Basic {
//		t.Error("По умолчанию должна создаваться basic-транзакция")
//	}
//
//	vr := tx.SetVersion(transaction.UpdateFeeAddress).Version()
//	if vr != transaction.UpdateFeeAddress {
//		t.Error("Expected", transaction.UpdateFeeAddress, "got", vr)
//	}
//}
//
//func TestGetSender(t *testing.T) {
//	snd := transaction.FromBytes(tx).Sender().ToBytes()
//	if !bytes.Equal(sndr, snd) {
//		t.Error("Expected", hex.EncodeToString(sndr), "got", hex.EncodeToString(snd))
//	}
//}
//
//func TestSetSender(t *testing.T) {
//	tx := transaction.NewTransaction()
//	if !bytes.Equal(tx.Sender().ToBytes(), make([]byte, 34)) {
//		t.Error("По умолчанию должн быть пустой отправитель")
//	}
//
//	pub := make([]byte, 34)
//	_, _ = rand.Read(pub)
//	copy(pub[0:2], []byte{0, 0})
//	adr := address.FromBytes(pub)
//
//	tx.SetSender(adr)
//
//	res := tx.Sender().ToBytes()
//	if !bytes.Equal(res, pub) {
//		t.Error("Expected", hex.EncodeToString(pub), "got", hex.EncodeToString(res))
//	}
//}
//
//func TestGetRecipient(t *testing.T) {
//	rcp := transaction.FromBytes(tx).Recipient().ToBytes()
//	if !bytes.Equal(rcpt, rcp) {
//		t.Error("Expected", hex.EncodeToString(rcpt), "got", hex.EncodeToString(rcp))
//	}
//}
//
//func TestSetRecipient(t *testing.T) {
//	tx := transaction.NewTransaction()
//	if !bytes.Equal(tx.Recipient().ToBytes(), make([]byte, 34)) {
//		t.Error("По умолчанию должн быть пустой отправитель")
//	}
//
//	pub := make([]byte, 34)
//	_, _ = rand.Read(pub)
//	copy(pub[0:2], []byte{0, 0})
//	adr := address.FromBytes(pub)
//	tx.SetRecipient(adr)
//	res := tx.Recipient().ToBytes()
//
//	if !bytes.Equal(res, pub) {
//		t.Error("Expected", hex.EncodeToString(pub), "got", hex.EncodeToString(res))
//	}
//}
//
//func TestGetValue(t *testing.T) {
//	val := transaction.FromBytes(tx).Value()
//	if value != val {
//		t.Error("Expected", value, "got", val)
//	}
//}
//
//func TestSetValue(t *testing.T) {
//	tx := transaction.NewTransaction()
//	if tx.Value() != 0 {
//		t.Error("По умолчанию сумма долнжна быть 0")
//	}
//
//	pub := make([]byte, 8)
//	_, _ = rand.Read(pub)
//	val := binary.BigEndian.Uint64(pub)
//	tx.SetValue(val)
//	res := tx.Value()
//
//	if res != val {
//		t.Error("Expected", val, "got", res)
//	}
//}
//
//func TestGetNonce(t *testing.T) {
//	ncn := transaction.FromBytes(tx).Nonce()
//	if nonce != ncn {
//		t.Error("Expected", nonce, "got", ncn)
//	}
//}
//
//func TestSetNonce(t *testing.T) {
//	tx := transaction.NewTransaction()
//	if tx.Nonce() != 0 {
//		t.Error("По умолчанию nonce долнжна быть 0")
//	}
//
//	pub := make([]byte, 8)
//	_, _ = rand.Read(pub)
//	non := binary.BigEndian.Uint64(pub)
//	tx.SetNonce(non)
//	res := tx.Nonce()
//
//	if res != non {
//		t.Error("Expected", non, "got", res)
//	}
//}
//
//func TestGetSignature(t *testing.T) {
//	sig := transaction.FromBytes(tx).Signature()
//	if !bytes.Equal(sign, sig) {
//		t.Error("Expected", hex.EncodeToString(sign), "got", hex.EncodeToString(sig))
//	}
//}
//
//func TestSetSignature(t *testing.T) {
//	tx := transaction.NewTransaction()
//	if !bytes.Equal(tx.Signature(), make([]byte, 64)) {
//		t.Error("По умолчанию должн быть пустая подпись")
//	}
//
//	sig := make([]byte, 64)
//	_, _ = rand.Read(sig)
//	tx.SetSignature(sig)
//	res := tx.Signature()
//
//	if !bytes.Equal(res, sig) {
//		t.Error("Expected", hex.EncodeToString(sig), "got", hex.EncodeToString(res))
//	}
//}
//
//func TestVerify(t *testing.T) {
//	trx := transaction.FromBytes(tx)
//	err := trx.Verify()
//
//	if err != nil {
//		t.Error("Expected nil, got", err)
//	}
//
//	trx.SetNonce(100500)
//	err = trx.Verify()
//
//	if err == nil {
//		t.Error("Expected not nil")
//	}
//}
