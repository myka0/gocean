package main

import (
	"strings"
)

// Sprite represents a visual element that can have multiple animation frames
type Sprite struct {
	frames []Frame
	w, h   int
}

// Frame represents a single animation frame of a sprite
type Frame struct {
	image [][]string
}

// getWidth calculates the maximum width of any line in a multi-line string
func getWidth(s string) int {
	lines := strings.Split(s, "\n")
	maxWidth := 0

	// Find the longest line to determine sprite width
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	return maxWidth
}

// getHeight calculates the number of lines in a multi-line string
func getHeight(s string) int {
	lines := strings.Split(s, "\n")
	return len(lines)
}

// newSprite creates a Sprite from raw ASCII art and color mask data
func newSprite(rawFrames []string, rawMasks []string) Sprite {
	var frames []Frame

	// Calculate sprite dimensions from the first frame
	w := getWidth(rawFrames[0])
	h := getHeight(rawFrames[0])

	// Get the color palette for character colorization
	palette := createPalette()

	// Process each animation frame
	for f := range rawFrames {
		// Initialize frame with proper dimensions
		frames = append(frames, Frame{
			image: make([][]string, h),
		})

		frameLines := strings.Split(rawFrames[f], "\n")
		maskLines := strings.Split(rawMasks[f], "\n")

		for y := range len(frameLines) {
			for x := range len(frameLines[y]) {
				char := frameLines[y][x]

				// Convert 'b' to backtick
				if char == 'b' {
					char = '`'
				}

				// Get the color mask character
				var maskChar byte = '?'
				if x < len(maskLines[y]) {
					maskChar = maskLines[y][x]
				}

				// Apply color to the character and add to frame
				coloredChar := colorize(char, maskChar, palette)
				frames[f].image[y] = append(frames[f].image[y], coloredChar)
			}
		}
	}

	return Sprite{
		frames: frames,
		w:      w,
		h:      h,
	}
}
