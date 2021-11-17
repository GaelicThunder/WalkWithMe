package walker

import (
	"testing"
	"time"

	"googlemaps.github.io/maps"
)

func Test1hCalculation(t *testing.T) {
	var w ConcreteWalker

	routes := []maps.Route{
		{
			Legs: []*maps.Leg{
				{
					Duration: 50 * time.Minute,
				},
				{
					Duration: 23 * time.Minute,
				},
			},
		},
		{
			Legs: []*maps.Leg{
				{
					Duration: 2 * time.Hour,
				},
				{
					Duration: 2 * time.Hour,
				},
			},
		},
	}
	_, totalwalk, _ := w.computate1hWalk(routes)
	if totalwalk != (1*time.Hour + 13*time.Minute) {
		t.Fatalf("Expected 1h but having %+v\n", totalwalk)
	}
}
