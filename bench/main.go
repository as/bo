package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

import "github.com/as/bo"

func main() {
	b := make([]byte, 8)

	ad := int32(0x41424344)
	eh := int32(0x45464748)

	bo.P32b(b, ad)
	bo.P32b(b[4:], eh)

	fmt.Printf("%v\n", string(b))
	// ABCDEFGH

	ah := bo.G64b(b)

	fmt.Printf("%x\n", ah)
	// 4142434445464748

	bench(10e6)
}

func bench(n int) {
	b := make([]byte, n*4)
	v := int32(0x41424344)
	sw := new(Stopwatch)

	sp := func(s string) {
		fmt.Printf("%s: %d ops took %d ms\n", s, n, sw.End())
	}

	// Benchmark bo functions

	sw.Beg()
	for i := 0; i < n; i++ {
		bo.P32b(b[4*i:], v)
	}
	sp("go: bo.P32b")

	sw.Beg()
	for i := 0; i < n; i++ {
		bo.P32l(b[4*i:], v)
	}
	sp("go: bo.P32l")

	sw.Beg()
	for i := 0; i < n; i++ {
		bo.G32b(b[4*i:])
	}
	sp("go: bo.G32b")

	sw.Beg()
	for i := 0; i < n; i++ {
		bo.G32l(b[4*i:])
	}
	sp("go: bo.G32l")

	// Benchmark encoding/binary

	buf := new(bytes.Buffer)

	sw.Beg()
	for i := 0; i < n; i++ {
		binary.Write(buf, binary.BigEndian, v)
	}
	sp("go: binary: P32b")

	sw.Beg()
	for i := 0; i < n; i++ {
		binary.Write(buf, binary.LittleEndian, v)
	}
	sp("go: binary: P32l")

	sw.Beg()
	for i := 0; i < n; i++ {
		binary.Read(buf, binary.BigEndian, v)
	}
	sp("go: binary: G32b")

	sw.Beg()
	for i := 0; i < n; i++ {
		binary.Read(buf, binary.LittleEndian, v)
	}
	sp("go: binary: G32l")

}

type Stopwatch struct {
	time time.Time
}

func (s *Stopwatch) Beg() {
	s.time = time.Now()
}

func (s *Stopwatch) End() int {
	return int(time.Now().Sub(s.time) / 10e5)
}
