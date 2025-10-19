package gocean

import (
	"math"
	"math/rand"
	"time"
)

// movementBehavior defines common movement patterns for entities
type movementBehavior struct {
	velocity     int
	direction    int
	moveInterval time.Duration
	nextMove     time.Duration
}

// newHorizontalMovement creates a movement behavior for horizontal motion
func newHorizontalMovement(velocity, direction int) *movementBehavior {
	moveInterval := time.Second / time.Duration(velocity)
	return &movementBehavior{
		velocity:     velocity,
		direction:    direction,
		moveInterval: moveInterval,
		nextMove:     moveInterval,
	}
}

// newVerticalMovement creates a movement behavior for vertical motion
func newVerticalMovement(velocity, direction int) *movementBehavior {
	moveInterval := time.Second / time.Duration(velocity)
	return &movementBehavior{
		velocity:     velocity,
		direction:    direction,
		moveInterval: moveInterval,
		nextMove:     moveInterval,
	}
}

// updateHorizontal updates horizontal position based on delta time
func (mb *movementBehavior) updateHorizontal(e *entity, dt time.Duration) {
	mb.nextMove -= dt
	for mb.nextMove <= 0 {
		e.x += mb.direction
		mb.nextMove += mb.moveInterval
	}
}

// updateVertical updates vertical position based on delta time
func (mb *movementBehavior) updateVertical(e *entity, dt time.Duration) {
	mb.nextMove -= dt
	for mb.nextMove <= 0 {
		e.y += mb.direction
		mb.nextMove += mb.moveInterval
	}
}

// isOffScreenHorizontal checks if entity has moved off screen horizontally
func isOffScreenHorizontal(e *entity, screenWidth int) bool {
	return e.x < -e.s.w || e.x > screenWidth
}

// isOffScreenVertical checks if entity has moved off screen vertically
func isOffScreenVertical(e *entity, screenHeight int) bool {
	return e.y < -e.s.h || e.y > screenHeight
}

// shouldSpawnBubble determines if a bubble should spawn using Poisson distribution
func shouldSpawnBubble(dt time.Duration, rate float64) bool {
	expectedBubbles := rate * dt.Seconds()
	probability := 1 - math.Exp(-expectedBubbles)
	return rand.Float64() < probability
}

// randomFishVelocity generates a random velocity for fish movement
func randomFishVelocity() int {
	return rand.Intn(maxFishVelocity-minFishVelocity+1) + minFishVelocity
}

// randomSeaweedHeight generates a random height for seaweed
func randomSeaweedHeight() int {
	return rand.Intn(maxSeaweedHeight-minSeaweedHeight) + minSeaweedHeight
}

// calculateFishCount determines how many fish to spawn based on screen size
func calculateFishCount(screenWidth, screenHeight int) int {
	screenSize := (screenHeight - waterSurfaceOffset) * screenWidth
	return screenSize / fishDensityFactor
}

// collisionDetector provides collision detection functionality
type collisionDetector struct {
	pointX, pointY int
}

// newCollisionDetector creates a collision detector for a specific point
func newCollisionDetector(x, y int) *collisionDetector {
	return &collisionDetector{pointX: x, pointY: y}
}

// checkCollisionWithFish checks if the detector point collides with any fish
func (cd *collisionDetector) checkCollisionWithFish(m *model) *entity {
	for i := zFishStart; i < zFishEnd; i++ {
		fishies := m.entities[i]
		for _, fish := range fishies {
			if !fish.physical {
				continue
			}

			// Check if collision point is within fish bounds
			if cd.pointX >= fish.x && cd.pointX <= fish.x+fish.s.w &&
				cd.pointY >= fish.y && cd.pointY <= fish.y+fish.s.h {
				return fish
			}
		}
	}

	return nil
}
