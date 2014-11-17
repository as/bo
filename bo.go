package bo
// Package bo (byte order) provides functions for 
// packing and unpacking 16, 32, and 64 bit integers 
// to and from byte slices.

// Function naming convention
// P = Pack, G = Get (unpack)
// NN = size of variable in bits
// b = big endian, l = little endian

// P16l packs 16 bits into b via little endian
func P16l(b []byte, v int16) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)

}

// P16b packs 16 bits into b via big endian
func P16b(b []byte, v int16) {
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}

func P32l(b []byte, v int32) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func P32b(b []byte, v int32) {
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}

func P64l(b []byte, v int64) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

func P64b(b []byte, v int64) {
	b[0] = byte(v >> 56) 
	b[1] = byte(v >> 48) 
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32) 
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
}

// G16l unpacks 16 bits from b via little endian
func G16l(b []byte) int16 {
		return int16(b[0]) |
               int16(b[1]) << 8  
}

// G16b unpacks 16 bits from b via big endian
func G16b(b []byte) int16{
	return int16(b[0]) <<  8 |
           int16(b[1]) 
}

func G32l(b []byte) int32 {
    return int32(b[0])       | 
           int32(b[1]) <<  8 | 
           int32(b[2]) << 16 | 
           int32(b[3]) << 24
}
 
func G32b(b []byte) int32 { 
	return int32(b[0]) << 24 | 
           int32(b[1]) << 16 | 
           int32(b[2]) <<  8 | 
           int32(b[3])
}
 
func G64l(b []byte) int64 { 
	return int64(b[0])       | 
           int64(b[1]) <<  8 | 
           int64(b[2]) << 16 | 
           int64(b[3]) << 24 | 
           int64(b[4]) << 32 | 
           int64(b[5]) << 40 | 
           int64(b[6]) << 48 | 
           int64(b[7]) << 56 
}

func G64b(b []byte) int64 {
	return int64(b[0]) << 56 | 
           int64(b[1]) << 48 | 
           int64(b[2]) << 40 | 
           int64(b[3]) << 32 | 
           int64(b[4]) << 24 | 
           int64(b[5]) << 16 | 
           int64(b[6]) <<  8 | 
           int64(b[7])
}
