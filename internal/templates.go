package internal

import "text/template"

var (
	Interface = template.Must(template.New("interface").Parse(
		`
import "math/rand"

type SeedFunc func() int64
type SourceFunc func() rand.Source
type Source64Func func() rand.Source64
type RandFunc func() *rand.Rand

type Rand interface {
{{- range .}}
	{{ .Definition }}
{{- end}}
}`))

	Global = template.Must(template.New("global").Parse(
		`
var p = NewPoolSeedFunc(func() int64 { return 1 })

func SetSeedFunc(s SeedFunc)         { p.(*pool).setSeedFunc(s) }
func SetSourceFunc(s SourceFunc)     { p.(*pool).setSourceFunc(s) }
func SetSource64Func(s Source64Func) { p.(*pool).setSource64Func(s) }
func SetRandFunc(r RandFunc)         { p.(*pool).setRandFunc(r) }
{{ range .}}
func {{ .Definition }} {
{{- if .OutArgsList }}
	return p.{{ .Name }}({{ .InArgsList }})
{{- else}}
	p.{{ .Name }}({{ .InArgsList }})
{{- end }}
}
{{- end}}
`))

	Decorator = template.Must(template.New("decorator").Parse(
		`
import (
	"math/rand"
	"sync"
)

type pool struct {
	pool *sync.Pool
}

func newPool() *pool { return &pool{pool: &sync.Pool{}} }
func NewPoolSeedFunc(s SeedFunc) Rand {
	p := newPool()
	p.setSeedFunc(s)
	return p
}

func NewPoolSourceFunc(s SourceFunc) Rand {
	p := newPool()
	p.setSourceFunc(s)
	return p
}

func NewPoolSource64Func(s Source64Func) Rand {
	p := newPool()
	p.setSource64Func(s)
	return p
}

func NewPoolRandFunc(r RandFunc) Rand {
	p := newPool()
	p.setRandFunc(r)
	return p
}

func (p *pool) setSeedFunc(s SeedFunc) {
	p.pool.New = func() interface{} {
		return rand.New(rand.NewSource(s()))
	}
}

func (p *pool) setSourceFunc(s SourceFunc) {
	p.pool.New = func() interface{} { return rand.New(s()) }
}

func (p *pool) setSource64Func(s Source64Func) {
	p.pool.New = func() interface{} {
		return rand.New(s())
	}
}

func (p *pool) setRandFunc(r RandFunc) {
	p.pool.New = func() interface{} { return r() }
}
{{ range .}}
func (p *pool) {{ .Definition }} {
	r := p.pool.Get().(*rand.Rand)
{{- if .OutArgsList }}
	{{ .OutArgsList }} = r.{{ .Name }}({{ .InArgsList }})
{{- else}}
	r.{{ .Name }}({{ .InArgsList }})
{{- end }}
	p.pool.Put(r)
{{- if .OutArgsList }}
	return
{{- end }}
}
{{ end }}
`))
	Benchmark = template.Must(template.New("bench").Parse(`
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
{{ range .}}
func BenchmarkRand_{{ .Method.Name }}(b *testing.B) {
	p := NewPoolSeedFunc(func() int64 { return time.Now().UnixNano() })
	b.ResetTimer()
	b.Run("vanilla", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.New(rand.NewSource(time.Now().UnixNano())).{{ .Method.Name }}({{ .Args }})
			}
		})
	})
	b.Run("vanilla global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				rand.{{ .Method.Name }}({{ .Args }})
			}
		})
	})
	b.Run("rp", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				p.{{ .Method.Name }}({{ .Args }})
			}
		})
	})
	b.Run("rp global", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				{{ .Method.Name }}({{ .Args }})
			}
		})
	})
}
{{ end }}
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
}`))
)
