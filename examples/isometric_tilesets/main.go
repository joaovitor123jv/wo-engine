package main

import (
	"fmt"

	woengine "github.com/joaovitor123jv/wo-engine"
	woutils "github.com/joaovitor123jv/wo-engine/wo-utils"
	"github.com/veandco/go-sdl2/sdl"
)

// The rendered map is a 2D isometric map crafted using the Tiled Map Editor.
// 		The tiles on the layers are drawn from the tilesets, and each tile has a size of 128x64.
// 		The rendering is based on an isometric projection.
// 		The tile data is stored in a CSV format within the TMX file.
// 		Please note that this setup works best with the current Tiled Map Editor configurations and may not support other configurations or editors.

func gameLogic() {
	// Create a new graphics context for rendering the isometric tilemap
	context := woutils.NewContext("Isometric Tilemap Rendering Example")
	defer context.Destroy() // Ensure resources are cleaned up when the function exits

	// Set the FPS target to 100
	context.SetTargetFramerate(100)

	// Start the rendering context
	context.Start()

	fpsViewer := woutils.NewFramerateViewer()

	// Variables for tracking user interactions such as zoom and map movement
	isApplyingZoom := false
	zoomSourceY := int32(0)
	isMovingMap := false
	defaultZoom := float32(1)
	movementSourceX, movementSourceY := int32(0), int32(0)

	// Initialize the game map which is a 2D isometric tilemap using the specified TMX file
	gameMap := woutils.NewGameMap(&context, "Test Map", "assets/test.tmx")
	defer gameMap.Destroy() // Ensure the game map is cleaned up when no longer needed

	// Add the game map as a renderable entity within the context
	context.AddRenderable(&gameMap)
	context.AddRenderable(&fpsViewer)

	// Set up a mouse click listener to handle map movements and zoom functionality
	context.AddMouseClickListener(func(x, y int32, button uint8, isPressed bool) bool {
		if button == sdl.BUTTON_LEFT { // Check for left mouse button click
			if isPressed {
				isMovingMap = true                      // Start moving the map
				movementSourceX, movementSourceY = x, y // Store the initial click position
			} else {
				isMovingMap = false // Stop moving the map
			}
			return true
		} else if button == sdl.BUTTON_RIGHT { // Check for right mouse button click
			if isPressed {
				isApplyingZoom = true // Start applying zoom
				zoomSourceY = y       // Store the initial zoom position
			} else {
				isApplyingZoom = false // Stop applying zoom
			}
			return true
		} else if button == sdl.BUTTON_MIDDLE { // Check for middle mouse button click
			defaultZoom = 1.0                   // Reset zoom to default value
			context.Camera.SetZoom(defaultZoom) // Apply the default zoom to the game map
		}
		return false
	})

	// Set up a mouse movement listener to handle map dragging and zoom adjustment
	context.AddMouseMovementListener(func(x, y int32) bool {
		if isMovingMap { // If map is in moving state
			context.Camera.Translate(x-movementSourceX, y-movementSourceY) // Translate map based on mouse movement
			movementSourceX, movementSourceY = x, y                        // Update movement source coordinates
			return true
		}
		if isApplyingZoom { // If zoom is being applied
			defaultZoom += 0.01 * float32(zoomSourceY-y) // Adjust zoom based on mouse movement
			context.Camera.SetZoom(defaultZoom)          // Apply the calculated zoom level to the game map
			fmt.Println("Current zoom: ", defaultZoom)
		}
		return false
	})

	// Enter the main event loop to process events and render graphics
	context.MainLoop()
}

func main() {
	game := woengine.NewGame()    // Create a new game instance
	game.SetEntrypoint(gameLogic) // Set the entry point function for the game
	game.Run()                    // Start the game
}
