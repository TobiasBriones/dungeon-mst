/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

type Path struct {
	p00   Point
	p01   Point
	p10   Point
	p11   Point
	rect1 Rect
	rect2 Rect
}

func (p *Path) inBounds(rect *Rect) bool {
	rect1 := p.rect1
	rect2 := p.rect2
	return rect1.InBounds(rect) || rect2.InBounds(rect)
}
