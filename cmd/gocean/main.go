package main

import (
	"flag"
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

	// Debug and performance tracking
	debug     bool
	tickRate  time.Duration
	frameTime time.Duration
	fps       float64
}

// initialModel creates and initializes a new aquarium model
func initialModel(width, height int, debug bool, maxFPS int) *model {
	tickRate := time.Nanosecond
	if maxFPS > 0 {
		tickRate = time.Second / time.Duration(maxFPS)
	}

	m := &model{
		windowWidth:  width,
		windowHeight: height,
		lastTick:     time.Now(),
		paused:       false,
		entities:     make(map[int][]*Entity),
		debug:        debug,
		tickRate:     tickRate,
	}

	m.resetGame()
	return m
}

// resetGame resets the grid and all entities
func (m *model) resetGame() {
	m.entities = make(map[int][]*Entity)
	m.allocGrid()
	m.addEnvironment()
	m.addCastle()
	m.addAllSeaweed()
	m.addAllFish()
	m.addSpecial()
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
	return tea.Batch(tick(m.tickRate), tea.EnterAltScreen)
}

// main initializes and runs the aquarium application
func main() {
	debug := flag.Bool("debug", false, "Enable debug mode.")
	maxFPS := flag.Int("fps", 120, "Maximum frames per second. 0 for unlimited.")
	flag.Parse()

	p := tea.NewProgram(initialModel(defaultWidth, defaultHeight, *debug, *maxFPS))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
