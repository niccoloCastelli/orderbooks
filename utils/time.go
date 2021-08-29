package utils

import (
	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func NewTimeRange(start time.Time, end time.Time) TimeRange {
	return TimeRange{Start: start, End: end}
}

func TimeToProtoTsPtr(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{Seconds: t.Unix(), Nanos: int32(t.Nanosecond())}
}
func TimePtrToProtoTsPtr(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}
	return &timestamp.Timestamp{Seconds: t.Unix(), Nanos: int32(t.Nanosecond())}
}
func ProtoTsToTime(ts *timestamp.Timestamp) time.Time {
	return time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
}
func ProtoTsToTimePtr(ts *timestamp.Timestamp) *time.Time {
	t := time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
	return &t
}

func GogoProtoTsToTime(ts *types.Timestamp) time.Time {
	return time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
}
func GogoProtoTsToTimePtr(ts *types.Timestamp) *time.Time {
	t := time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
	return &t
}
func TimeToGogoProtoTs(ts *time.Time) types.Timestamp {
	if ts == nil {
		return types.Timestamp{Seconds: 0, Nanos: 0}
	}
	return types.Timestamp{Seconds: ts.Unix(), Nanos: int32(ts.Nanosecond())}
}
func TimeToGogoProtoTsPtr(ts *time.Time) *types.Timestamp {
	if ts == nil {
		return nil
	}
	pts := TimeToGogoProtoTs(ts)
	return &pts
}
