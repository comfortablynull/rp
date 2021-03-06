// Code generated by github.com/comfortablynull/rp. DO NOT EDIT.
package rp
import "math/rand"

type SeedFunc func() int64
type SourceFunc func() rand.Source
type Source64Func func() rand.Source64
type RandFunc func() *rand.Rand

type Rand interface {
	ExpFloat64() (of float64)
	Float32() (of float32)
	Float64() (of float64)
	Int() (oi int)
	Int31() (oi int32)
	Int31n(i int32) (oi int32)
	Int63() (oi int64)
	Int63n(i int64) (oi int64)
	Intn(i int) (oi int)
	NormFloat64() (of float64)
	Perm(i int) (oi []int)
	Read(u []uint8) (oi int, oe error)
	Shuffle(i int, f func(int, int))
	Uint32() (ou uint32)
	Uint64() (ou uint64)
}