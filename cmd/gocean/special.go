package main

import (
	"math/rand"
	"time"
)

// addSpecial creates special entities
func (m *model) addSpecial() {
	m.addShip()
}

// addShip creates a sailing ship that moves across the water surface
func (m *model) addShip() {
	var ship = []RawStaticArt{
		{
			frame: `     |    |    |
    )_)  )_)  )_)
   )___))___))___)\
  )____)____)_____)\\
_____|____|____|____\\\__
\                   /`,
			mask: `     y    y    y
    www  www  www
   wwwwwwwwwwwwwwww
  wwwwwwwwwwwwwwwwwww
yyyyyyyyyyyyyyyyyyyywwwyy
y???????????????????y`,
		},
		{
			frame: `         |    |    |
        (_(  (_(  (_(
      /(___((___((___(
    //(_____(____(____(
__///____|____|____|_____
    \                   /`,
			mask: `         y    y    y
        www  www  www
      wwwwwwwwwwwwwwww
    wwwwwwwwwwwwwwwwwww
yywwwyyyyyyyyyyyyyyyyyyyy
    y???????????????????y`,
		},
	}

	// Randomly select left or right facing ship
	num := int(rand.Intn(2))
	sprite := newSprite([]string{ship[num].frame}, []string{ship[num].mask})

	// Determine starting position and movement direction
	startX, direction := m.createMovement(num, sprite.w)

	e := &Entity{
		s:     sprite,
		x:     startX,
		y:     0,
		z:     zWaterGap0,
		alive: true,
	}

	// Create horizontal movement behavior for the ship
	horizontalMovement := NewHorizontalMovement(shipVelocity, direction)

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		horizontalMovement.UpdateHorizontal(ee, dt)

		// Kill ship when it moves completely off screen
		if IsOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	e.onDie = func(mm *model) { mm.addSpecial() }
	m.entities[e.z] = append(m.entities[e.z], e)
}
