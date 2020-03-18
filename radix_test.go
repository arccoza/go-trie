package trie

import (
	"testing"
	"github.com/k0kubun/pp"
	// "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	// os.Exit(m.Run())

	t := NewItemTrie()
	t.Put([]byte("remember"), 0x0001)
	t.Put([]byte("remain"), 0x0002)
	t.Put([]byte("rem"), 0x0004)

	pp.Println(t)
}
