package walker

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"googlemaps.github.io/maps"
)

type Walker interface {
	PositionComputation(string, string) (string, time.Duration, error)
}

func NewWalker() Walker {
	if os.Getenv("mode_mock") != "" {
		log.Printf("The walker will work in mock mode")
		return &MockedWalker{}
	}
	return &ConcreteWalker{}
}

type ConcreteWalker struct{}

// positionComputation computate where we are after 1h walk
func (w *ConcreteWalker) PositionComputation(from, to string) (string, time.Duration, error) {
	// initialise google api library
	envKey := os.Getenv("MAPS_KEY")
	if envKey == "" {
		log.Fatalf("no maps api key set")
	}
	c, err := maps.NewClient(maps.WithAPIKey(envKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// get the actual information
	r := &maps.DirectionsRequest{
		Origin:      from,
		Destination: to,
		Mode:        maps.TravelModeWalking,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		return "", time.Duration(0), fmt.Errorf("error while getting the directions: %s", err.Error())
	}

	fmt.Printf("%+v\n", route[0].Legs)

	// computate 1h walk
	return w.computate1hWalk(route)
}

func (w *ConcreteWalker) computate1hWalk(route []maps.Route) (string, time.Duration, error) {
	var totaltime time.Duration
	var actualPosition string
	for i := 0; i < len(route); i++ {
		for j := 0; j < len(route[i].Legs); j++ {
			totaltime += route[i].Legs[j].Duration
			actualPosition = route[i].Legs[j].EndAddress
			if totaltime >= time.Duration(1)*time.Hour {
				return actualPosition, totaltime, nil
			}
		}
		if totaltime >= time.Duration(1)*time.Hour {
			return actualPosition, totaltime, nil
		}
	}
	return actualPosition, totaltime, nil
}
