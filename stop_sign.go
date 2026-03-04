package main

import (
	"pfeifer.dev/mapd/maps"
)

func checkWayForStopSignChange(state *State, parent *Upcoming[bool], way maps.NextWayResult) (valid bool, val bool) {
	if way.Way.HasStopSign(way.IsForward) {
		return true, true
	}

	return false, parent.DefaultValue
}
