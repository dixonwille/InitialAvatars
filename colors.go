package avatar

import (
	"encoding/hex"
	"fmt"
)

//Color is the RGBA value of any color.
type Color struct {
	R int
	G int
	B int
	A float64
}

const (
	rgbaFormater = "rgba(%d, %d, %d, %f)"
)

//NewColor creates a color with the given Red, Green, Blue, and Alpha values.
//Red, Green, and Blue are integers with values from 0 to 255.
//Alpha is measure from 0.0 (transparent) to 1.0 (opaque).
func NewColor(r, g, b int, a float64) (*Color, error) {
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

//ColorFromHex creates the color from hex value.
//The Alpha channel will be 1.0.
func ColorFromHex(hexString string) (*Color, error) {
	if len(hexString) > 6 {
		return nil, newError(errValueOutOfRange, "Hex")
	}
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	return &Color{
		R: int(hexBytes[0]),
		G: int(hexBytes[1]),
		B: int(hexBytes[2]),
		A: 1.0,
	}, nil
}

func (c *Color) String() string {
	return fmt.Sprintf(rgbaFormater, c.R, c.G, c.B, c.A)
}
