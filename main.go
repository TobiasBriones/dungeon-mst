package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

var (
	bgImage     *ebiten.Image
	runnerImage *ebiten.Image
)

type Game struct {
	count  int
	runner Runner
}

func (g *Game) Update() error {
	g.count++

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			switch k {
			case ebiten.KeyUp, ebiten.KeyW:
				g.runner.Pos.Y--
			case ebiten.KeyDown, ebiten.KeyS:
				g.runner.Pos.Y++
			case ebiten.KeyLeft, ebiten.KeyA:
				g.runner.Pos.X--
			case ebiten.KeyRight, ebiten.KeyD:
				g.runner.Pos.X++
			}
		}
	}
	g.runner.Normalize()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBg(screen)
	drawSomeDungeons(screen)
	drawRunner(g, screen)
}

func NewGame() Game {
	runner := NewRunner()
	return Game{
		count:  0,
		runner: runner,
	}
}

type Rect struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

type Dungeon struct {
	rect        Rect
	brickImage  *ebiten.Image
	brickYImage *ebiten.Image
}

func (d *Dungeon) Width() int {
	return d.rect.Right - d.rect.Left
}

func (d *Dungeon) Height() int {
	return d.rect.Bottom - d.rect.Top
}

func (d *Dungeon) Cx() int {
	return d.rect.Left + d.Width()/2
}

func (d *Dungeon) Cy() int {
	return d.rect.Top + d.Height()/2
}

func (d *Dungeon) Overlaps(other Dungeon, margin int) bool {
	xo := (d.rect.Left-margin) <= (other.rect.Right+margin) &&
		(d.rect.Right+margin) >= (other.rect.Left-margin)

	yo := (d.rect.Top-margin) <= (other.rect.Bottom+margin) &&
		(d.rect.Bottom+margin) >= (other.rect.Top-margin)

	return xo && yo
}

func (d *Dungeon) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	sizeX := d.Width() / 128
	sizeY := d.Height() / 128
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Top))

	for i := 0; i < sizeX; i++ {
		screen.DrawImage(d.brickImage, op)
		op.GeoM.Translate(128, 0)
	}

	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Bottom-24))
	for i := 0; i < sizeX; i++ {
		screen.DrawImage(d.brickImage, op)
		op.GeoM.Translate(128, 0)
	}

	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Top))
	for i := 0; i < sizeY; i++ {
		screen.DrawImage(d.brickYImage, op)
		op.GeoM.Translate(0, 128)
	}

	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Right-24), float64(d.rect.Top))
	for i := 0; i < sizeY; i++ {
		screen.DrawImage(d.brickYImage, op)
		op.GeoM.Translate(0, 128)
	}
}

func NewDungeon(x0 int, y0 int, sizeX int, sizeY int) Dungeon {
	brickImg, _, err := ebitenutil.NewImageFromFile("./assets/brick.png")

	if err != nil {
		log.Fatal(err)
	}
	brickYImg, _, err := ebitenutil.NewImageFromFile("./assets/brick_y.png")

	if err != nil {
		log.Fatal(err)
	}
	rect := Rect{x0, y0, x0 + sizeX*128, y0 + sizeY*128}
	return Dungeon{rect, brickImg, brickYImg}
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

type Runner struct {
	Pos   image.Point
	Scale float64
}

func (r *Runner) Center() {
	r.Pos.X = int(-(frameWidth * r.Scale) / 2)
	r.Pos.Y = int(-(frameHeight * r.Scale) / 2)
	r.Pos.X += screenWidth / 2
	r.Pos.Y += screenHeight / 2
}

func (r *Runner) Normalize() {
	pos := &r.Pos

	if pos.X < 0 {
		pos.X = 0
	}
	if pos.X > screenWidth-int(frameWidth*r.Scale) {
		pos.X = screenWidth - int(frameWidth*r.Scale)
	}
	if pos.Y < 0 {
		pos.Y = 0
	}
	if pos.Y > screenHeight-int(frameHeight*r.Scale) {
		pos.Y = screenHeight - int(frameHeight*r.Scale)
	}
}

func NewRunner() Runner {
	runner := Runner{
		Pos:   image.Point{},
		Scale: 2,
	}
	runner.Center()
	return runner
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dungeon MST")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func init() {
	loadBg()
	loadRunner()
}

func loadBg() {
	rand.Seed(time.Now().UnixNano())
	bgNumber := rand.Intn(3) + 1
	bgName := "bg_" + strconv.Itoa(bgNumber) + ".png"
	bgImg, _, err := ebitenutil.NewImageFromFile("./assets/" + bgName)

	if err != nil {
		log.Fatal(err)
	}
	bgImage = bgImg
}

func loadRunner() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))

	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)
}

func drawBg(screen *ebiten.Image) {
	screen.DrawImage(bgImage, nil)
}

func drawRunner(g *Game, screen *ebiten.Image) {
	runner := g.runner
	pos := runner.Pos
	op := &ebiten.DrawImageOptions{}
	i := (g.count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	rect := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)

	op.GeoM.Scale(runner.Scale, runner.Scale)
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))
	screen.DrawImage(runnerImage.SubImage(rect).(*ebiten.Image), op)
}

func drawSomeDungeons(screen *ebiten.Image) {
	dungeons := genSomeDungeons()

	for _, dungeon := range dungeons {
		dungeon.Draw(screen)
	}
}

func genSomeDungeons() []Dungeon {
	return []Dungeon{
		NewDungeon(0, 0, 1, 1),
		NewDungeon(20, 540, 4, 1),
		NewDungeon(200, 140, 3, 2),
		NewDungeon(350, 90, 4, 1),
	}
}
