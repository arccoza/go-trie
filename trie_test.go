package trie

import (
	"os"
	"testing"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// buf := make([]byte, 4)
	// t := NewItemTrie()
	// // t.Put([]byte("remember"), 0x0001)
	// // t.Put([]byte("remain"), 0x0002)
	// // t.Put([]byte("rem"), 0x0004)

	// pp.Println(0x0000)
	// t.Put(Uint32ToBytes(buf, 0x0000), 0x0001)
	// pp.Println(0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0001), 0x0001)
	// pp.Println(0x0002)
	// t.Put(Uint32ToBytes(buf, 0x0002), 0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0003), 0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0004), 0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0005), 0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0006), 0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0007), 0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0008), 0x0001)
	// t.Put(Uint32ToBytes(buf, 0x0009), 0x0001)

	// pp.Println(0x001F)
	// t.Put(Uint32ToBytes(buf, 0x001F), 0x0001)
	// pp.Println(0x061C)
	// t.Put(Uint32ToBytes(buf, 0x061C), 0x0001)

	// pp.Println(t)


	// p := NewPrefix(R16, []byte{0xAB, 0xCD})
	// // p := Prefix{exp: 4, key: []byte{0xAB, 0xCD}}
	// pp.Println(p, R2, R4, R16, R256)
	// pp.Println(p.Get(0), p.Get(1), p.Len())
	// p2 := p.Slice(1, 4)
	// // p2 = p2.Slice(1, 3)
	// pp.Println(p2)
	// pp.Println(p2.Get(0), p2.Get(2), p2.Len())

	// pp.Println(bitCompare(0xFF, 0xFF, 4, 20))

	p := NewPrefix(R16, []byte{0xAB, 0xCD})
	pp.Println(p.relIdx(3))
	pp.Println(p.absIdx(1, 1))


	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestGet_R16(t *testing.T) {
	p := NewPrefix(R16, []byte{0xAB, 0xCD})
	want := []byte{0x0A, 0x0B, 0x0C, 0x0D}
	got := []byte{p.Get(0), p.Get(1), p.Get(2), p.Get(3)}

	assert.Equalf(t, want[0], got[0], "Prefix.Get(0) should equal %#X", want[0])
	assert.Equalf(t, want[1], got[1], "Prefix.Get(1) should equal %#X", want[1])
	assert.Equalf(t, want[2], got[2], "Prefix.Get(2) should equal %#X", want[2])
	assert.Equalf(t, want[3], got[3], "Prefix.Get(3) should equal %#X", want[3])

	assert.Panics(t, func() { p.Get(4) }, "Prefix.Get(4) should panic")
	assert.Panics(t, func() { p.Get(-1) }, "Prefix.Get(-1) should panic")
}

func TestSlice_R16(t *testing.T) {
	p := NewPrefix(R16, []byte{0xAB, 0xCD})
	p2 := p.Slice(1)
	p2want := []byte{0x0B, 0x0C, 0x0D}
	p2got := []byte{p2.Get(0), p2.Get(1), p2.Get(2)}
	p3 := p2.Slice(1, 2)
	p3want := []byte{0x0C}
	p3got := []byte{p3.Get(0)}
	p4 := p3.Slice(1, 1)

	assert.NotSame(t, &p, &p2, "Prefix.Slice should return a new Prefix")
	assert.Equal(t, p2, p.Slice(1, 4), "Prefix.Slice(1) should equal Prefix.Slice(1, 4)")

	assert.Equalf(t, p2want[0], p2got[0], "Prefix.Slice(1).Get(0) should equal %#X", p2want[0])
	assert.Equalf(t, p2want[1], p2got[1], "Prefix.Slice(1).Get(1) should equal %#X", p2want[1])
	assert.Equalf(t, p2want[2], p2got[2], "Prefix.Slice(1).Get(2) should equal %#X", p2want[2])
	assert.Panics(t, func() { p2.Get(3) }, "Prefix.Slice(1).Get(3) should panic")
	assert.Panics(t, func() { p2.Get(-1) }, "Prefix.Slice(1).Get(-1) should panic")

	assert.Equalf(t, p3want[0], p3got[0], "Prefix.Slice(1).Slice(1, 2).Get(0) should equal %#X", p3want[0])
	assert.Panics(t, func() { p3.Get(2) }, "Prefix.Slice(1).Slice(1, 2).Get(2) should panic")
	assert.Panics(t, func() { p3.Get(-1) }, "Prefix.Slice(1).Slice(1, 2).Get(-1) should panic")

	assert.Equalf(t, 0, p4.Len(), "Prefix.Slice(1).Slice(1, 2).Slice(1, 1).Len() should equal %#X", 0)
	assert.Panics(t, func() { p4.Get(0) }, "Prefix.Slice(1).Slice(1, 2).Slice(1, 1).Get(0) should panic")
	assert.Panics(t, func() { p4.Get(-1) }, "Prefix.Slice(1).Slice(1, 2).Slice(1, 1).Get(-1) should panic")
}

func TestLen_R16(t *testing.T) {
	p := NewPrefix(R16, []byte{0xAB, 0xCD})
	want := 4
	got := p.Len()
	p2 := p.Slice(4, 4)
	p2want := 0
	p2got := p2.Len()

	assert.Equalf(t, want, got, "Prefix.Len() should equal %#X", want)
	assert.Equalf(t, p2want, p2got, "Prefix.Slice(4, 4).Len() should equal %#X", want)
}
