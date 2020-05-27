package main

import "testing"

func BenchmarkModuloA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		modulo(13010818623874)
	}
}

func BenchmarkModuloB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		moduloBitwise(13010818623874)
	}
}
func BenchmarkModuloC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		modulo(505679007744)
	}
}

func BenchmarkModuloD(b *testing.B) {
	for n := 0; n < b.N; n++ {
		moduloBitwise(505679007744)
	}
}
