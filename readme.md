## What
Package bo (byte order) provides functions for packing and unpacking 16, 32, and 64 bit integers 
to and from byte slices.

## How
The first 8 letters of the alphabet

```
	b := make([]byte, 8)

	ad := int32(0x41424344)
	eh := int32(0x45464748)

	bo.P32b(b, ad)
	bo.P32b(b[4:], eh)

	fmt.Printf("%v\n", string(b))
	// ABCDEFGH
```

## Why
bo was written to pack data quickly.
bo runs 40% slower than inline C but 2200% faster than go package encoding/binary.
run bench/mark.sh to benchmark on your own machine.

## Install and Benchmark
```
	git clone https://github.com/as/bo
	cd bo/bench/
	chmod u+x ./mark.sh
	./mark.sh # needs gcc and go compiler
```
