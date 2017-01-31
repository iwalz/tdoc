package table

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func placepic(svg *svg.SVG, f io.Reader, x int, y int, w int, h int) error {
	//svg.Group(``, "")

	if Wireframe {
		// Renders the embedded wireframe
		svg.Rect(x, y, w, h, wireoptions)
	}

	b := new(bytes.Buffer)
	b.ReadFrom(f)
	d := b.String()
	d = strings.Replace(d, `<?xml version="1.0" encoding="utf-8"?>`, "", 1)
	r, _ := regexp.Compile("(?s)<svg.*?>")
	// @TODO: Very aws specific, but keep it for development
	d = r.ReplaceAllString(d, fmt.Sprintf(`<svg width="%dpx" height="%dpx" x="%d" y="%d" enable-background="new 0 0 100 100" viewBox="0 0 100 100">`, w, h, x, y))
	io.WriteString(svg.Writer, d)
	//svg.Gend()
	return nil
}

func text(svg *svg.SVG, x int, y int, caption string) error {
	svg.Text(x, y, caption, "font-size: 22")
	return nil
}
