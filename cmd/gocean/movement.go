package main

import (
	"math"
	"math/rand"
	"time"
)

// MovementBehavior defines common movement patterns for entities
type MovementBehavior struct {
	velocity     int
	direction    int
	moveInterval time.Duration
	nextMove     time.Duration
}

// NewHorizontalMovement creates a movement behavior for horizontal motion
func NewHorizontalMovement(velocity, direction int) *MovementBehavior {
	moveInterval := time.Second / time.Duration(velocity)
	return &MovementBehavior{
		velocity:     velocity,
		direction:    direction,
		moveInterval: moveInterval,
		nextMove:     moveInterval,
	}
}

// NewVerticalMovement creates a movement behavior for vertical motion
func NewVerticalMovement(velocity, direction int) *MovementBehavior {
	moveInterval := time.Second / time.Duration(velocity)
	return &MovementBehavior{
		velocity:     velocity,
		direction:    direction,
		moveInterval: moveInterval,
		nextMove:     moveInterval,
	}
}

// UpdateHorizontal updates horizontal position based on delta time
func (mb *MovementBehavior) UpdateHorizontal(e *Entity, dt time.Duration) {
	mb.nextMove -= dt
	for mb.nextMove <= 0 {
		e.x += mb.direction
		mb.nextMove += mb.moveInterval
	}
}

// UpdateVertical updates vertical position based on delta time
func (mb *MovementBehavior) UpdateVertical(e *Entity, dt time.Duration) {
	mb.nextMove -= dt
	for mb.nextMove <= 0 {
		e.y += mb.direction
		mb.nextMove += mb.moveInterval
	}
}

// IsOffScreenHorizontal checks if entity has moved off screen horizontally
func IsOffScreenHorizontal(e *Entity, screenWidth int) bool {
	return e.x < -e.s.w || e.x > screenWidth
}

// IsOffScreenVertical checks if entity has moved off screen vertically
func IsOffScreenVertical(e *Entity, screenHeight int) bool {
	return e.y < -e.s.h || e.y > screenHeight
}

// ShouldSpawnBubble determines if a bubble should spawn using Poisson distribution
func ShouldSpawnBubble(dt time.Duration, rate float64) bool {
	expectedBubbles := rate * dt.Seconds()
	probability := 1 - math.Exp(-expectedBubbles)
	return rand.Float64() < probability
}

// RandomFishVelocity generates a random velocity for fish movement
func RandomFishVelocity() int {
	return rand.Intn(maxFishVelocity-minFishVelocity+1) + minFishVelocity
}

// RandomSeaweedHeight generates a random height for seaweed
func RandomSeaweedHeight() int {
	return rand.Intn(maxSeaweedHeight-minSeaweedHeight) + minSeaweedHeight
}

// CalculateFishCount determines how many fish to spawn based on screen size
func CalculateFishCount(screenWidth, screenHeight int) int {
	screenSize := (screenHeight - waterSurfaceOffset) * screenWidth
	return screenSize / fishDensityFactor
}

// CollisionDetector provides collision detection functionality
type CollisionDetector struct {
	pointX, pointY int
}

// NewCollisionDetector creates a collision detector for a specific point
func NewCollisionDetector(x, y int) *CollisionDetector {
	return &CollisionDetector{pointX: x, pointY: y}
}

// CheckCollisionWithFish checks if the detector point collides with any fish
func (cd *CollisionDetector) CheckCollisionWithFish(m *model) *Entity {
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
