package bitset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitset(t *testing.T) {
	x := New(uint64(1))                      // Init at `1`
	assert.True(t, x.IsSet(64))              // Confirm the LSB is 1
	x.SetBit(1)                              // Set the MSB
	assert.True(t, x.IsSet(1))               // Confirm the MSB is 1
	assert.True(t, x.IsSet(64))              // Confirm the LSB is still 1
	x.UnsetBit(1)                            // Unset the MSB
	assert.False(t, x.IsSet(1))              // Confirm the MSB is 0
	assert.True(t, x.IsSet(64))              // Confirm the LSB is still 1
	assert.Equal(t, uint64(0b1), x.And(0b1)) // 0b1 & 0b1 == 0b1
	x.OrAssign(0b11 << 61)                   // Set bits 0b011...
	assert.True(t, x.IsSet(2))               // Confirm 0b01... >> (64-2) == 1
	assert.True(t, x.IsSet(3))               // Confirm 0b001... >> (64-2) == 1

	// Test setting all bits for each unsigned integer type
	//
	// We also test the `Bits()` iterator here

	// uint8
	const u8max uint8 = 1<<8 - 1
	b8 := New(uint8(0))
	b8.OrAssign(u8max)
	assert.Equal(t, b8.value, u8max)
	for i, v := range b8.Bits() {
		assert.Equal(t, uint8(1), v, "expected all bits set, [%d] is not", i)
	}

	// uint16
	const u16max uint16 = 1<<16 - 1
	b16 := New(uint16(0))
	b16.OrAssign(u16max)
	assert.Equal(t, u16max, b16.value)
	for i, v := range b16.Bits() {
		assert.Equal(t, uint16(1), v, "expected all bits set, [%d] is not", i)
	}

	// uint32
	const u32max uint32 = 1<<32 - 1
	b32 := New(uint32(0))
	b32.OrAssign(u32max)
	assert.Equal(t, u32max, b32.value)
	for i, v := range b32.Bits() {
		assert.Equal(t, uint32(1), v, "expected all bits set, [%d] is not", i)
	}

	// uint64
	const u64max uint64 = 1<<64 - 1
	b64 := New(uint64(0))
	b64.OrAssign(u64max)
	assert.Equal(t, u64max, b64.value)
	for i, v := range b64.Bits() {
		assert.Equal(t, uint64(1), v, "expected all bits set, [%d] is not", i)
	}

	// Test unsetting all bits for each unsigned integer type

	// uint8
	b8.AndNot(1<<8 - 1)
}
