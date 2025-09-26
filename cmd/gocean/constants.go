package main

import "time"

const (
	// Fame rate
	tickEvery = time.Second / 120

	// Z axis layers (higher draws first)
	zShark     = 2
	zFishStart = 3
	zFishEnd   = 20
	zSeaweed   = 21
	zCastle    = 22

	// Waterline layers
	zWaterLine0 = 9
	zWaterGap0  = 8
	zWaterLine1 = 7
	zWaterGap1  = 6
	zWaterLine2 = 5
	zWaterGap2  = 4
	zWaterLine3 = 3
	zWaterGap3  = 2

	// Aquarium dimensions
	minAquariumWidth  = 32
	minAquariumHeight = 16
	defaultWidth      = 80
	defaultHeight     = 24

	// Water surface offset from top
	waterSurfaceOffset = 9

	// Fish behavior constants
	fishDensityFactor = 350 // screen pixels per fish
	minFishVelocity   = 3   // chars/sec
	maxFishVelocity   = 30  // chars/sec
	bubbleSpawnRate   = 0.2 // bubbles per second per fish
	bubbleRiseSpeed   = 10  // chars/sec
	shipVelocity      = 10  // chars/sec

	// Seaweed constants
	minSeaweedHeight = 3
	maxSeaweedHeight = 7

	// Animation timing
	bubbleFrameDelay = 1 * time.Second
)

// Random Pallete
var (
	ansi16 = map[byte]string{
		'r': "1", 'R': "9",
		'g': "2", 'G': "10",
		'y': "3", 'Y': "11",
		'b': "4", 'B': "12",
		'm': "5", 'M': "13",
		'c': "6", 'C': "14",
		'w': "7", 'W': "15",
	}
)
