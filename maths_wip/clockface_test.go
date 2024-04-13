package mathswip

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{x: 150, y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSecondsInRadians(t *testing.T) {
	ts := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)

	want := math.Pi
	got := secondsInRadians(ts)

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
