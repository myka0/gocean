package gocean

import (
	"fmt"
	"math/rand"
	"unsafe"

	"github.com/charmbracelet/lipgloss"
)

// View renders the current state to a string for display
func (m *model) View() string {
	// Precalculate total buffer size
	total := 0
	for _, row := range m.grid {
		for _, s := range row {
			total += len(s)
		}
		total++ // newline character
	}

	// Reuse buffer if possible, otherwise allocate new one
	if cap(m.renderBuf) < total {
		m.renderBuf = make([]byte, total)
	} else {
		m.renderBuf = m.renderBuf[:total]
	}

	// Build the output buffer
	pos := copy(m.renderBuf, "\n")
	for _, row := range m.grid {
		for _, s := range row {
			pos += copy(m.renderBuf[pos:], s)
		}
		pos += copy(m.renderBuf[pos:], "\n")
	}

	// Zero-copy string creation using unsafe
	return unsafe.String(&m.renderBuf[0], len(m.renderBuf))
}

// createPalette generates a color palette for entity rendering
func createPalette() map[byte]lipgloss.Style {
	base := lipgloss.NewStyle().Bold(true)
	styles := map[byte]lipgloss.Style{}

	// Create randomized colors for digits 1-7, with 4 fixed as white for eye highlights
	for d := '1'; d <= '7'; d++ {
		key := byte(d)
		if d == '4' {
			styles[key] = base.Foreground(lipgloss.Color("7"))
			continue
		}
		styles[key] = base.Foreground(lipgloss.Color(fmt.Sprint(rand.Intn(6) + 1)))
	}

	return styles
}

// colorize applies color styling to a character based on its color mask
func colorize(c byte, m byte, palette map[byte]lipgloss.Style) string {
	base := lipgloss.NewStyle().Bold(true)
	// Return uncolored space characters as is
	if c == ' ' && m == ' ' {
		return string(c)
	}

	// Use palette colors for digits 1-7
	if m >= '1' && m <= '7' {
		return palette[m].Render(string(c))
	}

	// Use ANSI color mapping for other color codes
	colID, ok := ansi16[m]
	if !ok {
		return base.Foreground(lipgloss.Color("7")).Render(string(c))
	}
	return base.Foreground(lipgloss.Color(colID)).Render(string(c))
}
