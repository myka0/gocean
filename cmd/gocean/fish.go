package main

import (
	"math/rand"
	"time"
)

// 1: Body
// 2: Dorsal Fin
// 3: Flippers
// 4: Eye
// 5: Mouth
// 6: Tailfin
// 7: Gills
var fishies = []RawStaticArt{
	{
		frame: `       \
     ...\..,
\  /'       \
 >=     (  ' >
/  \      / /
    b"'"'/''`,
		mask: `       2
     1112111
6  11???????1
 66?????7??4?5
6  1??????3?1
    11111311`,
	},

	{
		frame: `      /
  ,../...
 /       '\  /
< '  )     =<
 \ \      /  \
  b'\'"'"'`,
		mask: `      2
  1112111
 1???????11  6
5?4??7?????66
 1?3??????1  6
  11311111`,
	},

	{
		frame: `    \
\ /--\
>=  (o>
/ \__/
    /`,
		mask: `    2
6 1111
66??745
6 1111
    3`,
	},

	{
		frame: `  /
 /--\ /
<o)  =<
 \__/ \
  \`,
		mask: `  2
 1111 6
547??66
 1111 6
  3`,
	},

	{
		frame: `       \:.
\;,   ,;\\\,,
  \\\;;:::::::o
  ///;;::::::::<
 /;b bb/////bb`,
		mask: `       222
666   1122211
  6661111111114
  66611111111115
 666 113333311`,
	},

	{
		frame: `      .:/
   ,,///;,   ,;/
 o:::::::;;///
>::::::::;;\\\
  ''\\\\\'' ';\`,
		mask: `      222
   1122211   666
 4111111111666
51111111111666
  113333311 666`,
	},

	{
		frame: `  __
><_'>
   '`,
		mask: `  11
61145
   3`,
	},

	{
		frame: ` __
<'_><
 b`,
		mask: ` 11
54116
 3`,
	},

	{
		frame: `   ..\,
>='   ('>
  '''/''`,
		mask: `   1121
661???745
  111311`,
	},

	{
		frame: `  ,/..
<')   b=<
 bb\bbb`,
		mask: `  1211
547???166
 113111`,
	},

	{
		frame: `   \
  / \
>=_('>
  \_/
   /`,
		mask: `   2
  1?1
661745
  111
   3`,
	},

	{
		frame: `  /
 / \
<')_=<
 \_/
  \`,
		mask: `  2
 1?1
547166
 111
  3`,
	},

	{
		frame: `  ,\
>=('>
  '/`,
		mask: `  12
66745
  13`,
	},

	{
		frame: ` /,
<')=<
 \b`,
		mask: ` 21
54766
 31`,
	},

	{
		frame: `  __
\/ o\
/\__/`,
		mask: `  11
61?41
61111`,
	},

	{
		frame: ` __
/o \/
\__/\`,
		mask: ` 11
14?16
11116`,
	},
}

// addAllFish spawns fish based on screen size
func (m *model) addAllFish() {
	fishCount := CalculateFishCount(m.windowWidth, m.windowHeight)
	for range fishCount {
		m.addFish()
	}
}

// addFish creates and adds a single fish entity to the aquarium
func (m *model) addFish() {
	fishNum := rand.Intn(len(fishies))
	fishRaw := fishies[fishNum]
	fish := newSprite([]string{fishRaw.frame}, []string{fishRaw.mask})

	depth := int(rand.Intn(zFishEnd-zFishStart)) + zFishStart
	velocity := RandomFishVelocity()
	startX, direction := m.createMovement(fishNum, fish.w)

	e := &Entity{
		s:        fish,
		x:        startX,
		y:        rand.Intn(m.windowHeight-waterSurfaceOffset-fish.h) + waterSurfaceOffset,
		z:        depth,
		alive:    true,
		physical: true,
	}

	// Store movement behavior in entity for tick updates
	horizontalMovement := NewHorizontalMovement(velocity, direction)

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		// Update horizontal movement
		horizontalMovement.UpdateHorizontal(ee, dt)

		x := ee.x
		if horizontalMovement.direction > 0 {
			x += ee.s.w
		}
		y := ee.y + ee.s.h/2

		// Spawn bubbles using Poisson distribution
		if ShouldSpawnBubble(dt, bubbleSpawnRate) {
			mm.addBubble(x, y, ee.z-1)
		}

		// Kill fish if it swims off screen
		if IsOffScreenHorizontal(ee, mm.windowWidth) {
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
	verticalMovement := NewVerticalMovement(bubbleRiseSpeed, -1)

	e := &Entity{
		s:          bubble,
		x:          x,
		y:          y,
		z:          z,
		frameDelay: bubbleFrameDelay,
		lastFrame:  time.Now(),
		alive:      true,
	}

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		// Animate bubble growth
		if ee.frame < len(ee.s.frames)-1 {
			ee.AdvanceFrame()
		}

		// Move bubble upward
		verticalMovement.UpdateVertical(ee, dt)

		// Kill bubble when it reaches the surface
		if ee.y <= waterSurfaceOffset {
			ee.alive = false
		}
	}

	m.entities[e.z] = append(m.entities[e.z], e)
}
