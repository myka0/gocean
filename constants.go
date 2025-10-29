package gocean

import "time"

const (
	// Z axis layers (higher draws first)
	zShark     = 2
	zFishStart = 3
	zFishEnd   = 20
	zSeaweed   = 21
	zCastle    = 22

	// Waterline layers
	zWaterGap0  = 9
	zWaterLine0 = 8
	zWaterGap1  = 7
	zWaterLine1 = 6
	zWaterGap2  = 5
	zWaterLine2 = 4
	zWaterGap3  = 3
	zWaterLine3 = 2

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
	swanVelocity      = 10  // chars/sec
	whaleVelocity     = 10  // chars/sec
	ducksVelocity     = 10  // chars/sec
	dolphinVelocity   = 20  // chars/sec
	fishHookVelocity  = 10  // chars/sec
	monsterVelocity   = 20  // chars/sec
	bigFishVelocity   = 30  // chars/sec
	sharkVelocity     = 30  // chars/sec

	// Seaweed constants
	minSeaweedHeight = 3
	maxSeaweedHeight = 7

	// Animation timing
	bubbleFrameDelay     = 1 * time.Second
	monsterFrameDelay    = 200 * time.Millisecond
	waterSpoutFrameDelay = 100 * time.Millisecond
	waterSpoutResetDelay = 1000 * time.Millisecond
	splatFrameDelay      = 500 * time.Millisecond
	duckFrameDelay       = 500 * time.Millisecond

	// Special entity Y positions
	yShipPosition    = 0
	yWhalePosition   = 3
	yMonsterPosition = 2
	ySwanPosition    = 1
	yDuckPosition    = 5
	yDolphinTop      = 1
	yDolphinMid      = 4
	yDolphinBottom   = 8
)

// ANSI color mapping
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
