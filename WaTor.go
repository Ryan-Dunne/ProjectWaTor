package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

const scale int = 2 //Size of entities
const width = 800
const height = 400
const fish = 1
const shark = 2
const noOfFish = 1
const noOfShark = 1

var wm [][]*creature

type creature struct {
	age, health, species int
	color                color.Color
	hasMoved             bool
}

var gridSize = []int{400, 400} //Width & Height

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}
var red color.Color = color.RGBA{255, 0, 0, 255}

var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}

func (g *Game) Update() error {
	for x := 1; x < width-1; x++ { //Across One pixel, when column complete
		for y := 1; y < height-1; y++ { //Iterates column
			if wm[x][y] == nil {
				continue
			}
			if wm[x][y].species == fish {

			}
		}
	}

	temp := buffer //Create copy of buffer(The updated grid)
	buffer = grid  //Buffer equals current grid state
	grid = temp    //temp(buffer) becomes new grid
	return nil
}

func (g *Game) Draw(window *ebiten.Image) {
	window.Fill(blue)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if wm[x][y] == nil {
				continue
			}
			if wm[x][y].species == fish { //If there's an entity
				window.Set(x, y, yellow)
			}
			if wm[x][y].species == shark {
				window.Set(x, y, red)
			}

		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game := &Game{}
	// Set up the world map as a 2-D Slice
	wm = make([][]*creature, width)
	for i := range wm {
		wm[i] = make([]*creature, height)
	}
	setUpFish()
	setUpShark()
	printGrid()

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("My Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func setUpFish() {
	for i := 0; i < noOfFish; i++ {
		spawnX := rand.Intn(width)
		spawnY := rand.Intn(height)
		if wm[spawnX][spawnY] == nil {
			wm[spawnX][spawnY] = &creature{
				age:      2,
				health:   5,
				species:  fish,
				hasMoved: false,
				color:    yellow,
			}
		}
	}
}

func setUpShark() {
	for i := 0; i < noOfShark; i++ {
		spawnX := rand.Intn(width)
		spawnY := rand.Intn(height)
		if wm[spawnX][spawnY] == nil {
			wm[spawnX][spawnY] = &creature{
				age:     2,
				health:  5,
				species: shark,
				color:   red,
			}
		}
	}
}

func printGrid() {
	fmt.Println("len(wm):", len(wm))
	if len(wm) > 0 {
		fmt.Println("len(wm[0]):", len(wm[0]))
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if wm[i][j] == nil {
				fmt.Print(" ")
			} else if wm[i][j].species == fish {
				fmt.Print("F")
			} else if wm[i][j].species == shark {
				fmt.Print("S")
			}
		}
	}
}
