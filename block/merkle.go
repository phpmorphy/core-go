package block

import (
	"crypto/sha256"
	"sync"
)

var mpool = sync.Pool{
	New: func() interface{} {
		return &[65535][32]byte{}
	},
}

func (b *Block) CalculateMerkleRoot() []byte {
	c := mpool.Get().(*[65535][32]byte)
	h := sha256.New()

	// step 1

	j := b.TxCount()

	for i := uint16(0); i < j; i++ {
		_, _ = h.Write(b.Transaction(i).Bytes)
		copy(c[i][0:32], h.Sum(nil))
		h.Reset()
	}

	// step 2

	min := func(a, b int) int {
		if a > b {
			return b
		}

		return a
	}

	next := func(count int) (nextCount, maxIdx int) {
		maxIdx = count - 1

		if count > 2 {
			count += count % 2
		}

		nextCount = count / 2

		return nextCount, maxIdx
	}

	for n, m := next(int(j)); n > 0; n, m = next(n) {
		for i := 0; i < n; i++ {
			k1 := i * 2
			k2 := min(k1+1, m)
			_, _ = h.Write(c[k1][0:32])
			_, _ = h.Write(c[k2][0:32])
			copy(c[i][0:32], h.Sum(nil))
			h.Reset()
		}
	}

	mpool.Put(c)

	return c[0][:]
}
