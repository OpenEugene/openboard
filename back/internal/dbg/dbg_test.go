package dbg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func BenchmarkDbgUse(b *testing.B) {
	SetOut(ioutil.Discard)
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
			SetOut(ioutil.Discard)
			Log("")
		}
	})
}
func BenchmarkDbgSetAndUseNil(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SetOut(nil)
			Log("")
		}
	})
}

func TestSetOut(t *testing.T) {
	Log("debug not set")

	buff := bytes.NewBuffer([]byte{})
	SetOut(buff)
	msg := "debug set to bytes buffer"
	Log(msg)
	got := buff.String()

	want := msg + "\n"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}

	SetOut(nil)
	buff.Reset()
	want = ""
	Log("debug set to nil")
	got = buff.String()
	if got != want {
		t.Errorf("want: nothing, got: %s", got)
	}

	SetOut(buff)
	buff.Reset()
	msg = "debug set to bytes buffer, %s time"
	Logf(msg, "second")
	got = buff.String()
	want = fmt.Sprintf(msg+"\n", "second")
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
