package model

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

const (
	horizontalUnitWidthPx  = 128
	horizontalUnitHeightPx = 24
	verticalUnitWidthPx    = horizontalUnitHeightPx
	verticalUnitHeightPx   = horizontalUnitWidthPx
)

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

func GetDungeonHorizontalUnitSize() Size {
	return Size{
		Width:  horizontalUnitWidthPx,
		Height: horizontalUnitHeightPx,
	}
}

func GetDungeonVerticalUnitSize() Size {
	return Size{
		Width:  verticalUnitWidthPx,
		Height: verticalUnitHeightPx,
	}
}
