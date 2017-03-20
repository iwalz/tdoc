package image

import svg "github.com/ajstarks/svgo"

func Text(svg *svg.SVG, x int, y int, caption string) error {
	svg.Text(x, y, caption, "font-size: 22")
	return nil
}
