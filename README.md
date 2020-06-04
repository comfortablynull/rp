# Rand Pool
TODO:
- Readme
- Tests

```
goos: linux
goarch: amd64
pkg: github.com/comfortablynull/rp
BenchmarkRand_ExpFloat64/vanilla-16              1326192               917 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_ExpFloat64/vanilla_global-16      20220693                59.7 ns/op             0 B/op          0 allocs/op
BenchmarkRand_ExpFloat64/rp-16                  544495420                2.28 ns/op            0 B/op          0 allocs/op
BenchmarkRand_ExpFloat64/rp_global-16           544715912                2.30 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Float32/vanilla-16                 1255000               955 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Float32/vanilla_global-16         20781561                56.9 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Float32/rp-16                     559278261                2.12 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Float32/rp_global-16              527536441                2.27 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Float64/vanilla-16                 1225156               983 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Float64/vanilla_global-16         21515570                55.0 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Float64/rp-16                     617679745                1.92 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Float64/rp_global-16              622484982                1.96 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int/vanilla-16                     1000000              1016 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Int/vanilla_global-16             22988792                53.1 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Int/rp-16                         630160885                1.91 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int/rp_global-16                  628468694                1.90 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int31/vanilla-16                   1000000              1026 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Int31/vanilla_global-16           23241214                52.3 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Int31/rp-16                       607221748                1.99 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int31/rp_global-16                609200158                1.93 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int31n/vanilla-16                  1000000              1031 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Int31n/vanilla_global-16          19454152                61.4 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Int31n/rp-16                      513438404                2.30 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int31n/rp_global-16               523482718                2.29 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int63/vanilla-16                   1000000              1027 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Int63/vanilla_global-16           23048594                52.6 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Int63/rp-16                       572844267                2.01 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int63/rp_global-16                601102017                1.98 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int63n/vanilla-16                  1000000              1081 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Int63n/vanilla_global-16          18096337                66.5 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Int63n/rp-16                      394916085                3.09 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Int63n/rp_global-16               412278001                3.02 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Intn/vanilla-16                    1000000              1076 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Intn/vanilla_global-16            19249204                64.2 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Intn/rp-16                        448952742                2.55 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Intn/rp_global-16                 485694772                2.54 ns/op            0 B/op          0 allocs/op
BenchmarkRand_NormFloat64/vanilla-16             1000000              1079 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_NormFloat64/vanilla_global-16     19454389                61.5 ns/op             0 B/op          0 allocs/op
BenchmarkRand_NormFloat64/rp-16                 509991326                2.37 ns/op            0 B/op          0 allocs/op
BenchmarkRand_NormFloat64/rp_global-16          528058783                2.27 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Perm/vanilla-16                     960505              1324 ns/op            6272 B/op          2 allocs/op
BenchmarkRand_Perm/vanilla_global-16              175233              6695 ns/op             896 B/op          1 allocs/op
BenchmarkRand_Perm/rp-16                         5378370               223 ns/op             896 B/op          1 allocs/op
BenchmarkRand_Perm/rp_global-16                  5218536               222 ns/op             896 B/op          1 allocs/op
BenchmarkRand_Uint32/vanilla-16                  1000000              1084 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Uint32/vanilla_global-16          23022897                52.3 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Uint32/rp-16                      589479589                2.02 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Uint32/rp_global-16               622832031                1.96 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Uint64/vanilla-16                  1094836              1088 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Uint64/vanilla_global-16          22298776                55.2 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Uint64/rp-16                      551955560                2.05 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Uint64/rp_global-16               588850101                2.01 ns/op            0 B/op          0 allocs/op
BenchmarkRand_Shuffle/vanilla-16                  954602              1100 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Shuffle/vanilla_global-16           195165              5977 ns/op               0 B/op          0 allocs/op
BenchmarkRand_Shuffle/rp-16                     14226909                86.0 ns/op            16 B/op          1 allocs/op
BenchmarkRand_Shuffle/rp_global-16              14538800                81.0 ns/op            16 B/op          1 allocs/op
BenchmarkRand_Read/vanilla-16                     901999              1139 ns/op            5376 B/op          1 allocs/op
BenchmarkRand_Read/vanilla_global-16             1000000              1765 ns/op               0 B/op          0 allocs/op
BenchmarkRand_Read/rp-16                        17296686                65.2 ns/op             0 B/op          0 allocs/op
BenchmarkRand_Read/rp_global-16                 18179086                65.7 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/comfortablynull/rp   82.082s
```