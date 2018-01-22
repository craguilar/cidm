package services

import (
	"testing"
)

//https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
// Benchmark function for Failure
func BenchmarkHandleBearerAuthS(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		HandleBearerAuth("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNydWl6QGludGVsZW56LmNvbSIsImV4cCI6MTUxMTM3NzY4MCwiaWF0IjoxNTA4Njk1NjgwLCJpc3MiOiJzb3VyY2UuaW50ZWxlbnouY29tIiwic3ViIjoiand0LmludGVsZW56LmNvbSJ9.rpEFTp3aV4HwXFsi73R999prVr50RYpfAB6TQ_3Wyo4")
	}
}

// Benchmark function for Failure
func BenchmarkHandleBearerAuthF(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		HandleBearerAuth("Bearer test")
	}
}
