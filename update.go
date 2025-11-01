package gocean

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// tickMsg represents a timer tick for animation updates
type tickMsg time.Time

// tick creates a timer command for animation updates with precise timing
func tick(tickRate time.Duration) tea.Cmd {
	return tea.Every(tickRate, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles all incoming messages and updates the model state
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleWindowResize(msg)
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	case tickMsg:
		return m.handleTick()
	}

	return m, nil
}

// handleWindowResize recreates the model with new dimensions
func (m *model) handleWindowResize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	width, height := max(defaultWidth, msg.Width), max(defaultHeight, msg.Height)
	m.windowWidth, m.windowHeight = width, height
	m.resetGame()
	return m, nil
}

// handleKeyPress processes keyboard input
func (m *model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "p":
		m.paused = !m.paused
		return m, nil
	case "r":
		m.resetGame()
		return m, nil
	}
	return m, nil
}

// handleTick processes animation frame updates
func (m *model) handleTick() (tea.Model, tea.Cmd) {
	now := time.Now()
	dt := now.Sub(m.lastTick)
	m.lastTick = now

	if !m.paused {
		m.updateSimulation(dt)
		if m.debug {
			m.addDebugText(dt)
		}
	}

	return m, tick(m.tickRate)
}

// addDebugText writes a single line of debug text to the first row
func (m *model) addDebugText(dt time.Duration) {
	fps := float64(time.Second) / float64(dt)
	text := fmt.Sprintf("FPS: %7.1f  DT: %7d us", fps, dt.Microseconds())

	row := m.grid[0]
	for i := 0; i < len(row) && i < len(text); i++ {
		row[i] = string(text[i])
	}
}

// updateSimulation runs one frame of the aquarium simulation
func (m *model) updateSimulation(dt time.Duration) {
	m.clearGrid()
	m.cullDeadEntities()
	m.updateAndRenderEntities(dt)
}

// cullDeadEntities removes dead entities from the simulation efficiently
func (m *model) cullDeadEntities() {
	m.aliveTemp = make(map[int][]*entity)
	for z, entities := range m.entities {
		for _, e := range entities {
			if e.alive {
				m.aliveTemp[z] = append(m.aliveTemp[z], e)
			}
		}
	}
	m.entities, m.aliveTemp = m.aliveTemp, m.entities
}

// updateAndRenderEntities processes all living entities for one frame
func (m *model) updateAndRenderEntities(dt time.Duration) {
	// Iterate through entity layers in reverse order
	for z := zCastle; z >= 0; z-- {
		entities := m.entities[z]
		for _, e := range entities {
			// Update entity position and state
			if e.alive && e.onTick != nil {
				e.onTick(m, e, dt)
			}

			// Handle entity death cleanup
			if !e.alive && e.onDie != nil {
				e.onDie(m)
			}

			// Render entity to the grid
			m.render(e)
		}
	}
}

// render draws an entity to the grid at its current position
func (m *model) render(e *entity) {
	frame := e.s.frames[e.frame]

	ex, ey := e.x, e.y
	h, w := m.windowHeight, m.windowWidth

	// Draw each character of the entity sprite
	for y := 0; y < e.s.h && ey+y < h; y++ {
		for x := 0; x < len(frame.image[y]) && ex+x < w; x++ {
			// Skip if position is off-screen to the left or top
			if ex+x < 0 || ey+y < 0 {
				continue
			}

			// Skip transparent pixels
			if frame.image[y][x] == " " {
				continue
			}

			// Draw the character to the grid
			m.grid[ey+y][ex+x] = frame.image[y][x]
		}
	}
}

// clearGrid resets all grid cells to spaces for the next frame
func (m *model) clearGrid() {
	for _, row := range m.grid {
		for i := range row {
			row[i] = " "
		}
	}
}
