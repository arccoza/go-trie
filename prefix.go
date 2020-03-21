package trie

import (
	"github.com/k0kubun/pp"
	"math"
)

type radix byte

const (
	R2 radix = 1 << iota
	R4
	R16
	R256
)

type Prefix struct {
	rdx, exp, div byte
	ptr, lth int
	key []byte
}

func (p *Prefix) Len() int {
	return len(p.key) * 8 / int(p.exp)
}

func (p *Prefix) Get(idx int) byte {
	div := 8 / int(p.exp)
	ia, ib := idx/int(div), idx%int(div)
	msk := byte(math.Pow(2, float64(p.exp)) - 1)
	pp.Println(div, ia, ib, msk)
	return p.key[ia] >> (int(p.exp) * (div - 1 - ib)) & msk
}

func (p Prefix) Slice(a, b int) Prefix {
	p.ptr, p.lth = a, b
	return p
}

// func (pa *Prefix) Append(pb *Prefix) Prefix {

// }
