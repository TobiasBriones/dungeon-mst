/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
	"math/rand"
)

const (
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
	rect    Rect
	barrier Barrier
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
	return NewPoint(d.Cx(), d.Cy())
}

func (d *Dungeon) GetPathFor(dungeon *Dungeon) *Path {
	center := d.Center()
	hp1x := min(center.X(), dungeon.Cx())
	hp2x := max(center.X(), dungeon.Cx())
	hpy := dungeon.Cy()
	hl := Line{
		p1: Point{hp1x, hpy},
		p2: Point{hp2x, hpy},
	}

	vp1x := center.X()
	vp1y := min(center.Y(), dungeon.Cy())
	vp2y := max(center.Y(), dungeon.Cy())
	vl := Line{
		p1: Point{vp1x, vp1y},
		p2: Point{vp1x, vp2y},
	}

	path := NewPath(hl, vl)
	return &path
}

func (d *Dungeon) Intersects(rect *Rect) bool {
	return d.rect.Intersects(rect)
}

func (d *Dungeon) InBounds(rect *Rect) bool {
	return d.rect.InBounds(rect)
}

func (d *Dungeon) CanMoveTowards(movement Movement, rect *Rect) bool {
	if !d.InBounds(rect) {
		return true
	}
	return !d.barrier.WillCollide(movement, rect)
}

func (d *Dungeon) DrawBarrier(screen *ebiten.Image) {
	d.barrier.Draw(screen)
}

func (d *Dungeon) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Draw Background
	rect := image.Rect(0, 0, d.rect.Width()-2*wallWidth, d.rect.Height()-2*wallWidth)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(d.rect.Left()+wallWidth), float64(d.rect.Top()+wallWidth))
	screen.DrawImage(bgImage.SubImage(rect).(*ebiten.Image), op)
}

func (d *Dungeon) RandomPoint(p int) Point {
	x := rand.Intn(d.Width()-wallWidth*2-p) + d.rect.Left() + wallWidth
	y := rand.Intn(d.Height()-wallWidth*2-p) + d.rect.Top() + wallWidth
	return Point{x, y}
}

func NewDungeon(p0 Point, factor DimensionFactor) Dungeon {
	x0 := p0.X()
	y0 := p0.Y()
	w := factor.Width * horizontalUnitWidthPx
	h := factor.Height * horizontalUnitWidthPx
	rect := NewRect(x0, y0, x0+w, y0+h)
	barrier := NewBarrier(rect, factor)
	return Dungeon{
		rect,
		barrier,
	}
}

type DungeonJSON struct {
	*RectJSON
	*BarrierJSON
}

func (d *DungeonJSON) ToDungeon() *Dungeon {
	dungeon := NewDungeon(
		NewPoint(d.RectJSON.Left, d.RectJSON.Top),
		*d.BarrierJSON.Factor,
	)
	return &dungeon
}

func NewDungeonJSON(d *Dungeon) *DungeonJSON {
	rect := NewRectJSON(&d.rect)
	barrier := NewBarrierJSON(&d.barrier)
	return &DungeonJSON{
		RectJSON:    rect,
		BarrierJSON: barrier,
	}
}

type DimensionFactor struct {
	Width  int
	Height int
}

type Wall struct {
	rect  Rect
	image *ebiten.Image
}

type WallJSON struct {
	*RectJSON
}

func (w *WallJSON) ToWall() *Wall {
	wall := &Wall{*w.RectJSON.ToRect(), nil}
	return wall
}

func NewWallJSON(w *Wall) *WallJSON {
	return &WallJSON{NewRectJSON(&w.rect)}
}

type Barrier struct {
	factor     DimensionFactor
	leftWall   Wall
	topWall    Wall
	rightWall  Wall
	bottomWall Wall
}

func (b *Barrier) WillCollide(movement Movement, objRect *Rect) bool {
	return WillCollide(movement, &b.leftWall.rect, objRect) ||
		WillCollide(movement, &b.topWall.rect, objRect) ||
		WillCollide(movement, &b.rightWall.rect, objRect) ||
		WillCollide(movement, &b.bottomWall.rect, objRect)
}

func (b *Barrier) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	wFactor := b.factor.Width
	hFactor := b.factor.Height
	blockWidth := horizontalUnitHeightPx

	// Draw Top
	op.GeoM.Reset()
	op.GeoM.Translate(float64(b.topWall.rect.Left()), float64(b.topWall.rect.Top()))
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Bottom
	op.GeoM.Reset()
	op.GeoM.Translate(float64(b.bottomWall.rect.Left()), float64(b.bottomWall.rect.Bottom()-blockWidth))
	for i := 0; i < wFactor; i++ {
		screen.DrawImage(brickImage, op)
		op.GeoM.Translate(horizontalUnitWidthPx, 0)
	}

	// Draw Left
	op.GeoM.Reset()
	op.GeoM.Translate(float64(b.leftWall.rect.Left()), float64(b.leftWall.rect.Top()))
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(brickYImage, op)
		op.GeoM.Translate(0, horizontalUnitWidthPx)
	}

	// Draw Right
	op.GeoM.Reset()
	op.GeoM.Translate(float64(b.rightWall.rect.Right()-blockWidth), float64(b.rightWall.rect.Top()))
	for i := 0; i < hFactor; i++ {
		screen.DrawImage(brickYImage, op)
		op.GeoM.Translate(0, horizontalUnitWidthPx)
	}
}

func NewBarrier(rect Rect, factor DimensionFactor) Barrier {
	return Barrier{
		factor: factor,
		leftWall: Wall{
			rect: NewRect(
				rect.Left(),
				rect.Top(),
				rect.Left()+wallWidth,
				rect.Bottom(),
			),
			image: brickYImage,
		},
		topWall: Wall{
			rect: NewRect(
				rect.Left(),
				rect.Top(),
				rect.Right(),
				rect.Top()+wallWidth,
			),
			image: brickImage,
		},
		rightWall: Wall{
			rect: NewRect(
				rect.Right()-wallWidth,
				rect.Top(),
				rect.Right(),
				rect.Bottom(),
			),
			image: brickYImage,
		},
		bottomWall: Wall{
			rect: NewRect(
				rect.Left(),
				rect.Bottom()-wallWidth,
				rect.Right(),
				rect.Bottom(),
			),
			image: brickImage,
		},
	}
}

type BarrierJSON struct {
	Factor         *DimensionFactor
	LeftWallJSON   *WallJSON
	TopWallJSON    *WallJSON
	RightWallJSON  *WallJSON
	BottomWallJSON *WallJSON
}

func (b *BarrierJSON) ToBarrier() *Barrier {
	factor := b.Factor
	rect := NewRect(
		b.LeftWallJSON.RectJSON.Left,
		b.TopWallJSON.RectJSON.Top,
		b.RightWallJSON.RectJSON.Right,
		b.BottomWallJSON.RectJSON.Bottom,
	)
	barrier := NewBarrier(rect, *factor)
	return &barrier
}

func NewBarrierJSON(b *Barrier) *BarrierJSON {
	return &BarrierJSON{
		Factor:         &b.factor,
		LeftWallJSON:   NewWallJSON(&b.leftWall),
		TopWallJSON:    NewWallJSON(&b.topWall),
		RightWallJSON:  NewWallJSON(&b.rightWall),
		BottomWallJSON: NewWallJSON(&b.bottomWall),
	}
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
	return NewDimension(
		horizontalUnitWidthPx,
		horizontalUnitHeightPx,
	)
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
