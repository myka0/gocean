package gocean

import (
	"math/rand"
	"strings"
	"time"

	"github.com/myka0/gocean/internal/art"
)

// Special entity types
type specialEntityType int

const (
	ship specialEntityType = iota
	whale
	monster
	bigFish
	shark
	fishHook
	swan
	ducks
	dolphins
	specialEntityCount
)

// addSpecial creates special entities using improved selection and configuration
func (m *model) addSpecial() {
	specialType := specialEntityType(rand.Intn(int(specialEntityCount)))

	switch specialType {
	case ship:
		m.addShip()
	case whale:
		m.addWhale()
	case monster:
		m.addMonster()
	case bigFish:
		m.addBigFish()
	case shark:
		m.addShark()
	case fishHook:
		m.addFishHook()
	case swan:
		m.addSwan()
	case ducks:
		m.addDucks()
	case dolphins:
		m.addDolphins()
	}
}

// createHorizontalEntity creates a basic horizontal-moving entity with common setup
func (m *model) createHorizontalEntity(dir int, y int, z int, velocity int, sprite sprite) *entity {
	startX, direction := m.createMovement(dir, sprite.w)

	e := &entity{
		s:     sprite,
		x:     startX,
		y:     y,
		z:     z,
		alive: true,
	}

	horizontalMovement := newHorizontalMovement(velocity, direction)

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		horizontalMovement.updateHorizontal(ee, dt)

		if isOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	e.onDie = func(mm *model) { mm.addSpecial() }
	return e
}

// createAnimatedHorizontalEntity creates an animated horizontal-moving entity
func (m *model) createAnimatedHorizontalEntity(dir int, y int, z int, velocity int, frameDelay time.Duration, sprite sprite) *entity {
	startX, direction := m.createMovement(dir, sprite.w)

	e := &entity{
		s:          sprite,
		x:          startX,
		y:          y,
		z:          z,
		frameDelay: frameDelay,
		alive:      true,
	}

	horizontalMovement := newHorizontalMovement(velocity, direction)

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		ee.AdvanceFrame()
		horizontalMovement.updateHorizontal(ee, dt)

		if isOffScreenHorizontal(ee, mm.windowWidth) {
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

	e := &entity{
		s:     sprite,
		x:     startX,
		y:     yWhalePosition,
		z:     zWaterGap2,
		alive: true,
	}

	horizontalMovement := newHorizontalMovement(whaleVelocity, direction)

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		horizontalMovement.updateHorizontal(ee, dt)

		// Kill whale when it moves completely off screen
		if isOffScreenHorizontal(ee, mm.windowWidth) {
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
func (m *model) addWaterSpout(x int, y int, dir int) *entity {
	// Create water spout sprite
	mask := make([]string, len(art.WaterSpoutFrames))
	for i := range mask {
		mask[i] = art.WaterSpoutMask
	}
	sprite := newSprite(art.WaterSpoutFrames, mask)

	e := &entity{
		s:          sprite,
		x:          x,
		y:          y,
		z:          zWaterGap2,
		frameDelay: waterSpoutFrameDelay,
		alive:      true,
	}

	// Create horizontal movement behavior
	horizontalMovement := newHorizontalMovement(whaleVelocity, dir)

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		// Create delay when spout resets
		if ee.frame == 0 {
			ee.frameDelay = waterSpoutResetDelay
		}
		ee.AdvanceFrame()
		if ee.frame == 1 {
			ee.frameDelay = waterSpoutFrameDelay
		}

		horizontalMovement.updateHorizontal(ee, dt)
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

	e := &entity{
		s:     sprite,
		x:     startX,
		y:     rand.Intn(m.windowHeight-waterSurfaceOffset-sprite.h) + waterSurfaceOffset,
		z:     zShark,
		alive: true,
	}

	// Create horizontal movement behavior
	horizontalMovement := newHorizontalMovement(sharkVelocity, direction)

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		horizontalMovement.updateHorizontal(ee, dt)

		// Calculate teeth position with improved logic
		teethX := ee.x + 50
		if direction == -1 {
			teethX = ee.x + 3
		}
		teethY := ee.y + 8

		// Use collision detector for better performance
		collisionDetector := newCollisionDetector(teethX, teethY)
		if fish := collisionDetector.checkCollisionWithFish(mm); fish != nil {
			fish.alive = false
			mm.addSplat(teethX, teethY)
			mm.addFish()
		}

		if isOffScreenHorizontal(ee, mm.windowWidth) {
			ee.alive = false
		}
	}

	e.onDie = func(mm *model) { mm.addSpecial() }
	m.entities[e.z] = append(m.entities[e.z], e)
}

// addSplat creates a splat effect using entity pooling
func (m *model) addSplat(x, y int) {
	splat := newSprite(art.SplatFrames, art.SplatMasks)

	e := &entity{
		s:          splat,
		x:          x - 3,
		y:          y - 3,
		z:          zShark - 1,
		frameDelay: splatFrameDelay,
		alive:      true,
	}

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
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

	hook := &entity{
		s:     sprite,
		x:     x,
		y:     -sprite.h,
		z:     zWaterGap1,
		alive: true,
	}

	// Create vertical movement behavior
	up := newVerticalMovement(fishHookVelocity, -1)
	down := newVerticalMovement(fishHookVelocity, 1)
	fishMovement := newVerticalMovement(fishHookVelocity, -1)
	verticalMovement := down

	// Create fishing line and movement behavior
	line := m.addFishingLine(x, sprite.h, maxHeight)
	lineUp := newVerticalMovement(fishHookVelocity, -1)
	lineDown := newVerticalMovement(fishHookVelocity, 1)
	lineMovement := lineDown

	hookedFish := false

	hook.onTick = func(mm *model, ee *entity, dt time.Duration) {
		if hookedFish || ee.y+ee.s.h < maxHeight {
			verticalMovement.updateVertical(ee, dt)
			lineMovement.updateVertical(line, dt)
		}

		pointX := ee.x + 1
		pointY := ee.y + 3

		// Use collision detector for fish hook
		if !hookedFish {
			collisionDetector := newCollisionDetector(pointX, pointY)
			if fish := collisionDetector.checkCollisionWithFish(mm); fish != nil {
				hookedFish = true
				verticalMovement = up
				lineMovement = lineUp
				fish.onTick = func(mm *model, ee *entity, dt time.Duration) {
					fishMovement.updateVertical(fish, dt)
					if isOffScreenVertical(ee, mm.windowHeight) {
						ee.alive = false
					}
				}
			}
		}

		// Kill fish hook when it moves completely off screen
		if isOffScreenVertical(ee, mm.windowHeight) {
			ee.alive = false
			line.alive = false
		}
	}

	hook.onDie = func(mm *model) { mm.addSpecial() }
	m.entities[hook.z] = append(m.entities[hook.z], hook)
	m.entities[line.z] = append(m.entities[line.z], line)
}

// addFishingLine creates a fishing line with improved validation
func (m *model) addFishingLine(x int, y int, maxHeight int) *entity {
	lineHeight := maxHeight - y

	// Create fishing line sprite
	lineFrame := strings.Repeat("|\n", lineHeight)
	lineMask := strings.Repeat("w\n", lineHeight)
	lineSprite := newSprite([]string{lineFrame}, []string{lineMask})

	line := &entity{
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
func (m *model) createDolphin(dir int, num int) *entity {
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

	e := &entity{
		s:     sprite,
		x:     startX - (15 * num * direction),
		y:     y,
		z:     zWaterGap3,
		alive: true,
	}

	// Create horizontal movement behavior
	horizontalMovement := newHorizontalMovement(dolphinVelocity, direction)
	up := newVerticalMovement(dolphinVelocity, -1)
	down := newVerticalMovement(dolphinVelocity, 1)
	verticalMovement := up

	e.onTick = func(mm *model, ee *entity, dt time.Duration) {
		horizontalMovement.updateHorizontal(ee, dt)

		// Animate dolphin vertical movement
		switch ee.y {
		case yDolphinTop:
			verticalMovement = down
			ee.frame = 1
		case yDolphinBottom:
			verticalMovement = up
			ee.frame = 0
		}
		verticalMovement.updateVertical(ee, dt/3)

		// Kill dolphin when it moves completely off screen
		if direction == 1 && ee.x > 0 || direction == -1 && ee.x < mm.windowWidth {
			if isOffScreenHorizontal(ee, mm.windowWidth) {
				ee.alive = false
			}
		}
	}

	return e
}
