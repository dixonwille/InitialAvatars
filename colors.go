package avatar

//Color is the RGBA value of any color.
type Color struct {
	R int
	G int
	B int
	A float64
}

//NewColor creates a color with the given Red, Green, Blue, and Alpha values.
//Red, Green, and Blue are integers with values from 0 to 255.
//Alpha is measure from 0.0 (transparent) to 1.0 (opaque).
func newColor(r, g, b int, a float64) (*Color, error) {
	if r < 0 || r > 255 {
		return nil, newError(errValueOutOfRange, "Red")
	}
	if g < 0 || g > 255 {
		return nil, newError(errValueOutOfRange, "Green")
	}
	if b < 0 || b > 255 {
		return nil, newError(errValueOutOfRange, "Blue")
	}
	if a < 0.0 || a > 1.0 {
		return nil, newError(errValueOutOfRange, "Alpha")
	}
	return &Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}, nil
}

//ChangeAlpha changes the alpha channel to specified value.
//Alpha is measure from 0.0 (transparent) to 1.0 (opaque).
func (c *Color) changeAlpha(a float64) {
	c.A = a
}
