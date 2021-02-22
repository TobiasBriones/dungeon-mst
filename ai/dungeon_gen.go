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
	xMap := map[int]bool{}
	yMap := map[int]bool{}

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

		// Check if there's a dungeon aligned to this one already
		for i := 0; i <= model.PathWidthPx; i++ {
			if xMap[rect.Left+i] ||
				xMap[rect.Cx()-model.PathWidthPx/2+i] ||
				xMap[rect.Right-i] ||
				yMap[rect.Top+i] ||
				yMap[rect.Cy()-model.PathWidthPx/2+i] ||
				yMap[rect.Bottom-i] {
				shouldContinue = true
				break
			}
		}
		if shouldContinue {
			continue
		}

		// Update corners
		xMap[rect.Left] = true
		xMap[rect.Cx()] = true
		xMap[rect.Right] = true
		yMap[rect.Top] = true
		yMap[rect.Cy()] = true
		yMap[rect.Bottom] = true

		// Fill wall widths to avoid paths colliding with walls
		for i := 1; i <= model.PathWidthPx; i++ {
			xMap[rect.Left+i] = true
			xMap[rect.Right-i] = true
			xMap[rect.Cx()-model.PathWidthPx/2+i] = true
			yMap[rect.Top+i] = true
			yMap[rect.Bottom-i] = true
			yMap[rect.Cy()-model.PathWidthPx/2+i] = true
		}

		// Add the dungeon
		dungeon := model.NewDungeon(p0, factor)
		dungeons = append(dungeons, &dungeon)
	}
	return dungeons
}

func GetNeighborhoods(dungeons []*model.Dungeon) {
	var tree []*model.Dungeon
	done := map[*model.Dungeon]bool{}

	tree = append(tree, dungeons[0])
	done[dungeons[0]] = true

	for true {
		if len(tree) == len(dungeons) {
			break
		}
		var a *model.Dungeon
		var b *model.Dungeon
		minDistance := 100000

		for _, d1 := range tree {
			p1 := d1.Center()

			for _, d2 := range dungeons {
				if done[d2] {
					continue
				}

				p2 := d2.Center()
				distance := model.Distance(p1, p2)

				if distance < minDistance {
					minDistance = distance
					a = d1
					b = d2
				}
			}
		}

		if a != nil && b != nil {
			tree = append(tree, b)
			done[b] = true

			a.AddNeighbor(b)
		}
	}
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
