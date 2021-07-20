package dbg

import (
	"io/ioutil"
	"testing"
)

func BenchmarkDbgUse(b *testing.B) {
	SetDebugOut(ioutil.Discard)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Log("")
		}
	})
}
func BenchmarkDbgUseNil(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Log("")
		}
	})
}
func BenchmarkDbgSetAndUse(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SetDebugOut(ioutil.Discard)
			Log("")
		}
	})
}
func BenchmarkDbgSetAndUseNil(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SetDebugOut(nil)
			Log("")
		}
	})
}
