package main

import (
	"math/rand"
	"strings"
	"time"

	"gocean/internal/art"
)

// addEnvironment creates the water surface waves that appear at the top of the screen
func (m *model) addEnvironment() {
	// Build the complete wave pattern by repeating segments
	for i, s := range art.WaterSegments {
		rep := (m.windowWidth / len(s)) + 1
		var image, mask string
		for range rep {
			image += s
			// Create mask with cyan color for all characters
			mask += strings.Repeat("c", len(s))
		}

		// Create the environment sprite and entity
		sprite := newSprite([]string{image}, []string{mask})
		e := &Entity{
			s:     sprite,
			x:     0,
			y:     5 + i,
			z:     zWaterLine0 - i*2,
			alive: true,
		}

		m.entities[e.z] = append(m.entities[e.z], e)
	}
}

// addCastle creates the decorative castle that appears on the right side of the screen.
func (m *model) addCastle() {
	// Create castle sprite and position it on the right side of the screen
	sprite := newSprite([]string{art.Castle.Frame}, []string{art.Castle.Mask})
	e := &Entity{
		s:     sprite,
		x:     m.windowWidth - 32,
		y:     m.windowHeight - 13,
		z:     zCastle,
		alive: true,
	}

	m.entities[e.z] = append(m.entities[e.z], e)
}

// addAllSeaweed populates the ocean floor with seaweed based on screen width
func (m *model) addAllSeaweed() {
	seaweedCount := int(m.windowWidth / 15)
	for range seaweedCount {
		m.addSeaweed()
	}
}

// addSeaweed creates a single animated seaweed plant at a random position
func (m *model) addSeaweed() {
	height := RandomSeaweedHeight()

	// Create two animation frames for swaying motion
	frameA := strings.Repeat("(\n )\n", height)
	frameA = strings.TrimSuffix(frameA, "\n")

	frameB := strings.Repeat(" )\n(\n", height)
	frameB = strings.TrimSuffix(frameB, "\n")

	// Color mask
	mask := strings.Repeat("gg\ngg\n", height)
	mask = strings.TrimSuffix(mask, "\n")

	sp := newSprite([]string{frameA, frameB}, []string{mask, mask})

	// Random positioning and animation speed
	x := rand.Intn(m.windowWidth-2) + 1
	y := m.windowHeight - height
	animSpeed := rand.Intn(100+1) + 400

	e := &Entity{
		s:          sp,
		x:          x,
		y:          y,
		z:          zSeaweed,
		frameDelay: time.Duration(animSpeed) * time.Millisecond,
		alive:      true,
	}

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) { ee.AdvanceFrame() }
	e.onDie = func(mm *model) { mm.addSeaweed() }
	m.entities[e.z] = append(m.entities[e.z], e)
}
