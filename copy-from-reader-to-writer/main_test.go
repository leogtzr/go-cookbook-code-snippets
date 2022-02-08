package main

import "testing"

// go test -bench=. -benchmem
// -bench=. run all the benchmark tests in the current package.
// -benchmem to show memory benchmarks.
func BenchmarkReadWrite(b *testing.B) {
	readWrite()
}

func BenchmarkCopy(b *testing.B) {
	copy()
}
