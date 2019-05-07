package altr

import (
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// CSVFromStrings ...
func CSVFromStrings(ss []string) string {
	return strings.Join(ss, ",")
}

// LimitUint32 ...
func LimitUint32(n uint32) uint32 {
	if n == 0 {
		return 1<<32 - 1
	}
	return n
}

// Timestamp ...
func Timestamp(t time.Time, valid bool) *timestamp.Timestamp {
	var ts *timestamp.Timestamp
	if valid {
		ts, _ = ptypes.TimestampProto(t)
	}
	return ts
}
