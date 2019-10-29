// Surface computes an SVG rendering of a 3D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xYRange       = 30.0                // axis ranges (-xYRange..+xYRange)
	xYScale       = width / 2 / xYRange // pixesl per x or y unit
	zScale        = height * 0.4        //pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: withe; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xYRange * (float64(i)/cells - 0.5)
	y := xYRange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xYScale
	sy := height/2 + (x+y)*sin30*xYScale - z*zScale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distrance from (0,0)

	return math.Sin(r) / r
}
