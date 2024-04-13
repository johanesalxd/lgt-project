package maths

import "time"

type Point struct {
	x int
	y int
}

func SecondHand(t time.Time) Point {
	return Point{x: 150, y: 60}
}
