package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	Red             = color.RGBA{0xff, 0x0, 0x0, 0xff}
	Yellow          = color.RGBA{0xff, 0xff, 0x0, 0xff}
	Blue            = color.RGBA{0x00, 0x00, 0xff, 0xff}
	mplusNormalFont font.Face
	groupYellow     []*Atom
	groupBlue       []*Atom
	groupRed        []*Atom
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomX() int {
	return rand.Intn(screenWidth)
}
func randomY() int {
	return rand.Intn(screenHeight)
}

type Atom struct {
	x, y, vx, vy float64
	color        color.Color
}

func (g *Game) createGroup(num int, color color.Color) []*Atom {
	var group []*Atom
	for i := 0; i < num; i++ {
		a := &Atom{x: float64(randomX()), y: float64(randomY()), color: color}
		g.atoms = append(g.atoms, a)
		group = append(group, a)
	}
	return group
}

func (g *Game) draw(screen *ebiten.Image) {
	for _, a := range g.atoms {
		ebitenutil.DrawRect(screen, a.x, a.y, 5, 5, a.color)
	}
}

func rule(group1, group2 []*Atom, g float64) {
	for _, a := range group1 {
		var fx, fy float64
		for _, b := range group2 {
			dx := a.x - b.x
			dy := a.y - b.y
			d := math.Sqrt(dx*dx + dy*dy)
			if d > 0 && d < 50 {
				force := (g * 1) / d
				fx += force * dx
				fy += force * dy
			}
		}
		a.vx = (a.vx + fx) * 0.5
		a.vy = (a.vy + fy) * 0.5
		a.x += a.vx
		a.y += a.vy
		if a.x <= 10 || a.x >= screenWidth-10 {
			a.vx *= -1
		}
		if a.y <= 10 || a.y >= screenHeight-10 {
			a.vy *= -1
		}
	}
}

type Game struct {
	counter int

	atoms []*Atom
}

func (g *Game) Update() error {
	// Change the text color for each second.

	if g.counter%ebiten.TPS() == 0 {

	}

	rule(groupBlue, groupRed, -0.3)
	rule(groupRed, groupBlue, -0.3)
	rule(groupRed, groupRed, 0.03)
	rule(groupYellow, groupRed, 0.20)

	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.RGBA{0x99, 0xcc, 0xff, 0xff})
	screen.Fill(color.Black)

	g.draw(screen)
	ebitenutil.DrawRect(screen, screenWidth, screenHeight, 10, 10, color.Opaque)
	// Draw info
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	text.Draw(screen, msg, mplusNormalFont, 20, 40, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth * 2, screenHeight * 2
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Live")
	g := &Game{}
	groupYellow = g.createGroup(400, Yellow)
	groupRed = g.createGroup(200, Red)
	groupBlue = g.createGroup(200, Blue)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
