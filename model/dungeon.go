/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
)

const (
	PathWidthPx            = 16
	CollisionLeft          = 0
	CollisionTop           = 1
	CollisionRight         = 2
	CollisionBottom        = 3
	horizontalUnitWidthPx  = 64
	horizontalUnitHeightPx = 12
)

var (
	bgImage     = getDungeonBgImage()
	brickImage  = getBrickImage()
	brickYImage = getBrickYImage()
)

type Dungeon struct {
	rect   Rect
	factor DimensionFactor
	paths  []*pathTrace
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

func (d *Dungeon) Collides(rect *Rect) int {
	if !d.rect.Intersects(rect) {
		return -1
	}
	subRect := Rect{
		Left:   d.rect.Left + PathWidthPx,
		Top:    d.rect.Top + PathWidthPx,
		Right:  d.rect.Right - PathWidthPx,
		Bottom: d.rect.Bottom - PathWidthPx,
	}
	collision := -1

	if rect.Left < subRect.Left {
		collision = 0
	} else if rect.Top < subRect.Top {
		collision = 1
	} else if rect.Right > subRect.Right {
		collision = 2
	} else if rect.Bottom > subRect.Bottom {
		collision = 3
	}
	return collision
}

func (d *Dungeon) Intersects(rect *Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Dungeon) AddNeighbor(dungeon *Dungeon) {
	center := d.Center()
	sw := PathWidthPx / 2
	path := &pathTrace{
		p00: Point{center.X - sw, center.Y + sw},
		p01: Point{center.X - sw, dungeon.Cy() + sw},
		p10: Point{center.X - sw, dungeon.Cy() - sw},
		p11: Point{dungeon.Cx() - sw, dungeon.Cy() - sw},
	}
	d.paths = append(d.paths, path)
}

func (d *Dungeon) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	wFactor := d.factor.Width
	hFactor := d.factor.Height
	blockWidth := horizontalUnitHeightPx

	// Draw Background
	rect := image.Rect(0, 0, d.rect.Width(), d.rect.Height())

	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Top))
	screen.DrawImage(bgImage.SubImage(rect).(*ebiten.Image), op)

	// Draw Top
	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Top))
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Bottom
	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Bottom-blockWidth))
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Left
	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left), float64(d.rect.Top))
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(brickYImage, op)
		op.GeoM.Translate(0, horizontalUnitWidthPx)
	}

	// Draw Right
	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Right-blockWidth), float64(d.rect.Top))
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(brickYImage, op)
		op.GeoM.Translate(0, horizontalUnitWidthPx)
	}

	// Draw the Neighborhood
	d.drawNeighborhood(screen)
}

func (d *Dungeon) drawNeighborhood(screen *ebiten.Image) {
	for _, path := range d.paths {
		d.drawPath(path, screen)
	}
}

func (d *Dungeon) drawPath(path *pathTrace, screen *ebiten.Image) {
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
	x0 := p0.X
	y0 := p0.Y
	w := factor.Width * horizontalUnitWidthPx
	h := factor.Height * horizontalUnitWidthPx
	rect := Rect{x0, y0, x0 + w, y0 + h}
	return Dungeon{
		rect,
		factor,
		[]*pathTrace{},
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

func getDungeonBgImage() *ebiten.Image {
	bgImg, _, err := ebitenutil.NewImageFromFile("./assets/dungeon_bg.png")

	if err != nil {
		log.Fatal(err)
	}
	return bgImg
}

func getBrickImage() *ebiten.Image {
	brickImg, _, err := ebitenutil.NewImageFromFile("./assets/brick.png")

	if err != nil {
		log.Fatal(err)
	}
	return brickImg
}

func getBrickYImage() *ebiten.Image {
	brickYImg, _, err := ebitenutil.NewImageFromFile("./assets/brick_y.png")

	if err != nil {
		log.Fatal(err)
	}
	return brickYImg
}

func GetDungeonHorizontalUnitSize() Dimension {
	return Dimension{
		Width:  horizontalUnitWidthPx,
		Height: horizontalUnitHeightPx,
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
