package mapping

// Rect is a helper type to represent a rectangular area.
type Rect struct {
	X1, Y1, X2, Y2 int
}

// NewRect returns a new Rect type with its properties calculated.
func NewRect(x int, y int, w int, h int) *Rect {
	return &Rect{
		X1: x,
		Y1: y,
		X2: x + w,
		Y2: y + h,
	}
}

// Center returns the center of the Rect.
func (r *Rect) Center() (x, y int) {
	return (r.X1 + r.X2) / 2, (r.Y1 + r.Y2) / 2
}

// Intersect returns a bool representing if the Rect intersects with another.
func (r *Rect) Intersect(other *Rect) bool {
	if r.X1 <= other.X2 && r.X2 >= other.X1 && r.Y1 <= other.Y2 && r.Y2 >= other.Y1 {
		return true
	}
	return false
}
