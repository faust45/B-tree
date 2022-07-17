package main

import (
	// "bytes"
	"fmt"
	// "sort"
	// "syscall"
	// "github.com/edsrzf/mmap-go"
	// "golang.org/x/sys/unix"
	// "os"
)

func main() {
	fmt.Printf("Start\n")
	// keys := [][]byte{[]byte("xxx"), []byte("11x")}
	key := Int64ToBytes(150)
	value := []byte("alegro")
	k1 := Int64ToBytes(500)
	v1 := []byte("alegro5")
	k2 := Int64ToBytes(730)
	v2 := []byte("alegro730")
	k3 := Int64ToBytes(830)
	v3 := []byte("alegro830")

	// k4 := Int64ToBytes(151)
	// v4 := []byte("alegro830")
	// k5 := Int64ToBytes(155)
	// v5 := []byte("alegro830")
	// k6 := Int64ToBytes(157)
	// v6 := []byte("alegro830")

	Put(key, value)
	Put(k1, v1)
	Put(k2, v2)
	Put(k3, v3)
	// Put(k4, v4)
	// Put(k5, v5)
	// Put(k6, v6)
	PrintTree()

	// f, _ := os.OpenFile("./file", os.O_RDWR, 0644)
	// defer f.Close()
	// // f.Write(make([]byte, 100))
	// // f.Sync()

	// sz := 100000
	// mmap, err := syscall.Mmap(int(f.Fd()), 0, sz, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	// if err != nil {
	// 	fmt.Printf("madvise: %s", err)
	// }

	// // Advise the kernel that the mmap is accessed randomly.
	// err = syscall.Madvise(mmap, syscall.MADV_RANDOM)
	// if err != nil && err != syscall.ENOSYS {
	// 	// Ignore not implemented error in kernel because it still works.
	// 	fmt.Printf("madvise: %s", err)
	// }

	// fmt.Printf("%d\n", len(mmap))
	// mmap[200] = 'X'
	// f.Sync()
}
