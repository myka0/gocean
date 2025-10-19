package gocean

import (
	"time"
)

// entity represents any drawable object in the ocean simulation
type entity struct {
	// Visual representation
	s sprite

	// Position
	x, y, z int

	// Animation state
	frame      int
	frameDelay time.Duration
	lastFrame  time.Time

	// Physics and lifecycle
	physical bool
	alive    bool

	// Behavior hooks
	onTick func(*model, *entity, time.Duration)
	onDie  func(*model)
}

// AdvanceFrame updates the entity's animation frame based on its frameDelay
func (e *entity) AdvanceFrame() {
	// Only advance if enough time has passed since the last frame change
	if time.Since(e.lastFrame) >= e.frameDelay {
		// Cycle to next frame, wrapping around to 0 after the last frame
		e.frame = (e.frame + 1) % max(1, len(e.s.frames))
		e.lastFrame = time.Now()
	}
}
