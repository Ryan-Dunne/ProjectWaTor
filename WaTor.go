package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

const scale int = 5 //Size of entities
const width = 600
const height = 600

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}
var red color.Color = color.RGBA{255, 0, 0, 255}

var NumShark = 80   //Starting Shark Population
var NumFish = 150   //Starting Fish Population
var FishBreed = 10  //Num of chronons that must pass before fish can breed
var SharkBreed = 12 //Num of chronons that must pass before sharks can breed
var Starve = 15     //Num of time sharks can go without food before death
// var GridSize		//Dimensions of world
var Threads = 1 //Num of threads to use

var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}
var count int = 0

func Chronon() {
	time.Sleep(5 * time.Millisecond)
}

func (g *Game) Update() error {
	Chronon()
	for x := 1; x < width-1; x++ { //Across One pixel, when column complete
		for y := 1; y < height-1; y++ { //Iterates column
			buffer[x][y] = 0 //Buffer defaults to no entity

			n := grid[x-1][y+0] + //Checks the neighbourhood - W,S,N,E
				grid[x+0][y-1] +
				grid[x+0][y+1] +
				grid[x+1][y+0]

			if grid[x][y] == 0 && n == 3 { //If tile empty & 3 entities surround
				buffer[x][y] = 1 //Create new entity
			} else if n < 2 || n > 3 { //If there's too few or too many entities, leave current cell blank
				buffer[x][y] = 0
			} else {
				buffer[x][y] = grid[x][y]
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
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[x][y] == 1 { //If there's an entity
						window.Set(x*scale+i, y*scale+j, red) //Change pixel colour
					}
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("My Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func setupFish() {
	for i := 0; i < NumFish; i++ {
		rand.Int() //Generate random coords to spawn fishs
	}
}
