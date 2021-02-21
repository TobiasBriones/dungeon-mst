package model

type Point struct {
	X int
	Y int
}

type Dimension struct {
	Width  int
	Height int
}

func (s *Dimension) SemiWidth() int {
	return s.Width / 2
}

func (s *Dimension) SemiHeight() int {
	return s.Height / 2
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
	xi := r.Left <= rect.Right &&
		r.Right >= rect.Left

	yi := r.Top <= rect.Bottom &&
		r.Bottom >= rect.Top

	return xi && yi
}
