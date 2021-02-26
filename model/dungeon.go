/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

const (
	CollisionLeft          = 0
	CollisionTop           = 1
	CollisionRight         = 2
	CollisionBottom        = 3
	horizontalUnitWidthPx  = 64
	horizontalUnitHeightPx = 12
	wallWidth              = horizontalUnitHeightPx
)

var (
	bgImage     = getDungeonBgImage()
	brickImage  = getBrickImage()
	brickYImage = getBrickYImage()
)

type Dungeon struct {
	rect   Rect
	factor DimensionFactor
	paths  []*Path
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
		Left:   d.rect.Left + wallWidth,
		Top:    d.rect.Top + wallWidth,
		Right:  d.rect.Right - wallWidth,
		Bottom: d.rect.Bottom - wallWidth,
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

	// If there's a collision check whether it is on a path
	if collision != -1 {
		for _, path := range d.paths {
			if path.inBounds(rect) {
				// If it has collision on a wall but it is placed inside a path
				// then let it go
				collision = -1
				break
			}
		}
	}
	return collision
}

func (d *Dungeon) Intersects(rect *Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Dungeon) AddNeighbor(dungeon *Dungeon) {
	path := d.getPathFor(dungeon)
	d.paths = append(d.paths, path)
	dungeon.paths = append(dungeon.paths, path)
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
		path.Draw(screen)
	}
}

func (d *Dungeon) getPathFor(dungeon *Dungeon) *Path {
	center := d.Center()
	hp1x := min(center.X, dungeon.Cx())
	hp2x := max(center.X, dungeon.Cx())
	hpy := dungeon.Cy()
	hl := Line{
		p1: Point{hp1x, hpy},
		p2: Point{hp2x, hpy},
	}

	vp1x := center.X
	vp1y := min(center.Y, dungeon.Cy())
	vp2y := max(center.Y, dungeon.Cy())
	vl := Line{
		p1: Point{vp1x, vp1y},
		p2: Point{vp1x, vp2y},
	}

	path := NewPath(hl, vl)
	return &path
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
		[]*Path{},
	}
}

type DimensionFactor struct {
	Width  int
	Height int
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
