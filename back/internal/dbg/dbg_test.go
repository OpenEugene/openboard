package dbg

import (
	"os"
	"testing"
)

func BenchmarkDbgUse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		useDbgLogExample()
	}
}

func useDbgLogExample() {
	Log("\ndebug is writing to nothing")

	SetDebugOut(os.Stdout)

	Log("debug has been set to write to stdout")

	Logf("writing various formats: %t, %s, %d", true, "word", 79)

	SetDebugOut(nil)

	Log("write to nothing again")
}
