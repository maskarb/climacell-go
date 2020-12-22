package main

import (
	"os"
	"time"

	climacellv4 "github.com/maskarb/climacell-go/climacell/v4"
)

func main() {
	start := time.Now()
	end := start.Add(time.Hour * 48)
	c := climacellv4.NewClient(os.Getenv("CLIMACELL_API_KEY"))

	opts := &climacellv4.TimelineListOptions{
		Location: climacellv4.Geometry{
			Type:        "Point",
			Coordinates: []string{"-78.613375", "35.816735"},
		},
		Fields:    []string{"temperature"},
		StartTime: start.Format(time.RFC3339),
		EndTime:   end.Format(time.RFC3339),
		TimeSteps: []string{"1d"},
	}

}
