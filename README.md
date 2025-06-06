# Overview

`bitset` is a basic generic bitset package, supporting all unsigned integer 
(8-64) types.

# Examples

```go
b8 := bitset.New(uint8(0))       // 0b00000000
b8.SetBit(1)                     // 0b10000000 (Set and Unset mutate)
b8.SetBit(2)                     // 0b11000000 (Set and Unset mutate)
b8.UnsetBit(2)                   // 0b10000000 (Set and Unset mutate)
fmt.Println(b8.IsSet(1))         // 0b10000000 "true"
b8.OrAssign(1 << 6)              // 0b11000000 (*Assign methods mutate)
b8.And(1<<7)                     // 0b10000000 (Does NOT mutate)
b8.Or(1<<5)                      // 0b11100000 (Does NOT mutate)
fmt.Printf("%08b\n", b8.Value()) // 0b11000000
b8.AndNotAssign(1<<8-1)          // 0b00000000 (*Assign methods mutate)
```
