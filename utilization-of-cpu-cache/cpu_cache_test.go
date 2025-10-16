package utilization_of_cpu_cache

import "testing"

func BenchmarkCountryCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := CountryCount(users)
		if m == nil {
			b.Fatal(m)
		}
	}
}
