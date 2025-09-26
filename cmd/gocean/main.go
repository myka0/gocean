package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
)

// tickMsg represents a timer tick for animation updates
type tickMsg time.Time

// model holds the application state for the aquarium simulation
type model struct {
	// Display dimensions
	windowHeight int
	windowWidth  int

	// Rendering and simulation state
	grid     [][]string
	paused   bool
	lastTick time.Time
	entities map[int][]*Entity

	// Performance optimization buffers
	renderBuf []byte
	aliveTemp map[int][]*Entity

	// Frame timing control
	lastFrame time.Time
}

// initialModel creates and initializes a new aquarium model
func initialModel(width, height int) *model {
	m := &model{
		windowWidth:  width,
		windowHeight: height,
		lastTick:     time.Now(),
		paused:       false,
		entities:     make(map[int][]*Entity),
		lastFrame:    time.Now(),
	}

	// Initialize the rendering grid
	m.allocGrid()

	// Create environment and creatures
	m.addEnvironment()
	m.addCastle()
	m.addAllSeaweed()
	m.addAllFish()
	m.addSpecial()

	return m
}

// allocGrid initializes the 2D rendering grid
func (m *model) allocGrid() {
	m.grid = make([][]string, m.windowHeight)
	for y := range m.grid {
		m.grid[y] = make([]string, m.windowWidth)
	}
}

// Init implements the Bubble Tea interface for initialization
func (m *model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}

// main initializes and runs the aquarium application
func main() {
	p := tea.NewProgram(initialModel(defaultWidth, defaultHeight))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
