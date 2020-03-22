package trie

import (
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

func (p *Prefix) Len() int {
	return p.lth
}

func (p *Prefix) Get(idx int) byte {
	if idx >= p.lth {
		panic(fmt.Errorf("index out of range %v", idx)
	}

	idx += p.ptr
	ia, ib := idx/int(p.div), idx%int(p.div)
	msk := byte(p.rdx - 1)
	pp.Println(p.div, ia, ib, msk)
	return p.key[ia] >> (int(p.exp) * (int(p.div) - 1 - ib)) & msk
}

// Use Receiver type of Prefix instead of *Prefix
// so that we get a simple copy of p.
func (p Prefix) Slice(a, b int) Prefix {
	p.ptr, p.lth = a, b
	return p
}

// func (pa *Prefix) Append(pb *Prefix) Prefix {

// }
