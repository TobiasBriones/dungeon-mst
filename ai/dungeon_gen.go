package ai

import (
	"dungeon-mst/model"
	"math"
	"math/rand"
)

const n = 100000
const maxWidthFactor = 8
const maxHeightFactor = 5

func GenerateDungeons(dimension model.Dimension) []*model.Dungeon {
	var dungeons []*model.Dungeon
	minDim := getMinSize()
	maxDim := getMaxSize()

	for i := 0; i < n; i++ {
		p := getRandomPoint(dimension, maxDim)
		factor := getRandomFactor()
		w := factor.Width * minDim.Width
		h := factor.Height * minDim.Width
		l := p.X - w/2
		t := p.Y - h/2
		p0 := model.Point{X: l, Y: t}
		rect := &model.Rect{
			Left:   l,
			Top:    t,
			Right:  l + w,
			Bottom: t + h,
		}
		shouldContinue := false

		for _, dungeon := range dungeons {
			if dungeon.Intersects(rect) {
				shouldContinue = true
				break
			}
		}
		if shouldContinue {
			continue
		}
		dungeon := model.NewDungeon(p0, factor)
		dungeons = append(dungeons, &dungeon)
	}
	return dungeons
}

func getMinSize() model.Dimension {
	baseSize := model.GetDungeonHorizontalUnitSize().Width
	return model.Dimension{
		Width:  baseSize,
		Height: baseSize,
	}
}

func getMaxSize() model.Dimension {
	baseSize := model.GetDungeonHorizontalUnitSize().Width
	return model.Dimension{
		Width:  maxWidthFactor * baseSize,
		Height: maxHeightFactor * baseSize,
	}
}

func getRandomPoint(dimension model.Dimension, maxDim model.Dimension) model.Point {
	cx := maxDim.SemiWidth() + int(float64(dimension.Width-maxDim.Width)*rand.Float64())
	cy := maxDim.SemiHeight() + int(float64(dimension.Height-maxDim.Height)*rand.Float64())
	return model.Point{X: cx, Y: cy}
}

func getRandomFactor() model.DimensionFactor {
	wFactor := 1 + int(math.Floor(float64(maxWidthFactor)*rand.Float64()))
	hFactor := 1 + int(math.Floor(float64(maxHeightFactor)*rand.Float64()))
	return model.DimensionFactor{Width: wFactor, Height: hFactor}
}
