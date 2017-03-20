package image

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	svg "github.com/ajstarks/svgo"
)

type Rewriter interface {
	SetX(int)
	X() int
	SetY(int)
	Y() int
	SetWidth(int)
	Width() int
	SetHeight(int)
	Height() int
	SetName(string)
	Name() string
	Place(*svg.SVG, io.Reader) error
}

type Rewrite struct {
	data        string
	name        string
	svg         *regexp.Regexp
	widthRegex  *regexp.Regexp
	heightRegex *regexp.Regexp
	xRegex      *regexp.Regexp
	yRegex      *regexp.Regexp
	classRegex  *regexp.Regexp
	x           int
	y           int
	width       int
	height      int
}

func NewRewrite() *Rewrite {
	svg := regexp.MustCompile("(?si)<\\?xml.*?\\?>.*?(<svg.*?>)")
	widthRegex := regexp.MustCompile("(?si)\\swidth\\s*=\\s*[\"|']?.*?[\"|']?\\s")
	heightRegex := regexp.MustCompile("(?si)\\sheight\\s*=\\s*[\"|']?.*?[\"|']?\\s")
	xRegex := regexp.MustCompile("(?si)\\sx\\s*=\\s*[\"|']?.*?[\"|']?\\s")
	yRegex := regexp.MustCompile("(?si)\\sy\\s*=\\s*[\"|']?.*?[\"|']?\\s")
	classRegex := regexp.MustCompile("(?si)\\sclass\\s*=\\s*[\"|']?(.*?)[\"|']?\\s")

	r := &Rewrite{
		svg:         svg,
		widthRegex:  widthRegex,
		heightRegex: heightRegex,
		xRegex:      xRegex,
		yRegex:      yRegex,
		classRegex:  classRegex,
	}

	return r
}

func (r *Rewrite) Data(i string) {
	r.data = i
}

func (r *Rewrite) rewriteSVGTag() {
	match := r.svg.FindAllStringSubmatch(r.data, 1)
	if match == nil {
		return
	}
	source := match[0][0]
	svg := match[0][1]
	appendix := ""

	// Rewrite width
	if r.widthRegex.MatchString(svg) {
		svg = r.widthRegex.ReplaceAllString(svg, fmt.Sprintf(" width='%dpx' ", r.width))
	} else {
		appendix = appendix + fmt.Sprintf(" width='%dpx' ", r.width)
	}

	// Rewrite height
	if r.heightRegex.MatchString(svg) {
		svg = r.heightRegex.ReplaceAllString(svg, fmt.Sprintf(" height='%dpx' ", r.height))
	} else {
		appendix = appendix + fmt.Sprintf(" height='%dpx' ", r.height)
	}

	// Rewrite x
	if r.xRegex.MatchString(svg) {
		svg = r.xRegex.ReplaceAllString(svg, fmt.Sprintf(" x='%dpx' ", r.x))
	} else {
		appendix = appendix + fmt.Sprintf(" x='%dpx' ", r.x)
	}

	// Rewrite y
	if r.yRegex.MatchString(svg) {
		svg = r.yRegex.ReplaceAllString(svg, fmt.Sprintf(" y='%dpx' ", r.y))
	} else {
		appendix = appendix + fmt.Sprintf(" y='%dpx' ", r.y)
	}

	if appendix != "" {
		svg = strings.Replace(svg, ">", appendix+">", 1)
	}

	r.data = strings.Replace(r.data, source, svg, 1)
}

func (r *Rewrite) rewriteClassesAndIds() {
	match := r.classRegex.FindAllStringSubmatch(r.data, -1)
	classNames := make(map[string]bool, 0)

	for fi, _ := range match {
		for si, v := range match[fi] {
			if si == 1 {
				classNames[v] = true
			}
		}
	}

	if r.name == "" {
		return
	}

	for class, _ := range classNames {
		r.data = strings.Replace(r.data, class, r.name+"_"+class, -1)
	}
}

func (r *Rewrite) SetName(name string) {
	r.name = name
}

func (r *Rewrite) Name() string {
	return r.name
}

func (r *Rewrite) SetX(x int) {
	r.x = x
}

func (r *Rewrite) SetY(y int) {
	r.y = y
}

func (r *Rewrite) X() int {
	return r.x
}

func (r *Rewrite) Y() int {
	return r.y
}

func (r *Rewrite) SetWidth(width int) {
	r.width = width
}

func (r *Rewrite) SetHeight(height int) {
	r.height = height
}

func (r *Rewrite) Width() int {
	return r.width
}

func (r *Rewrite) Height() int {
	return r.height
}

func (r *Rewrite) Rewrite() string {
	r.rewriteSVGTag()
	r.rewriteClassesAndIds()

	return r.data
}

func (r *Rewrite) Place(svg *svg.SVG, reader io.Reader) error {

	b := new(bytes.Buffer)
	b.ReadFrom(reader)

	r.Data(b.String())
	r.data = r.Rewrite()

	io.WriteString(svg.Writer, r.data)

	return nil
}
