package trie

import (
	"testing"
	"github.com/k0kubun/pp"
	// "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	// os.Exit(m.Run())

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
	p := NewPrefix(R16, []byte{0xAB, 0xCD})
	// p := Prefix{exp: 4, key: []byte{0xAB, 0xCD}}
	pp.Println(p, R2, R4, R16, R256)
	pp.Println(p.Get(0), p.Get(1), p.Len())
}
