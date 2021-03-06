// Public Domain (-) 2010-2011 The Ampify Authors.
// See the Ampify UNLICENSE file for details.

package argo

import (
	"amp/big"
	"bytes"
	"fmt"
	"testing"
)

func Buffer() *bytes.Buffer {
	return bytes.NewBuffer([]byte{})
}

func TestWriteSize(t *testing.T) {

	tests := map[uint64]string{
		0:                    "\x00",
		123456789:            "\x95\x9a\xef:",
		18446744073709551615: "\xff\xff\xff\xff\xff\xff\xff\xff\xff\x01",
	}

	for value, expected := range tests {
		buf := Buffer()
		WriteSize(value, buf)
		if string(buf.Bytes()) != expected {
			t.Errorf("Got unexpected encoding for %d: %q", value, buf.Bytes())
		}
	}

}

func testWriteInt(t *testing.T) {
	N := int64(8322944)
	buf := Buffer()
	WriteInt(N, buf)
	fmt.Printf("%q\n", string(buf.Bytes()))
}

func testWriteBigInt(t *testing.T) {
	N := big.NewInt(8322944)
	buf := Buffer()
	WriteBigInt(N, buf)
	fmt.Printf("%q\n", string(buf.Bytes()))
}


func testWriteIntOrdering(t *testing.T) {

	buf := Buffer()
	WriteInt(-10258176, buf)
	prev := string(buf.Bytes())

	var i int64

	for i = -10258175; i < 10258175; i++ {
		buf.Reset()
		WriteInt(i, buf)
		cur := string(buf.Bytes())
		if prev >= cur {
			t.Errorf("Lexicographical ordering failure for %d -- %q >= %q", i, prev, cur)
		}
		prev = cur
	}

}

func testWriteBigIntOrdering(t *testing.T) {

	buf := Buffer()
	WriteBigInt(big.NewInt(-10258176), buf)
	prev := string(buf.Bytes())

	var i int64

	for i = -10258175; i < 10258175; i++ {
		buf.Reset()
		WriteBigInt(big.NewInt(i), buf)
		cur := string(buf.Bytes())
		if prev >= cur {
			t.Errorf("Lexicographical ordering failure for %d -- %q >= %q", i, prev, cur)
		}
		prev = cur
	}

}

func decimal(value string) *big.Decimal {
	dec, _ := big.NewDecimal(value)
	return dec
}

func TestWriteDecimalOrdering(t *testing.T) {

	buf := Buffer()
	WriteDecimal(decimal("0"), buf)
	prev := string(buf.Bytes())

	tests := []string{
		"0.02",
		"0.0201",
		"0.05",
		"2",
		"2.30001",
		"2.30002",
	}

	for _, value := range tests {
		buf.Reset()
		WriteDecimal(decimal(value), buf)
		cur := string(buf.Bytes())
		if prev >= cur {
			left, right := decimal(value).Components()
			t.Errorf("Lexicographical ordering failure for %s (%s, %s) -- %q >= %q",
				value, left, right, prev, cur)
		}
		prev = cur
	}

}


func BenchmarkWriteSize(b *testing.B) {
	buf := Buffer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		WriteSize(123456789, buf)
	}
}

func BenchmarkWriteNumber(b *testing.B) {
	buf := Buffer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		WriteNumber("123456789", buf)
	}
}
