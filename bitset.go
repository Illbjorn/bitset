package bitset

import (
	"iter"
	"log"
	"unsafe"
)

func New[T unsigned](value T) Bitset[T] {
	bits := uint8(unsafe.Sizeof(value) * 8)
	return Bitset[T]{
		value: value,
		bits:  bits,
	}
}

type (
	unsigned interface {
		~uint8 | ~uint16 | ~uint32 | ~uint64
	}

	Bitset[T unsigned] struct {
		value T
		bits  uint8
	}

	Bitset8  = Bitset[uint8]
	Bitset16 = Bitset[uint16]
	Bitset32 = Bitset[uint32]
	Bitset64 = Bitset[uint64]
)

/*                                  Read-only                                 */

func (self Bitset[T]) Value() T {
	return self.value
}

func (self Bitset[T]) And(other T) T {
	return self.value & other
}

func (self Bitset[T]) AndNot(other T) T {
	return self.value &^ other
}

func (self Bitset[T]) Or(other T) T {
	return self.value | other
}

func (self Bitset[T]) OrNot(other T) T {
	return self.value | ^other
}

func (self Bitset[T]) Xor(other T) T {
	return self.value ^ other
}

// IsSet reports whether the ith (`bit`) bit is set.
//
// The bit position is 1-based, with 1 being the most significant bit and 64
// the least significant.
//
// NOTE: Values < 1 or > 64 are ignored!
func (self Bitset[T]) IsSet(bit uint8) bool {
	return self.value>>(self.bits-bit)&0b1 == 1
}

// Bits returns an iterator over each bit of the underlying unsigned integer
// type. Returned values are the 1-based bit position (from MSB->LSB)
func (self Bitset[T]) Bits() iter.Seq2[uint8, T] {
	return func(yield func(uint8, T) bool) {
		for i := range self.bits {
			i := i + 1
			bit := self.bits - i
			if !yield(i, self.value>>bit&0b1) {
				return
			}
		}
	}
}

/*                                 Mutative                                 */

// SetBit sets the ith (`bit`) bit.
//
// The bit position is 1-based, with 1 being the most significant bit and 64
// the least significant.
//
// NOTE: Values < 1 or > 64 are ignored!
func (self *Bitset[T]) SetBit(bit uint8) {
	mask := T(1) << (self.bits - bit)
	self.value |= mask
}

// UnsetBit unsets (zeroes) the ith (`bit`) bit.
//
// The bit position is 1-based, with 1 being the most significant bit and 64
// the least significant.
//
// NOTE: Values < 1 or > 64 are ignored!
func (self *Bitset[T]) UnsetBit(bit uint8) {
	mask := T(1) << (self.bits - bit)
	self.value &^= mask
}

// AndAssign performs a bitwise-and of `self` and `other`, storing the result
// to `self`.
func (self *Bitset[T]) AndAssign(other T) {
	log.Printf("got [%08b]", other)
	self.value &= other
}

// AndAssign performs a bitwise-and-not of `self` and `other`, storing the
// result to `self`.
func (self *Bitset[T]) AndNotAssign(other T) {
	self.value &^= other
}

func (self *Bitset[T]) OrAssign(other T) {
	self.value |= other
}
