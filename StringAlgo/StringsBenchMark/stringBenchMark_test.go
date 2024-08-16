package StringsBenchMark

import (
	"testing"
)

func BenchmarkPlusConcat(b *testing.B) {
	s := randomString(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		plusConcat(10000, s)
	}
}
func BenchmarkSprintfConcat(b *testing.B) {
	s := randomString(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sprintfConcat(10000, s)
	}
}

func BenchmarkBuilderConcat(b *testing.B) {
	s := randomString(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builderConcat(10000, s)
	}
}

func BenchmarkJoinConcat(b *testing.B) {
	s := randomString(10)
	strs := []string{s}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		joinConcat(10000, strs)
	}
}

func BenchmarkBufferConcat(b *testing.B) {
	s := randomString(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bufferConcat(10000, s)
	}
}
func BenchmarkByteConcat(b *testing.B) {
	s := randomString(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byteConcat(10000, s)
	}
}
func BenchmarkPreByteConcat(b *testing.B) {
	s := randomString(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		preByteConcat(10000, s)
	}
}
