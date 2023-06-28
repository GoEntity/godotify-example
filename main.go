//go get github.com/GoEntity/godotify
//go get github.com/hajimehoshi/ebiten

//player_godotified.png is already generated in `img`` folder because I ran `go run main.go` to convert player.png as an example
//customize your path to save img files

package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"

	"log"
	"os"

	"github.com/GoEntity/godotify"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	player *ebiten.Image
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1900, 1400
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screenCenterX := float64(screen.Bounds().Dx()) / 2
	screenCenterY := float64(screen.Bounds().Dy()) / 2

	imageCenterX := float64(g.player.Bounds().Dx()) / 2
	imageCenterY := float64(g.player.Bounds().Dy()) / 2

	op.GeoM.Translate(screenCenterX-imageCenterX, screenCenterY-imageCenterY)

	screen.DrawImage(g.player, op)
}

func (g *Game) init() {
	var err error
	g.player, _, err = ebitenutil.NewImageFromFile("img/player_godotified.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func main() {
	//set up for GoDotify
	//of course, you can set path however you want it
	inputFile := "img/player.png"
	outputFile := "img/player_godotified.png"
	config := godotify.Config{
		Intensity: 0.7,
	}

	//godotify it (intake ur img, process it, and spit it out)
	err := godotify.GoDotify(inputFile, outputFile, config)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	//open result img
	file, err := os.Open(outputFile)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer file.Close()

	//decode
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	//convert to ebiten image
	yourImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	//now you can use `yourImage` as an ebiten.Image in your game

	game := &Game{}
	game.player = yourImage
	game.init()

	ebiten.SetWindowSize(1900, 1400)
	ebiten.SetWindowTitle("godotify example.. show player")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
