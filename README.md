## Benchmarks
```bash
go test -run=. -bench=. -benchtime=20s -benchmem
```

```
goos: windows
goarch: amd64
pkg: github.com/mikkael131/go-leet-discount-distribution
cpu: 12th Gen Intel(R) Core(TM) i7-12700K
Benchmark_applyDiscountWithDonation/find_allocation_with_the_least_donation-20          1000000000              21.27 ns/op            0 B/op          0 allocs/op
Benchmark_applyDiscountWithDonation/exceeding_discount-20                               1000000000              21.21 ns/op            0 B/op          0 allocs/op
Benchmark_applyDiscountWithDonation/find_the_no_donation_distribution-20                774245772               31.01 ns/op            0 B/op          0 allocs/op
Benchmark_applyDiscountWithDonation/prioritize_lowest_overall_discount_percentage-20    1000000000              21.23 ns/op            0 B/op          0 allocs/op
```