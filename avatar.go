package avatar

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ajstarks/svgo"
)

//Avatar is used to create the svg
type Avatar struct {
	svg      *svg.SVG
	initials string
	color    Color
	size     int
	title    string
	desc     string
}

const (
	//Since this is an SVG we can keep this constant.
	//User can change the width and height in html if they need to.
	size = 200
	desc = "This Avatar was created using InitialAvatars from https://github.com/dixonwille/InitialAvatars"
)

//NewAvatar populates the Avatar stucture so that you can create the Avatar.
//Does not populate the SVG field until we know what we are writting to.
func NewAvatar(initials string, color Color) (*Avatar, error) {
	if len(initials) > 2 {
		return nil, newError(errInvalidValue, "Initials can only have 2 letters")
	}
	initials = strings.ToUpper(initials)
	var title = "Avatar for " + initials
	return &Avatar{
		initials: initials,
		color:    color,
		size:     size,
		title:    title,
		desc:     desc,
	}, nil
}

//CreateFile writes the Avatar to the file specified.
func (a *Avatar) CreateFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	a.svg = svg.New(f)
	a.makeSVG()
	return nil
}

//Create writes the Avatar svg to the io.writer.
func (a *Avatar) Create(w io.Writer) {
	a.svg = svg.New(w)
	a.makeSVG()
}

func (a *Avatar) makeSVG() {
	a.svg.Startview(a.size, a.size, 0, 0, a.size, a.size)
	a.svg.Title(a.title)
	a.svg.Desc(a.desc)
	a.svg.Group(`class="avatar"`)
	a.svg.Circle(a.size/2, a.size/2, a.size/2, `class="avatarBackground"`, fmt.Sprintf("fill:%s", a.color.String()))
	a.svg.Text(a.size/2, a.size/2, a.initials, `class="avatarInitials"`, fmt.Sprintf(`font-family:Tahoma, Geneva, sans-serif;
		text-anchor:middle;
		alignment-baseline:central;
		font-size:%fpx;
		fill:white;`, float64(a.size)*0.50))
	a.svg.Gend()
	a.svg.End()
}
