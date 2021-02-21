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

func (r *Rect) Width() int {
	return r.Right - r.Left
}

func (r *Rect) SemiWidth() int {
	return r.Width() / 2
}

func (r *Rect) Height() int {
	return r.Bottom - r.Top
}

func (r *Rect) SemiHeight() int {
	return r.Height() / 2
}

func (r *Rect) Intersects(rect *Rect) bool {
	xo := r.Left <= rect.Right &&
		r.Right >= rect.Left

	yo := r.Top <= rect.Bottom &&
		r.Bottom >= rect.Top

	return xo && yo
}
