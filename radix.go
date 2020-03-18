package trie

import (
	// "fmt"
	"github.com/k0kubun/pp"
	"github.com/arccoza/go-trie/generic"
)

type Item = generic.Type

type ItemTrie struct{
	root node
}

type node struct {
	prefix []byte
	hmask byte
	tmask byte
	leaf bool
	value Item
	MergeOp func(Item, Item) Item
	edges edges
}

type edges []*node

func NewItemTrie() *ItemTrie {
	return &ItemTrie{root: node{edges: make(edges, 16)}}
}

func (t *ItemTrie) Get(key []byte) *node {
	return t.root.get(key, 0xF0)
}

func (t *ItemTrie) Put(key []byte, value Item) {
	// Make sure to copy the key so there is no ref to it externally,
	// preventing unseen modification.
	key = append([]byte(nil), key...)
	t.root.put(&node{prefix: key, hmask: 0xF0, tmask: 0x0F, leaf: true, value: value})
}

func (t *ItemTrie) Del(key []byte) {
	// return t.root.del(key, 0xF0)
}

func (n *node) IsLeaf() bool {
	return n.leaf
}

func (e edges) add(n *node) {
	idx := getNibble(n.prefix[0], n.hmask)
	e[idx] = n
}

func (head *node) split(brk int, msk byte, off int) (tail *node) {
	if brk < len(head.prefix) {
		tail = &node{
			prefix: head.prefix[brk + off:],
			hmask: ^msk,
			tmask: head.tmask,
			leaf: head.leaf,
			value: head.value,
			edges: head.edges,
		}
		head.prefix = head.prefix[:brk + off + 1]
		head.edges = make(edges, 16)
		head.tmask, head.leaf, head.value = msk, false, 0
	}

	return tail
}

func (n *node) chop(brk int, msk byte, off int) bool {
	if brk < len(n.prefix) {
		n.prefix = n.prefix[brk + off:]
		n.hmask = ^msk
		return true
	}
	return false
}

func (prv *node) put(n *node) {
	pp.Println("PUT")
	for frm, brk, msk, off := walk(prv, n.prefix, n.hmask); brk >= 0 ; frm, brk, msk, off = walk(frm, n.prefix, n.hmask) {
		// SPLIT EXISTING
		if tail := frm.split(brk, msk, off); tail != nil {
			// pp.Println("SPLIT", tail)
			frm.edges.add(tail)
		}
		// MATCH
		if !n.chop(brk, msk, off) {
			// pp.Println("MATCH", n)
			frm.leaf = true
			frm.value = n.value // TODO: Integrate MergeOp fn
			return
		// NO MATCH
		} else {
			// pp.Println("NO MATCH")
			prv = frm
		}
	}

	if prv.edges == nil {
		prv.edges = make(edges, 16)
	}
	prv.edges.add(n)
}

func (rt *node) get(key []byte, nib byte) *node {
	for frm, brk, msk, off := walk(rt, key, nib); brk >= 0 ; frm, brk, msk, off = walk(frm, key, msk) {
		// SHORT
		if len(key) < len(frm.prefix) {
			return frm
		// BRANCH
		} else if brk < len(key) {
			key = key[brk + off:]
			msk = ^msk
		// MATCH
		} else {
			return frm
		}
	}

	return nil
}

func walk(from *node, key []byte, nib byte) (*node, int, byte, int) {
	if len(key) == 0 { return nil, -1, nib, 0 }
	idx := getNibble(key[0], nib)
	n := from.edges[idx]
	if n == nil { return nil, -1, nib, 0 }
	// if len(n.prefix) == 0 { return nil, -1, nib, 0 }

	brk, msk, off := 0, byte(0x0F), -1
	for minLen := minInt(len(n.prefix), len(key)); brk < minLen; brk++ {
		if a, b := n.prefix[brk], key[brk]; a != b {
			if (a ^ b) & 0xF0 == 0 {
				msk = 0xF0
				off = 0
			} else {
				off = 1
				brk--
			}
			break
		}
	}

	// If the entire prefix matched, do a tail check to ensure the last element
	// isn't the first nibble of the byte, otherwise we need to break there.
	// You don't need to do this sort of check on the head, which nibble of the byte
	// the head actually starts on because of the nature of prefix tries,
	// the prefixes have to match (ie. the earlier nibbles have to match).
	if brk == len(n.prefix) && n.tmask == 0xF0 {
		brk, msk, off = brk - 1, 0xF0, 0
	}

	return n, brk, msk, off
}

func getNibble(b, m byte) byte {
	n := b & m
	if n > 0x0F { n = n >> 4 }
	return n
}

func minInt(a, b int) int {
	if a < b { return a }
	return b
}
