package bcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodePrimitives(t *testing.T) {
	// bool
	assert.Equal(t, []byte{1}, MustEncodeToBytes(true))
	assert.Equal(t, []byte{0}, MustEncodeToBytes(false))

	// int8 & uint8
	assert.Equal(t, []byte{0xFF}, MustEncodeToBytes(int8(-1)))
	assert.Equal(t, []byte{1}, MustEncodeToBytes(uint8(1)))

	// int16 & uint16
	assert.Equal(t, []byte{0xCC, 0xED}, MustEncodeToBytes(int16(-4660)))
	assert.Equal(t, []byte{0x34, 0x12}, MustEncodeToBytes(uint16(4660)))

	// int32 & uint32
	assert.Equal(t, []byte{0x88, 0xA9, 0xCB, 0xED}, MustEncodeToBytes(int32(-305419896)))
	assert.Equal(t, []byte{0x78, 0x56, 0x34, 0x12}, MustEncodeToBytes(uint32(305419896)))

	// int64 & uint64
	assert.Equal(t, []byte{0x00, 0x11, 0x32, 0x54, 0x87, 0xA9, 0xCB, 0xED}, MustEncodeToBytes(int64(-1311768467750121216)))
	assert.Equal(t, []byte{0x00, 0xEF, 0xCD, 0xAB, 0x78, 0x56, 0x34, 0x12}, MustEncodeToBytes(uint64(1311768467750121216)))
}

func TestEncodeOption(t *testing.T) {
	assert.Equal(t, []byte{0}, MustEncodeToBytes((*uint8)(nil)))

	val := uint8(8)
	assert.Equal(t, []byte{1, 8}, MustEncodeToBytes(&val))
}

func TestEncodeArray(t *testing.T) {
	val := [3]uint16{1, 2, 3}
	assert.Equal(t, []byte{1, 0, 2, 0, 3, 0}, MustEncodeToBytes(val))
}

func TestEncodeSlice(t *testing.T) {
	val1 := []uint16{1, 2}
	assert.Equal(t, []byte{2, 1, 0, 2, 0}, MustEncodeToBytes(val1))

	val2 := make([]uint8, 9487)
	encoded := MustEncodeToBytes(val2)
	assert.Equal(t, 9487+2, len(encoded))
	assert.Equal(t, []byte{0x8f, 0x4a}, encoded[:2])
}

func TestEncodeString(t *testing.T) {
	val := "çå∞≠¢õß∂ƒ∫"
	expecting := []byte{
		24, 0xc3, 0xa7, 0xc3, 0xa5, 0xe2, 0x88, 0x9e, 0xe2, 0x89, 0xa0, 0xc2,
		0xa2, 0xc3, 0xb5, 0xc3, 0x9f, 0xe2, 0x88, 0x82, 0xc6, 0x92, 0xe2, 0x88, 0xab,
	}
	assert.Equal(t, expecting, MustEncodeToBytes(val))
}

func TestEncodeBytes(t *testing.T) {
	val := []byte{1, 2, 3}
	assert.Equal(t, []byte{3, 1, 2, 3}, MustEncodeToBytes(val))
}

func TestEncodeStruct(t *testing.T) {
	type Foo struct {
		Bool bool
	}

	type Bar struct {
		Foo
		Bytes []byte
		Label string
	}

	type Zoo struct {
		Inner Bar
		Name  string
	}

	bar := Bar{
		Foo:   Foo{true},
		Bytes: []byte{0xC0, 0xDE},
		Label: "a",
	}
	assert.Equal(t, []byte{1, 2, 0xC0, 0xDE, 1, 'a'}, MustEncodeToBytes(bar))

	zoo := Zoo{
		Inner: bar,
		Name:  "b",
	}
	assert.Equal(t, []byte{1, 2, 0xC0, 0xDE, 1, 'a', 1, 'b'}, MustEncodeToBytes(zoo))
}
