package main

import (
	"math/rand"
	"strings"
	"time"

	"gocean/internal/art"
)

// Special entity types
type SpecialEntityType int

const (
	Ship SpecialEntityType = iota
	Whale
	Monster
	BigFish
	Shark
	FishHook
	Swan
	Ducks
	Dolphins
	SpecialEntityCount
)

// addSpecial creates special entities using improved selection and configuration
func (m *model) addSpecial() {
	specialType := SpecialEntityType(rand.Intn(int(SpecialEntityCount)))

	switch specialType {
	case Ship:
		m.addShip()
	case Whale:
		m.addWhale()
	case Monster:
		m.addMonster()
	case BigFish:
		m.addBigFish()
	case Shark:
		m.addShark()
	case FishHook:
		m.addFishHook()
	case Swan:
		m.addSwan()
	case Ducks:
		m.addDucks()
	case Dolphins:
		m.addDolphins()
	}
}

// createHorizontalEntity creates a basic horizontal-moving entity with common setup
func (m *model) createHorizontalEntity(dir int, y int, z int, velocity int, sprite Sprite) *Entity {
	startX, direction := m.createMovement(dir, sprite.w)

	e := &Entity{
		s:     sprite,
		x:     startX,
		y:     y,
		z:     z,
		alive: true,
	}

	horizontalMovement := NewHorizontalMovement(velocity, direction)

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		horizontalMovement.UpdateHorizontal(ee, dt)

		if IsOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	e.onDie = func(mm *model) { mm.addSpecial() }
	return e
}

// createAnimatedHorizontalEntity creates an animated horizontal-moving entity
func (m *model) createAnimatedHorizontalEntity(dir int, y int, z int, velocity int, frameDelay time.Duration, sprite Sprite) *Entity {
	startX, direction := m.createMovement(dir, sprite.w)

	e := &Entity{
		s:          sprite,
		x:          startX,
		y:          y,
		z:          z,
		frameDelay: frameDelay,
		alive:      true,
	}

	horizontalMovement := NewHorizontalMovement(velocity, direction)

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		ee.AdvanceFrame()
		horizontalMovement.UpdateHorizontal(ee, dt)

		if IsOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	e.onDie = func(mm *model) { mm.addSpecial() }
	return e
}

// addShip creates a sailing ship that moves across the water surface
func (m *model) addShip() {
	dir := rand.Intn(2)
	sprite := newSprite([]string{art.Ship[dir].Frame}, []string{art.Ship[dir].Mask})

	e := m.createHorizontalEntity(dir, yShipPosition, zWaterGap1, shipVelocity, sprite)
	m.entities[e.z] = append(m.entities[e.z], e)
}

// addBigFish creates a big fish
func (m *model) addBigFish() {
	dir := rand.Intn(2)
	sprite := newSprite([]string{art.BigFish[dir].Frame}, []string{art.BigFish[dir].Mask})
	y := rand.Intn(m.windowHeight-waterSurfaceOffset-sprite.h) + waterSurfaceOffset

	e := m.createHorizontalEntity(dir, y, zShark, bigFishVelocity, sprite)
	m.entities[e.z] = append(m.entities[e.z], e)
}

// addSwan creates a swan that moves across the water surface
func (m *model) addSwan() {
	dir := rand.Intn(2)
	sprite := newSprite([]string{art.Swan[dir].Frame}, []string{art.Swan[dir].Mask})

	e := m.createHorizontalEntity(dir, ySwanPosition, zWaterGap3, swanVelocity, sprite)
	m.entities[e.z] = append(m.entities[e.z], e)
}

// addMonster creates a monster that moves across the water surface
func (m *model) addMonster() {
	dir := rand.Intn(2)
	sprite := newSprite(art.MonsterFrames[dir], art.MonsterMasks[dir])

	e := m.createAnimatedHorizontalEntity(dir, yMonsterPosition, zWaterGap2, monsterVelocity, monsterFrameDelay, sprite)
	m.entities[e.z] = append(m.entities[e.z], e)
}

// addDucks creates ducks that move across the water surface
func (m *model) addDucks() {
	dir := rand.Intn(2)
	mask := []string{art.DuckMasks[dir], art.DuckMasks[dir], art.DuckMasks[dir]}
	sprite := newSprite(art.DuckFrames[dir], mask)

	e := m.createAnimatedHorizontalEntity(dir, yDuckPosition, zWaterGap3, ducksVelocity, duckFrameDelay, sprite)
	m.entities[e.z] = append(m.entities[e.z], e)
}

// addWhale creates a whale that moves across the water surface with water spout
func (m *model) addWhale() {
	// Randomly select left or right facing whale
	num := rand.Intn(2)
	sprite := newSprite([]string{art.Whale[num].Frame}, []string{art.Whale[num].Mask})

	// Determine starting position and movement direction
	startX, direction := m.createMovement(num, sprite.w)

	e := &Entity{
		s:     sprite,
		x:     startX,
		y:     yWhalePosition,
		z:     zWaterGap2,
		alive: true,
	}

	horizontalMovement := NewHorizontalMovement(whaleVelocity, direction)

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		horizontalMovement.UpdateHorizontal(ee, dt)

		// Kill whale when it moves completely off screen
		if IsOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	// Create water spout with proper positioning
	spoutX := startX + 11
	if direction == -1 {
		spoutX = startX + 1
	}
	waterSpout := m.addWaterSpout(spoutX, 0, direction)

	e.onDie = func(mm *model) {
		waterSpout.alive = false
		mm.addSpecial()
	}

	m.entities[e.z] = append(m.entities[e.z], e)
	m.entities[waterSpout.z] = append(m.entities[waterSpout.z], waterSpout)
}

// addWaterSpout creates a water spout animation with improved frame management
func (m *model) addWaterSpout(x int, y int, dir int) *Entity {
	// Create water spout sprite
	mask := make([]string, len(art.WaterSpoutFrames))
	for i := range mask {
		mask[i] = art.WaterSpoutMask
	}
	sprite := newSprite(art.WaterSpoutFrames, mask)

	e := &Entity{
		s:          sprite,
		x:          x,
		y:          y,
		z:          zWaterGap2,
		frameDelay: waterSpoutFrameDelay,
		alive:      true,
	}

	// Create horizontal movement behavior
	horizontalMovement := NewHorizontalMovement(whaleVelocity, dir)

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		// Create delay when spout resets
		if ee.frame == 0 {
			ee.frameDelay = waterSpoutResetDelay
		}
		ee.AdvanceFrame()
		if ee.frame == 1 {
			ee.frameDelay = waterSpoutFrameDelay
		}

		horizontalMovement.UpdateHorizontal(ee, dt)
	}

	return e
}

// addShark creates a shark with improved collision detection
func (m *model) addShark() {
	// Randomly select left or right facing shark
	num := rand.Intn(2)
	sprite := newSprite([]string{art.Shark[num].Frame}, []string{art.Shark[num].Mask})

	// Determine starting position and movement direction
	startX, direction := m.createMovement(num, sprite.w)

	e := &Entity{
		s:     sprite,
		x:     startX,
		y:     rand.Intn(m.windowHeight-waterSurfaceOffset-sprite.h) + waterSurfaceOffset,
		z:     zShark,
		alive: true,
	}

	// Create horizontal movement behavior
	horizontalMovement := NewHorizontalMovement(sharkVelocity, direction)

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		horizontalMovement.UpdateHorizontal(ee, dt)

		// Calculate teeth position with improved logic
		teethX := ee.x + 50
		if direction == -1 {
			teethX = ee.x + 3
		}
		teethY := ee.y + 8

		// Use collision detector for better performance
		collisionDetector := NewCollisionDetector(teethX, teethY)
		if fish := collisionDetector.CheckCollisionWithFish(mm); fish != nil {
			fish.alive = false
			mm.addSplat(teethX, teethY)
			mm.addFish()
		}

		if IsOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	e.onDie = func(mm *model) { mm.addSpecial() }
	m.entities[e.z] = append(m.entities[e.z], e)
}

// addSplat creates a splat effect using entity pooling
func (m *model) addSplat(x, y int) {
	splat := newSprite(art.SplatFrames, art.SplatMasks)

	e := &Entity{
		s:          splat,
		x:          x - 3,
		y:          y - 3,
		z:          zShark - 1,
		frameDelay: splatFrameDelay,
		alive:      true,
	}

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		ee.AdvanceFrame()
		if ee.frame == 0 {
			ee.alive = false
		}
	}

	m.entities[e.z] = append(m.entities[e.z], e)
}

// addFishHook creates a fish hook with improved bounds checking and validation
func (m *model) addFishHook() {
	maxHeight := m.windowHeight/4 + m.windowHeight/2
	x := 10 + rand.Intn(m.windowWidth-40)

	sprite := newSprite([]string{art.FishHook.Frame}, []string{art.FishHook.Mask})

	hook := &Entity{
		s:     sprite,
		x:     x,
		y:     -sprite.h,
		z:     zWaterGap1,
		alive: true,
	}

	// Create vertical movement behavior
	up := NewVerticalMovement(fishHookVelocity, -1)
	down := NewVerticalMovement(fishHookVelocity, 1)
	fishMovement := NewVerticalMovement(fishHookVelocity, -1)
	verticalMovement := down

	// Create fishing line and movement behavior
	line := m.addFishingLine(x, sprite.h, maxHeight)
	lineUp := NewVerticalMovement(fishHookVelocity, -1)
	lineDown := NewVerticalMovement(fishHookVelocity, 1)
	lineMovement := lineDown

	hookedFish := false

	hook.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		if hookedFish || ee.y+ee.s.h < maxHeight {
			verticalMovement.UpdateVertical(ee, dt)
			lineMovement.UpdateVertical(line, dt)
		}

		pointX := ee.x + 1
		pointY := ee.y + 3

		// Use collision detector for fish hook
		if !hookedFish {
			collisionDetector := NewCollisionDetector(pointX, pointY)
			if fish := collisionDetector.CheckCollisionWithFish(mm); fish != nil {
				hookedFish = true
				verticalMovement = up
				lineMovement = lineUp
				fish.onTick = func(mm *model, ee *Entity, dt time.Duration) {
					fishMovement.UpdateVertical(fish, dt)
					if IsOffScreenVertical(ee, mm.windowHeight) {
						ee.alive = false
					}
				}
			}
		}

		// Kill fish hook when it moves completely off screen
		if IsOffScreenVertical(ee, mm.windowHeight) {
			ee.alive = false
			line.alive = false
		}
	}

	hook.onDie = func(mm *model) { mm.addSpecial() }
	m.entities[hook.z] = append(m.entities[hook.z], hook)
	m.entities[line.z] = append(m.entities[line.z], line)
}

// addFishingLine creates a fishing line with improved validation
func (m *model) addFishingLine(x int, y int, maxHeight int) *Entity {
	lineHeight := maxHeight - y

	// Create fishing line sprite
	lineFrame := strings.Repeat("|\n", lineHeight)
	lineMask := strings.Repeat("w\n", lineHeight)
	lineSprite := newSprite([]string{lineFrame}, []string{lineMask})

	line := &Entity{
		s:     lineSprite,
		x:     x + 4,
		y:     -lineHeight - y,
		z:     zWaterGap1,
		alive: true,
	}

	return line
}

// addDolphins creates dolphins with improved movement logic
func (m *model) addDolphins() {
	dir := rand.Intn(2)
	d1 := m.createDolphin(dir, 0)
	d2 := m.createDolphin(dir, 1)
	d3 := m.createDolphin(dir, 2)

	d3.onDie = func(mm *model) { mm.addSpecial() }

	m.entities[d1.z] = append(m.entities[d1.z], d1)
	m.entities[d2.z] = append(m.entities[d2.z], d2)
	m.entities[d3.z] = append(m.entities[d3.z], d3)
}

// createDolphin creates a single dolphin with improved boundary checking
func (m *model) createDolphin(dir int, num int) *Entity {
	var y int
	switch num {
	case 1:
		y = yDolphinMid
	case 2:
		y = yDolphinBottom - 1
	default:
		y = yDolphinTop
	}

	frames := art.DolphinFrames[dir]
	masks := art.DolphinMasks[dir][num]
	sprite := newSprite(frames, masks)

	// Determine starting position and movement direction
	startX, direction := m.createMovement(dir, sprite.w)

	e := &Entity{
		s:     sprite,
		x:     startX - (15 * num * direction),
		y:     y,
		z:     zWaterGap3,
		alive: true,
	}

	// Create horizontal movement behavior
	horizontalMovement := NewHorizontalMovement(dolphinVelocity, direction)
	up := NewVerticalMovement(dolphinVelocity, -1)
	down := NewVerticalMovement(dolphinVelocity, 1)
	verticalMovement := up

	e.onTick = func(mm *model, ee *Entity, dt time.Duration) {
		horizontalMovement.UpdateHorizontal(ee, dt)

		// Animate dolphin vertical movement
		switch ee.y {
		case yDolphinTop:
			verticalMovement = down
			ee.frame = 1
		case yDolphinBottom:
			verticalMovement = up
			ee.frame = 0
		}
		verticalMovement.UpdateVertical(ee, dt/3)

		// Kill dolphin when it moves completely off screen
		if direction == 1 && ee.x > 0 || direction == -1 && ee.x < mm.windowWidth {
			if IsOffScreenHorizontal(ee, mm.windowWidth) {
				ee.alive = false
			}
		}
	}

	return e
}
