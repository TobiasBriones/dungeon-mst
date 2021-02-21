package model

type Point struct {
	X int
	Y int
}

type Size struct {
	Width  int
	Height int
}

type Rect struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

func (r *Rect) Intersects(rect *Rect) bool {
	xo := r.Left <= rect.Right &&
		r.Right >= rect.Left

	yo := r.Top <= rect.Bottom &&
		r.Bottom >= rect.Top

	return xo && yo
}
