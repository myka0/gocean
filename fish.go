package gocean

import (
	"math/rand"
	"time"

	"github.com/myka0/gocean/internal/art"
)

// addAllFish spawns fish based on screen size
func (m *model) addAllFish() {
	fishCount := calculateFishCount(m.windowWidth, m.windowHeight)
	for range fishCount {
		m.addFish()
	}
}

// addFish creates and adds a single fish entity to the aquarium
func (m *model) addFish() {
	fishNum := rand.Intn(len(art.Fishies))
	fishRaw := art.Fishies[fishNum]
	fish := newSprite([]string{fishRaw.Frame}, []string{fishRaw.Mask})

	depth := int(rand.Intn(zFishEnd-zFishStart)) + zFishStart
	velocity := randomFishVelocity()
	startX, direction := m.createMovement(fishNum, fish.w)

	e := &entity{
		s:        fish,
		x:        startX,
		y:        rand.Intn(m.windowHeight-waterSurfaceOffset-fish.h) + waterSurfaceOffset,
		z:        depth,
		alive:    true,
		physical: true,
	}

	// Store movement behavior in entity for tick updates
	horizontalMovement := newHorizontalMovement(velocity, direction)

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		// Update horizontal movement
		horizontalMovement.updateHorizontal(ee, dt)

		bubbleX := ee.x
		if horizontalMovement.direction > 0 {
			bubbleX += ee.s.w
		}
		bubbleY := ee.y + ee.s.h/2

		// Spawn bubbles using Poisson distribution
		if shouldSpawnBubble(dt, bubbleSpawnRate) {
			mm.addBubble(bubbleX, bubbleY, ee.z-1)
		}

		// Kill fish if it swims off screen
		if isOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	e.onDie = func(mm *model) { mm.addFish() }
	m.entities[e.z] = append(m.entities[e.z], e)
}

// createFishMovement determines initial position and direction for a fish
func (m *model) createMovement(num, width int) (int, int) {
	if num%2 == 1 {
		// Fish swimming left
		return m.windowWidth - 1, -1
	}
	// Fish swimming right
	return 1 - width, 1
}

// addBubble creates a bubble entity that rises from the given position
func (m *model) addBubble(x, y, z int) {
	bubble := newSprite([]string{".", "o", "O"}, []string{"c", "c", "c"})
	verticalMovement := newVerticalMovement(bubbleRiseSpeed, -1)

	e := &entity{
		s:          bubble,
		x:          x,
		y:          y,
		z:          z,
		frameDelay: bubbleFrameDelay,
		lastFrame:  time.Now(),
		alive:      true,
	}

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		// Animate bubble growth
		if ee.frame < len(ee.s.frames)-1 {
			ee.AdvanceFrame()
		}

		// Move bubble upward
		verticalMovement.updateVertical(ee, dt)

		// Kill bubble when it reaches the surface
		if ee.y <= waterSurfaceOffset {
			ee.alive = false
		}
	}

	m.entities[e.z] = append(m.entities[e.z], e)
}
