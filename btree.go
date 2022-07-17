package main

import (
	"bytes"
	"fmt"
	// "sort"
	"encoding/binary"
)

var (
	maxNodeSize = 3
	splitAt     = 2
	keySize     = 8
	root        = &node{}
)

type Keys []Key
type Key []byte

type node struct {
	isLeaf     bool
	isOverflow bool
	keys       Keys
	children   []*node
	values     [][]byte
}

func init() {
	root.keys = Keys{Int64ToBytes(100)}
	n := leaf(Int64ToBytes(10), []byte("astruda"))
	n1 := leaf(Int64ToBytes(700), []byte("alex"))
	root.children = []*node{n, n1}
}

func PrintTree() {
	root.print()
}

func Put(key, value []byte) {
	path := root.search(key)

	leafNode := path[0]
	leafNode.add(key, value)

	for _, n := range path {
		if n.isOverflow {
			fmt.Println("overflow")
			if n == root {
				fmt.Println("split root")
				a := make([]Key, splitAt-1)
				b := make([]Key, maxNodeSize-splitAt)
				ach := make([]*node, splitAt)
				bch := make([]*node, splitAt)

				copy(a, n.keys[:splitAt-1])
				copy(b, n.keys[splitAt:])

				copy(ach, n.children[:splitAt])
				copy(bch, n.children[splitAt:])

				n.keys = Keys{n.keys[splitAt-1]}
				n1 := &node{keys: a, children: ach}
				n2 := &node{keys: b, children: bch}
				n.children = []*node{n1, n2}
			} else {
				a := make([]Key, splitAt-1)
				b := make([]Key, maxNodeSize-splitAt)

				copy(a, n.keys[:splitAt-1])
				copy(b, n.keys[splitAt-1:])

				n.keys = a
				newNode := &node{isLeaf: n.isLeaf, keys: b}
				path[1].promote(newNode)
			}

			n.isOverflow = false
		} else {
			break
		}

	}
}

func (n *node) promote(node *node) error {
	if n.isLeaf {
		return fmt.Errorf("promote works only for non leavs nodes")
	}

	defer n.checkIsOverflow()
	key := node.keys[0]
	for i, k := range n.keys {
		if bytes.Compare(key, k) < 0 {
			n.keys = sliceInsert(n.keys, i, key)
			n.children = sliceInsertCh(n.children, i+1, node)

			return nil
		}
	}

	n.keys = append(n.keys, key)
	n.children = append(n.children, node)

	return nil
}

func (n *node) checkIsOverflow() {
	if len(n.keys) == maxNodeSize {
		n.isOverflow = true
	}
}

func (n *node) add(key, value []byte) error {
	if !n.isLeaf {
		return fmt.Errorf("add works only for leavs nodes")
	}

	defer n.checkIsOverflow()

	for i, k := range n.keys {
		if bytes.Compare(key, k) < 0 {
			n.keys = sliceInsert(n.keys, i, key)
			return nil
		}
	}

	n.keys = append(n.keys, key)
	return nil
}

func (n *node) search(key []byte) []*node {
	if n.isLeaf {
		return []*node{n}
	}

	countKeys := len(n.keys)
	if 0 < countKeys {
		ch := n.children[countKeys]
		for i := 0; i < countKeys; i++ {
			if bytes.Compare(key, n.keys[i]) < 0 {
				ch = n.children[i]
				break
			}
		}

		path := ch.search(key)
		return append(path, n)
	}

	return nil
}

func (n *node) printNode() {
	if n.isLeaf {
		fmt.Printf("leaf -- ")
	} else {
		fmt.Printf("node -- ")
	}

	for _, k := range n.keys {
		fmt.Printf("%d, ", BytesToInt64(k))
	}
	fmt.Print("\n")
}

func (n *node) print() {
	if n.isLeaf {
		fmt.Printf("leaf -- ")
	}

	for _, k := range n.keys {
		fmt.Printf("%d, ", BytesToInt64(k))
	}

	fmt.Printf("\n")
	for _, ch := range n.children {
		ch.print()
	}

}

func leaf(key, value []byte) *node {
	n := node{isLeaf: true, keys: Keys{key}, values: [][]byte{value}}
	return &n
}

func Int64ToBytes(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
	}
	return buff.Bytes()
}

func BytesToInt64(b []byte) int64 {
	var n int64

	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &n)
	return n
}

func sliceInsert(coll Keys, i int, key Key) Keys {
	if len(coll) <= i {
		return append(coll, key)
	}
	coll = append(coll[:i+1], coll[i:]...)
	coll[i] = key
	return coll
}

func sliceInsertCh(coll []*node, i int, key *node) []*node {
	if len(coll) == i {
		return append(coll, key)
	}

	coll = append(coll[:i+1], coll[i:]...)
	coll[i] = key
	return coll
}
