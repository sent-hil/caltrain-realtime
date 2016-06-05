package main

import (
	"log"

	"github.com/sent-hil/caltrain-realtime"
)

func main() {
	// SanFrancisco is end of line, has only SouthBound service.
	assert(caltrain.SanFrancisco, caltrain.SouthBound, 3)
	assert(caltrain.SanFrancisco, caltrain.NorthBound, 0)

	// Weekend only service.
	assert(caltrain.Broadway, caltrain.SouthBound, 0)
	assert(caltrain.Broadway, caltrain.NorthBound, 0)
	assert(caltrain.Atherton, caltrain.SouthBound, 0)
	assert(caltrain.Atherton, caltrain.NorthBound, 0)

	assert(caltrain.Gilroy, caltrain.SouthBound, 0)
	assert(caltrain.Gilroy, caltrain.NorthBound, 0)

	assert(caltrain.PaloAlto, caltrain.SouthBound, 3)
	assert(caltrain.PaloAlto, caltrain.NorthBound, 3)

	assert(caltrain.MountainView, caltrain.SouthBound, 3)
	assert(caltrain.MountainView, caltrain.NorthBound, 3)

	// Possibly SouthBound has only 2 train.
	assert(caltrain.Tamien, caltrain.SouthBound, 2)
	assert(caltrain.Tamien, caltrain.NorthBound, 3)
}

func assert(s caltrain.Station, d caltrain.Direction, l int) {
	timings, err := caltrain.GetRealTimings(s, d)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return
	}
	if len(timings) != l {
		log.Printf(
			"ERROR: Expected '%d' durations for '%s'->'%d', was '%d': '%v'",
			l, s, d, len(timings), timings,
		)
		return
	}
	if len(timings) > 0 {
		log.Printf("Passed '%s'->'%d': '%v'.", s, d, timings)
	} else {
		log.Printf("Passed '%s'->'%d': no trains.", s, d)
	}
}
