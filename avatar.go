package avatar

import (
	"io"

	"github.com/ajstarks/svgo"
)

//Avatar is used to create the svg
type Avatar struct {
	SVG      *svg.SVG
	Initials string
	Colors   []Color
}

//NewAvatar populates the Avatar stucture so that you can create the Avatar.
//Does not populate the SVG field until we know what we are writting to.
func NewAvatar(initials string, colors ...Color) *Avatar {
	return &Avatar{
		Initials: initials,
		Colors:   colors,
	}
}

//CreateFile writes the Avatar to the file specified.
func (a *Avatar) CreateFile(filename string) {

}

//Create writes the Avatar svg to the io.writer.
func (a *Avatar) Create(w io.Writer) {
	a.SVG = svg.New(w)
}
