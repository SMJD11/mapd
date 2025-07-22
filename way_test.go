package main

import (
	"testing"

	"capnproto.org/go/capnp/v3"
	"github.com/bradleyjkemp/cupaloy"
)

func TestGetNextStopSign(t *testing.T) {
	// Define a struct to hold all results for a single snapshot
	type testResult struct {
		Found     bool
		Latitude  float64
		Longitude float64
		Error     error
	}

	// --- Common Test Data Setup ---
	arena := capnp.MultiSegment([][]byte{})
	_, seg, _ := capnp.NewMessage(arena)

	// Way 1 (Current Way)
	way1, _ := NewWay(seg)
	way1.SetName("Main St")
	nodes1, _ := way1.NewNodes(4)
	nodes1.At(0).SetLatitude(0)
	nodes1.At(0).SetLongitude(0)
	nodes1.At(1).SetLatitude(0)
	nodes1.At(1).SetLongitude(1)
	nodes1.At(2).SetLatitude(0)
	nodes1.At(2).SetLongitude(2)
	nodes1.At(3).SetLatitude(0)
	nodes1.At(3).SetLongitude(3)
	stopNodes1, _ := way1.NewStopNodes(1)
	stopNodes1.At(0).SetLatitude(0)
	stopNodes1.At(0).SetLongitude(2)
	stopNodes1.At(0).SetDirection("forward")

	// Way 2 (Next Way)
	way2, _ := NewWay(seg)
	way2.SetName("Main St")
	nodes2, _ := way2.NewNodes(3)
	nodes2.At(0).SetLatitude(0)
	nodes2.At(0).SetLongitude(3)
	nodes2.At(1).SetLatitude(0)
	nodes2.At(1).SetLongitude(4)
	nodes2.At(2).SetLatitude(0)
	nodes2.At(2).SetLongitude(5)
	stopNodes2, _ := way2.NewStopNodes(1)
	stopNodes2.At(0).SetLatitude(0)
	stopNodes2.At(0).SetLongitude(5)

	// Way 3 (Another Next Way)
	way3, _ := NewWay(seg)
	way3.SetName("Side St")
	nodes3, _ := way3.NewNodes(3)
	nodes3.At(0).SetLatitude(0)
	nodes3.At(0).SetLongitude(5)

	// --- Test Scenarios ---
	t.Run("Stop sign is ahead on the current way", func(t *testing.T) {
		currentWay := CurrentWay{
			Way: way1,
			Distance: DistanceResult{
				LineStart: nodes1.At(0),
				LineEnd:   nodes1.At(1),
			},
		}
		nextWays := []NextWayResult{{Way: way2, IsForward: true}}
		isForward := true

		stopSign, found, err := GetNextStopSign(currentWay, nextWays, isForward)
		result := testResult{Found: found, Error: err}
		if found {
			result.Latitude = stopSign.Latitude()
			result.Longitude = stopSign.Longitude()
		}
		cupaloy.SnapshotT(t, result)
	})

	t.Run("Stop sign is behind on the current way", func(t *testing.T) {
		currentWay := CurrentWay{
			Way: way1,
			Distance: DistanceResult{
				LineStart: nodes1.At(2),
				LineEnd:   nodes1.At(3),
			},
		}
		nextWays := []NextWayResult{{Way: way2, IsForward: true}}
		isForward := true

		stopSign, found, err := GetNextStopSign(currentWay, nextWays, isForward)
		result := testResult{Found: found, Error: err}
		if found {
			result.Latitude = stopSign.Latitude()
			result.Longitude = stopSign.Longitude()
		}
		cupaloy.SnapshotT(t, result)
	})

	t.Run("Stop sign is on a future way", func(t *testing.T) {
		way1NoStops, _ := NewWay(seg)
		nodes1Copy, _ := way1NoStops.NewNodes(4)
		nodes1Copy.At(0).SetLatitude(0)
		nodes1Copy.At(0).SetLongitude(0)

		currentWay := CurrentWay{
			Way: way1NoStops,
			Distance: DistanceResult{
				LineStart: nodes1Copy.At(0),
				LineEnd:   nodes1Copy.At(1),
			},
		}
		nextWays := []NextWayResult{{Way: way2, IsForward: true}, {Way: way3, IsForward: true}}
		isForward := true

		stopSign, found, err := GetNextStopSign(currentWay, nextWays, isForward)
		result := testResult{Found: found, Error: err}
		if found {
			result.Latitude = stopSign.Latitude()
			result.Longitude = stopSign.Longitude()
		}
		cupaloy.SnapshotT(t, result)
	})

	t.Run("Stop sign direction does not match", func(t *testing.T) {
		currentWay := CurrentWay{
			Way: way1,
			Distance: DistanceResult{
				LineStart: nodes1.At(0),
				LineEnd:   nodes1.At(1),
			},
		}
		nextWays := []NextWayResult{{Way: way2, IsForward: true}}
		isForward := false // Traveling backward

		stopSign, found, err := GetNextStopSign(currentWay, nextWays, isForward)
		result := testResult{Found: found, Error: err}
		if found {
			result.Latitude = stopSign.Latitude()
			result.Longitude = stopSign.Longitude()
		}
		cupaloy.SnapshotT(t, result)
	})

	t.Run("No stop signs on path", func(t *testing.T) {
		way1NoStops, _ := NewWay(seg)
		nodes1Copy, _ := way1NoStops.NewNodes(2)
		nodes1Copy.At(0).SetLatitude(0)

		way2NoStops, _ := NewWay(seg)
		way2NoStops.NewNodes(1)

		currentWay := CurrentWay{
			Way: way1NoStops,
			Distance: DistanceResult{
				LineStart: nodes1Copy.At(0),
				LineEnd:   nodes1Copy.At(1),
			},
		}
		nextWays := []NextWayResult{{Way: way2NoStops, IsForward: true}}
		isForward := true

		stopSign, found, err := GetNextStopSign(currentWay, nextWays, isForward)
		result := testResult{Found: found, Error: err}
		if found {
			result.Latitude = stopSign.Latitude()
			result.Longitude = stopSign.Longitude()
		}
		cupaloy.SnapshotT(t, result)
	})
}
