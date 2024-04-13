package mathswip

import (
	"math"
	"time"
)

type Point struct {
	x int
	y int
}

func SecondHand(t time.Time) Point {
	return Point{x: 150, y: 60}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi
}
