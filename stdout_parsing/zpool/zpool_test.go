package main

import "testing"

func BenchmarkRunCommand(b *testing.B) {
	for n := 0; n < 100000; n++ {
		runCommand()
	}
}
