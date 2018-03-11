#!/usr/bin/env python3

"""Generate word lists for use in listings-golang.sty"""

# Text source: https://golang.org/pkg/builtin/

bif = '''
func append(slice []Type, elems ...Type) []Type
func cap(v Type) int
func close(c chan<- Type)
func complex(r, i FloatType) ComplexType
func copy(dst, src []Type) int
func delete(m map[Type]Type1, key Type)
func imag(c ComplexType) FloatType
func len(v Type) int
func make(t Type, size ...IntegerType) Type
func new(Type) *Type
func panic(v interface{})
func print(args ...Type)
func println(args ...Type)
func real(c ComplexType) FloatType
func recover() interface{}
'''

print('BUILT-IN FUNCTIONS')
for lin in bif.split('\n'):
    if lin.strip():
        print(lin.replace('func ','').split('(')[0], end=',')
print()

types = '''
type bool
type byte
type complex64
type complex128
type error
type float32
type float64
type int
type int8
type int16
type int32
type int64
type rune
type string
type uint
type uint8
type uint16
type uint32
type uint64
type uintptr
type ComplexType
type FloatType
type IntegerType
type Type
type Type1
'''

print('BUILT-IN TYPES')
for lin in types.split('\n'):
    if lin.strip():
        print(lin.replace('type ',''), end=',')
print()
