// Code generated by github.com/comfortablynull/rp. DO NOT EDIT.
package rp

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// For global test
	SetSeedFunc(func() int64 { return time.Now().UnixNano() })
	os.Exit(m.Run())
}

func BenchmarkRand_ExpFloat64(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).ExpFloat64()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.ExpFloat64()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.ExpFloat64()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ExpFloat64()
			}
		})
	})
}

func BenchmarkRand_Float32(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Float32()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Float32()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Float32()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Float32()
			}
		})
	})
}

func BenchmarkRand_Float64(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Float64()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Float64()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Float64()
			}
		})
	})
}

func BenchmarkRand_Int(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Int()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Int()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Int()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Int()
			}
		})
	})
}

func BenchmarkRand_Int31(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Int31()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Int31()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Int31()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Int31()
			}
		})
	})
}

func BenchmarkRand_Int31n(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100)
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Int31n(100)
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Int31n(100)
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Int31n(100)
			}
		})
	})
}

func BenchmarkRand_Int63(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Int63()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Int63()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Int63()
			}
		})
	})
}

func BenchmarkRand_Int63n(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100)
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Int63n(100)
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Int63n(100)
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Int63n(100)
			}
		})
	})
}

func BenchmarkRand_Intn(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Intn(100)
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Intn(100)
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Intn(100)
			}
		})
	})
}

func BenchmarkRand_NormFloat64(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).NormFloat64()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.NormFloat64()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.NormFloat64()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				NormFloat64()
			}
		})
	})
}

func BenchmarkRand_Perm(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Perm(100)
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Perm(100)
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Perm(100)
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Perm(100)
			}
		})
	})
}

func BenchmarkRand_Uint32(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Uint32()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Uint32()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Uint32()
			}
		})
	})
}

func BenchmarkRand_Uint64(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.Uint64()
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.Uint64()
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Uint64()
			}
		})
	})
}

func BenchmarkRand_Shuffle(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()

	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			arr := [100]int{}
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).Shuffle(100, func(i int, j int) {
					arr[i], arr[j] = arr[j], arr[i]
				})
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			arr := [100]int{}
			for pb.Next() {
				rand.Shuffle(100, func(i int, j int) {
					arr[i], arr[j] = arr[j], arr[i]
				})
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			arr := [100]int{}
			for pb.Next() {
				p.Shuffle(100, func(i int, j int) {
					arr[i], arr[j] = arr[j], arr[i]
				})
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			arr := [100]int{}
			for pb.Next() {
				Shuffle(100, func(i int, j int) {
					arr[i], arr[j] = arr[j], arr[i]
				})
			}
		})
	})
}

func BenchmarkRand_Read(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()

	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			bb := make([]byte, 1000)
			for pb.Next() {
				if _, err := rand.New(rand.NewSource(time.Now().UnixNano())).Read(bb); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			bb := make([]byte, 1000)
			for pb.Next() {
				if _, err := rand.Read(bb); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			bb := make([]byte, 1000)
			for pb.Next() {
				if _, err := p.Read(bb); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			bb := make([]byte, 1000)
			for pb.Next() {
				if _, err := Read(bb); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}
