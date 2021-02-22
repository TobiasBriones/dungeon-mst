package model

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"math"
)

const (
	PathWidthPx            = 16
	horizontalUnitWidthPx  = 64
	horizontalUnitHeightPx = 12
	verticalUnitWidthPx    = horizontalUnitHeightPx
	verticalUnitHeightPx   = horizontalUnitWidthPx
)

type Dungeon struct {
	rect         Rect
	factor       DimensionFactor
	brickImage   *ebiten.Image
	brickYImage  *ebiten.Image
	neighborhood []*Dungeon
}

func (d *Dungeon) Width() int {
	return d.rect.Width()
}

func (d *Dungeon) Height() int {
	return d.rect.Height()
}

func (d *Dungeon) Cx() int {
	return d.rect.Cx()
}

func (d *Dungeon) Cy() int {
	return d.rect.Cy()
}

func (d *Dungeon) Center() Point {
	return Point{
		X: d.Cx(),
		Y: d.Cy(),
	}
}

func (d *Dungeon) Overlaps(other *Dungeon, margin int) bool {
	xo := (d.rect.Left-margin) <= (other.rect.Right+margin) &&
		(d.rect.Right+margin) >= (other.rect.Left-margin)

	yo := (d.rect.Top-margin) <= (other.rect.Bottom+margin) &&
		(d.rect.Bottom+margin) >= (other.rect.Top-margin)

	return xo && yo
}

func (d *Dungeon) Intersects(rect *Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Dungeon) AddNeighbor(dungeon *Dungeon) {
	d.neighborhood = append(d.neighborhood, dungeon)
}

func (d *Dungeon) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	wFactor := d.factor.Width
	hFactor := d.factor.Height
	blockWidth := horizontalUnitHeightPx

	// Draw Top
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Top))
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(d.brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Bottom
	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Bottom-blockWidth))
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(d.brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Left
	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Top))
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(d.brickYImage, op)
		op.GeoM.Translate(0, verticalUnitHeightPx)
	}

	// Draw Right
	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Right-blockWidth), float64(d.rect.Top))
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(d.brickYImage, op)
		op.GeoM.Translate(0, verticalUnitHeightPx)
	}

	// Draw the Neighborhood
	d.drawNeighborhood(screen)
}

func (d *Dungeon) drawNeighborhood(screen *ebiten.Image) {
	for _, neighbor := range d.neighborhood {
		center := d.Center()
		sw := PathWidthPx / 2
		path := pathTrace{
			p00: Point{center.X - sw, center.Y + sw},
			p01: Point{center.X - sw, neighbor.Cy() + sw},
			p10: Point{center.X - sw, neighbor.Cy() - sw},
			p11: Point{neighbor.Cx() - sw, neighbor.Cy() - sw},
		}
		d.drawPath(path, screen)
	}
}

func (d *Dungeon) drawPath(path pathTrace, screen *ebiten.Image) {
	d.drawPathLine(path.p00, path.p01, screen)
	d.drawPathLine(path.p10, path.p11, screen)
}

func (d *Dungeon) drawPathLine(p1 Point, p2 Point, screen *ebiten.Image) {
	if p1.X == p2.X {
		x := p1.X
		y0 := p1.Y
		y1 := p2.Y
		d.drawVerticalPath(x, y0, y1, screen)
	} else {
		y := p1.Y
		x0 := p1.X
		x1 := p2.X
		d.drawHorizontalPath(y, x0, x1, screen)
	}
}

func (d *Dungeon) drawHorizontalPath(y int, x0 int, x1 int, screen *ebiten.Image) {
	x := min(x0, x1)
	w := int(math.Abs(float64(x0 - x1)))
	line := ebiten.NewImage(w, PathWidthPx)
	op := &ebiten.DrawImageOptions{}

	line.Fill(color.Gray{})
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(line, op)
}

func (d *Dungeon) drawVerticalPath(x int, y0 int, y1 int, screen *ebiten.Image) {
	y := min(y0, y1)
	h := int(math.Abs(float64(y0 - y1)))
	line := ebiten.NewImage(PathWidthPx, h)
	op := &ebiten.DrawImageOptions{}

	line.Fill(color.Gray{})
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(line, op)
}

func NewDungeon(p0 Point, factor DimensionFactor) Dungeon {
	brickImg, _, err := ebitenutil.NewImageFromFile("./assets/brick.png")

	if err != nil {
		log.Fatal(err)
	}
	brickYImg, _, err := ebitenutil.NewImageFromFile("./assets/brick_y.png")

	if err != nil {
		log.Fatal(err)
	}
	x0 := p0.X
	y0 := p0.Y
	w := factor.Width * horizontalUnitWidthPx
	h := factor.Height * verticalUnitHeightPx
	rect := Rect{x0, y0, x0 + w, y0 + h}
	return Dungeon{
		rect,
		factor,
		brickImg,
		brickYImg,
		[]*Dungeon{},
	}
}

type DimensionFactor struct {
	Width  int
	Height int
}

type pathTrace struct {
	p00 Point
	p01 Point
	p10 Point
	p11 Point
}

func GetDungeonHorizontalUnitSize() Dimension {
	return Dimension{
		Width:  horizontalUnitWidthPx,
		Height: horizontalUnitHeightPx,
	}
}

func GetDungeonVerticalUnitSize() Dimension {
	return Dimension{
		Width:  verticalUnitWidthPx,
		Height: verticalUnitHeightPx,
	}
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
