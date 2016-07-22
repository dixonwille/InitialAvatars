package avatar

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//Color is the RGBA value of any color.
type Color struct {
	R          int
	G          int
	B          int
	A          float64
	brightness float64
}

const (
	rgbaFormater  = "rgba(%d, %d, %d, %f)"
	maxBrightness = (0.2126 * 255.0) + (0.7152 * 255.0) + (0.0722 * 255.0)
)

//NewColor creates a color with the given Red, Green, Blue, and Alpha values.
//Red, Green, and Blue are integers with values from 0 to 255.
//Alpha is measure from 0.0 (transparent) to 1.0 (opaque).
func NewColor(r, g, b int, a float64) (*Color, error) {
	if r < 0 || r > 255 {
		return nil, newError(errValueOutOfRange, "Red needs to be between 0 and 255")
	}
	if g < 0 || g > 255 {
		return nil, newError(errValueOutOfRange, "Green needs to be between 0 and 255")
	}
	if b < 0 || b > 255 {
		return nil, newError(errValueOutOfRange, "Blue needs to be between 0 and 255")
	}
	if a < 0.0 || a > 1.0 {
		return nil, newError(errValueOutOfRange, "Alpha needs to be between 0.0 and 1.0")
	}

	brightness := (0.2126 * float64(r)) + (0.7152 * float64(g)) + (0.0722 * float64(b))
	return &Color{
		R:          r,
		G:          g,
		B:          b,
		A:          a,
		brightness: brightness,
	}, nil
}

//ColorFromHex creates the color from hex value.
//The Alpha channel will be 1.0.
func ColorFromHex(hexString string) (*Color, error) {
	if len(hexString) != 6 {
		return nil, newError(errInvalidValue, "Hex String Must be 6 Characters long")
	}
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, newError(errInvalidValue, "Hex String can only contain the Characters 0-9 and a-f")
	}
	r := int(hexBytes[0])
	g := int(hexBytes[1])
	b := int(hexBytes[2])
	return NewColor(r, g, b, 1.0)
}

//RandomColor will generate a random color to use.
//The alpha channel will stay at 1.0
func RandomColor() *Color {
	red := rand.Intn(256)
	green := rand.Intn(256)
	blue := rand.Intn(256)
	color, _ := NewColor(red, green, blue, 1.0)
	return color
}

func (c *Color) String() string {
	return fmt.Sprintf(rgbaFormater, c.R, c.G, c.B, c.A)
}
