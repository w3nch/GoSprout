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
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width/2, playerDest.Height/2), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
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
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	drawScene()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sproutling")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	rl.InitAudioDevice()
	backgroundMusic = rl.LoadMusicStream("resource/music/music.mp3") // Load music file
	rl.PlayMusicStream(backgroundMusic)
	MusicPause = false
	

	grassSprite = rl.LoadTexture("resource/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("resource/Characters/Basic Charakter Spritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 48, 48)
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
