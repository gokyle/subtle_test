package subtletest

import (
	"crypto/subtle"
	"testing"
)

func Benchmark0(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("BCDEFGHIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark1(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ACDEFGHIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark2(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABDEFGHIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark3(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCEFGHIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark4(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDFGHIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark5(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEGHIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark6(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFHIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark7(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGIJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark8(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHJKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark9(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHIKLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark10(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHIJLMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark11(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHIJKMNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark12(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHIJKLNOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark13(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHIJKLMOP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func Benchmark14(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHIJKLMNP")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}

func BenchmarkMatch(b *testing.B) {
	b.StopTimer()
	a := []byte("ABCDEFGHIJKLMNO")
	c := []byte("ABCDEFGHIJKLMNO")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		subtle.ConstantTimeCompare(a, c)
	}
}
