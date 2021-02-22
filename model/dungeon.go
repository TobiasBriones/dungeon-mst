package model

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"math"
)

const (
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
		w1 := int(math.Abs(float64(center.X-neighbor.Cx()))) + 1
		h1 := 12
		w2 := 12
		h2 := int(math.Abs(float64(center.Y-neighbor.Cy()))) + 1
		line1 := ebiten.NewImage(w1, h1)
		line2 := ebiten.NewImage(w2, h2)
		ltr := center.X-neighbor.Cx() < 0
		ttb := center.Y-neighbor.Cy() < 0

		op := &ebiten.DrawImageOptions{}

		line1.Fill(color.Gray{})
		line2.Fill(color.Gray{})

		if ltr {
			op.GeoM.Translate(float64(center.X), float64(center.Y))
		} else {
			op.GeoM.Translate(float64(neighbor.Cx()), float64(center.Y))
		}
		screen.DrawImage(line1, op)

		op.GeoM.Reset()
		if ttb {
			op.GeoM.Translate(float64(d.Cx()), float64(center.Y))
		} else {
			op.GeoM.Translate(float64(center.X), float64(neighbor.Cy()))
		}
		screen.DrawImage(line2, op)
	}
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
