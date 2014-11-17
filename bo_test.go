package bo
// Unit tests for the bo package. These functions are a means
// to an end and do not have internal error checks. If they
// are not called correctly they will generate a runtime panic
// during go test

import (
	"reflect"
	"testing"
)

type Tests struct {
	id     string      // test id
	f, g   interface{} // fn and inverse fn
	in, ex interface{} // input & expected val
}

func TestP(t *testing.T) {

	// Test table
	tt := []Tests{
		{"16l", P16l, G16l, int16(0x4142), "BA"},
		{"16b", P16b, G16b, int16(0x4142), "AB"},
		{"32l", P32l, G32l, int32(0x41424344), "DCBA"},
		{"32b", P32b, G32b, int32(0x41424344), "ABCD"},
		{"64l", P64l, G64l, int64(0x4142434445464748), "HGFEDCBA"},
		{"64b", P64b, G64b, int64(0x4142434445464748), "ABCDEFGH"},
	}

	// Test all functions and inverse functions
	for _, v := range tt {
		PTest(t, v.id, v.f, v.in, v.ex)
		GTest(t, v.id, v.g, v.ex, v.in) // swap ex and in
	}

}

// PTest tests bo functions that start with the letter
// P. fu is the function to call, in is the int16, int32
// or int64 input, and ex is a string representing the
// expected value. If fu does not yield the string-equivalent
// ex value the test fails.
func PTest(t *testing.T, id string, fu, in, ex interface{}) {
	fun := reflect.ValueOf(fu)
	inp := reflect.ValueOf(in)

	tmp := make([]byte, sizeof(in))
	buf := reflect.ValueOf(tmp)

	args := []reflect.Value{buf, inp}
	fun.Call(args)

	ac := string(buf.Interface().([]byte))
	if ac != ex.(string) {
		t.Logf("P%s: ac (%v) != ex (%v)\n", id, ac, ex)
		t.Fail()
	}
}

// GTest is similar to PTest. Except in and ex are transposed.
// in is a string and ex is an int16, int32, or int64 value.
func GTest(t *testing.T, id string, fu, in, ex interface{}) {
	fun := reflect.ValueOf(fu)
	inp := reflect.ValueOf([]byte(in.(string)))

	args := []reflect.Value{inp}
	ret := fun.Call(args)
	ac := ret[0].Interface()

	if ac != ex {
		t.Logf("G%s: ac (%v) != ex (%v)\n", id, ac, ex)
		t.Fail()
	}
}

// sizeof returns the size of an integer in bytes
func sizeof(i interface{}) int {
	return reflect.TypeOf(i).Bits() / 8
}
