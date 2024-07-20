package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth      = 1000
	screenHeight     = 480
	fullScreenWidth  = 1920 // Desired fullscreen width
	fullScreenHeight = 1080 // Desired fullscreen height
)

var (
	running         = true
	bkgColor        = rl.NewColor(147, 211, 196, 255)
	grassSprite     rl.Texture2D
	playerSprite    rl.Texture2D
	playerSrc       rl.Rectangle
	playerDest      rl.Rectangle
	playerSpeed     float32 = 3
	isFullScreen    bool
	backgroundMusic rl.Music
	MusicPause      bool
	cam             rl.Camera2D
	playerDir       int
	playerMoving    bool
	playerUp        bool
	playerRight     bool
	playerDown      bool
	playerLeft      bool
	playerFrame     int
	frameCount      int
)

func drawScene() {
	// Select the correct frame for the player based on direction
	playerSrc.X = float32(playerFrame) * playerSrc.Width
	playerSrc.Y = float32(playerDir) * playerSrc.Height
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width/2, playerDest.Height/2), 0, rl.White)
}

func input() {
	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerUp = true
		playerMoving = true
		playerDir = 1
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDown = true
		playerMoving = true
		playerDir = 0
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerLeft = true
		playerMoving = true
		playerDir = 2
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerRight = true
		playerMoving = true
		playerDir = 3
	}

	// Toggle fullscreen mode
	if rl.IsKeyPressed(rl.KeyF) {
		isFullScreen = !isFullScreen
		if isFullScreen {
			rl.ToggleFullscreen()
			rl.SetWindowSize(fullScreenWidth, fullScreenHeight)
		} else {
			rl.ToggleFullscreen()
			rl.SetWindowSize(screenWidth, screenHeight)
		}
	}

	// Exit the application
	if rl.IsKeyPressed(rl.KeyEscape) {
		running = false
	}

	// Pause/Resume music
	if rl.IsKeyPressed(rl.KeyQ) {
		MusicPause = !MusicPause
	}
}

func update() {
	running = !rl.WindowShouldClose()
	rl.UpdateMusicStream(backgroundMusic) // Update the music stream
	if MusicPause {
		rl.PauseMusicStream(backgroundMusic)
	} else {
		rl.ResumeMusicStream(backgroundMusic)
	}

	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		frameCount++
		if frameCount%8 == 0 {
			playerFrame++
			if playerFrame > 3 {
				playerFrame = 0
			}
		}
	} else {
		playerFrame = 0
	}

	// Update camera target to follow the player
	cam.Target = rl.NewVector2(
		playerDest.X+playerDest.Width/2,
		playerDest.Y+playerDest.Height/2,
	)
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)

	rl.BeginMode2D(cam)
	drawScene()
	rl.EndMode2D()

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sproutling")
	rl.SetExitKey(0)
	rl.SetTargetFPS(120) // Max FPS set to 120

	rl.InitAudioDevice()
	backgroundMusic = rl.LoadMusicStream("resource/music/music.mp3") // Load music file
	rl.PlayMusicStream(backgroundMusic)
	MusicPause = false

	grassSprite = rl.LoadTexture("resource/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("resource/Characters/Basic Charakter Spritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 48, 48)

	playerDir = 0 // Initialize playerDir

	cam = rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),  // camera position
		rl.NewVector2(float32(playerDest.X+playerDest.Width/2), float32(playerDest.Y+playerDest.Height/2)), // target position
		0, // rotation
		2, // zoom level increased to 2
	)
}

func quit() {
	rl.UnloadMusicStream(backgroundMusic) // Unload music stream
	rl.CloseAudioDevice()                 // Close audio device

	rl.UnloadTexture(playerSprite)
	rl.UnloadTexture(grassSprite)
	rl.CloseWindow()
}

func main() {
	for running {
		input()
		update()
		render()
	}
	quit()
}
