package trie

import (
	"fmt"
	"math"
	"github.com/k0kubun/pp"
)

type radix uint16

const (
	R2   radix = 2
	R4   radix = 4
	R16  radix = 16
	R256 radix = 256
)

func NewPrefix(rdx radix, key []byte) Prefix {
	exp := int(math.Log2(float64(rdx)))
	div := 8 / exp
	lth := len(key) * 8 / exp
	return Prefix{rdx: rdx, exp: byte(exp), div: byte(div), ptr: 0, lth: lth, key: key}
}

type Prefix struct {
	rdx      radix
	exp, div byte
	ptr, lth int
	key      []byte
}

func (p *Prefix) relIdx(i int) (int, int) {
	return i/int(p.div), i%int(p.div)
}

func (p *Prefix) supIdx(i int) int {
	return i/int(p.div)
}

func (p *Prefix) subIdx(i int) int {
	return i%int(p.div)
}

func (p *Prefix) supLen(l int) int {
	return l/int(p.div)
}

func (p *Prefix) Len() int {
	return p.lth
}

func (p *Prefix) Get(idx int) byte {
	if idx >= p.lth {
		panic(fmt.Errorf("Prefix.Get index out of range %v", idx))
	}

	idx += p.ptr
	ia, ib := idx/int(p.div), idx%int(p.div)
	msk := byte(p.rdx - 1)
	pp.Println(p.div, ia, ib, msk)
	return p.key[ia] >> (int(p.exp) * (int(p.div) - 1 - ib)) & msk
}

// Use receiver type of Prefix instead of *Prefix
// so that we get a simple, shallow copy of p.
func (p Prefix) Slice(args ...int) Prefix {
	a, b, nargs := 0, p.lth, len(args)
	if nargs > 0 {
		a = args[0]
	}
	if nargs > 1 {
		b = args[1]
	}

	if a < 0 || b < 0 {
		panic(fmt.Errorf("Prefix.Slice invalid index %v, %v (index must be non-negative)", a, b))
	} else if a > b {
		panic(fmt.Errorf("Prefix.Slice invalid index %v > %v", a, b))
	} else if a > p.lth || b > p.lth {
		panic(fmt.Errorf("Prefix.Slice bounds out of range %v, %v", a, b))
	} else {
		p.ptr, p.lth = p.ptr + a, b - a
	}

	return p
}

func bitMask(exp, n int) byte {
	return byte(math.Pow(2, float64(exp*n)) - 1) << (exp * (8/exp - n))
}

func bitCompare(a, b byte, exp, n int) (i int) {
	n = MinInt(n, 8/exp)
	for i = n - 1; i >= 0; i-- {
		msk := bitMask(exp, i + 1)
		if (a ^ b) & msk == 0 {
			break
		}
	}
	return
}
