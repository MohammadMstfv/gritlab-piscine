package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
)

const (
	W = 500
	H = 120
)

/* ---------------- COLOR ---------------- */

func rgb(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

var reset = "\x1b[0m"

/* ---------------- CANVAS ---------------- */

type Canvas struct {
	w, h int
	col  [][]string
	pix  [][]bool
}

func newCanvas(w, h int) *Canvas {
	c := &Canvas{
		w:   w,
		h:   h,
		col: make([][]string, h),
		pix: make([][]bool, h),
	}

	for i := range c.col {
		c.col[i] = make([]string, w)
		c.pix[i] = make([]bool, w)
	}

	return c
}

func (c *Canvas) set(x, y int, color string, on bool) {
	if x < 0 || y < 0 || x >= c.w || y >= c.h {
		return
	}

	c.pix[y][x] = on
	if on {
		c.col[y][x] = color
	}
}

/* ---------------- BRAILLE ENCODER ---------------- */

func (c *Canvas) render() string {
	out := ""

	for y := 0; y < c.h; y += 4 {
		for x := 0; x < c.w; x += 2 {

			var b rune = 0
			var color string

			// helper
			get := func(dx, dy int) bool {
				xx := x + dx
				yy := y + dy
				if xx >= 0 && yy >= 0 && xx < c.w && yy < c.h {
					if c.pix[yy][xx] {
						color = c.col[yy][xx]
						return true
					}
				}
				return false
			}

			if get(0, 0) {
				b |= 1
			}
			if get(0, 1) {
				b |= 2
			}
			if get(0, 2) {
				b |= 4
			}
			if get(1, 0) {
				b |= 8
			}
			if get(1, 1) {
				b |= 16
			}
			if get(1, 2) {
				b |= 32
			}
			if get(0, 3) {
				b |= 64
			}
			if get(1, 3) {
				b |= 128
			}

			if b == 0 {
				out += " "
			} else {
				out += color + string(rune(0x2800+b)) + reset
			}
		}
		out += "\n"
	}

	return out
}

/* ---------------- FLAGS ---------------- */

func wave(t, u, v, phase float64) float64 {
	return math.Sin(t*2+u*10+phase)*0.02 +
		math.Cos(t*1.5+v*6+phase)*0.015
}

/* ---------------- GEORGIA ---------------- */
func georgia(c *Canvas, ox, oy, w, h int, t float64) {
	red := rgb(200, 30, 40)
	white := rgb(245, 245, 245)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			u := float64(x) / float64(w)
			v := float64(y) / float64(h)

			v += wave(t, u, v, 0)

			col := white

			// helper: thick line
			line := func(a, b, thickness float64) bool {
				return math.Abs(a-b) < thickness
			}

			// --- central cross (proper proportions) ---
			if line(u, 0.5, 0.045) || line(v, 0.5, 0.070) {
				col = red
			}

			// --- bolnur-katskhuri crosses (true cross shape) ---
			drawSmallCross := func(cx, cy float64) bool {
				h := line(u, cx, 0.018) && math.Abs(v-cy) < 0.060
				vv := line(v, cy, 0.018) && math.Abs(u-cx) < 0.060
				return h || vv
			}

			if drawSmallCross(0.25, 0.25) ||
				drawSmallCross(0.75, 0.25) ||
				drawSmallCross(0.25, 0.75) ||
				drawSmallCross(0.75, 0.75) {
				col = red
			}

			c.set(ox+x, oy+y, col, true)
		}
	}
}

/* ---------------- IRAN (SHIR-O-KHORSID) ---------------- */
func iran(c *Canvas, ox, oy, w, h int, t float64) {
	green := rgb(0, 120, 60)
	white := rgb(245, 245, 245)
	red := rgb(200, 30, 40)
	gold := rgb(240, 200, 40)
	brown := rgb(110, 70, 30)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			u := float64(x) / float64(w)
			v := float64(y) / float64(h)

			v += wave(t, u, v, 2.5)

			// background tricolor
			col := white
			if v < 0.33 {
				col = green
			} else if v > 0.66 {
				col = red
			}

			// ---------------- SUN (circle + rays) ----------------
			dx := u - 0.5
			dy := v - 0.42
			r := math.Sqrt(dx*dx + dy*dy)
			angle := math.Atan2(dy, dx)

			// core
			if r < 0.11 {
				col = gold
			}

			// rays (spiky ring)
			if r > 0.11 && r < 0.18 {
				if int(angle*16)%2 == 0 {
					col = gold
				}
			}

			// ---------------- LION (structured silhouette) ----------------

			// body (rounded rectangle via distance approximation)
			body := u > 0.40 && u < 0.60 &&
				v > 0.62 && v < 0.82

			// head
			head := u > 0.47 && u < 0.55 &&
				v > 0.55 && v < 0.63

			// neck
			neck := u > 0.46 && u < 0.54 &&
				v > 0.60 && v < 0.68

			// tail hint
			tail := u > 0.60 && u < 0.65 &&
				v > 0.68 && v < 0.72

			if body || head || neck || tail {
				col = brown
			}

			c.set(ox+x, oy+y, col, true)
		}
	}
}

/* ---------------- FINLAND ---------------- */

func finland(c *Canvas, ox, oy, w, h int, t float64) {
	blue := rgb(0, 70, 160)
	white := rgb(245, 245, 245)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			u := float64(x) / float64(w)
			v := float64(y) / float64(h)

			v += wave(t, u, v, 5.0)

			col := white

			if math.Abs(u-0.35) < 0.08 || math.Abs(v-0.5) < 0.12 {
				col = blue
			}

			c.set(ox+x, oy+y, col, true)
		}
	}
}

/* ---------------- MAIN ---------------- */

func main() {
	panelW := W / 3

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		canvas := newCanvas(W, H)
		t := float64(time.Now().UnixNano()) / 1e9

		georgia(canvas, 0, 0, panelW, H, t)
		iran(canvas, panelW, 0, panelW, H, t)
		finland(canvas, 2*panelW, 0, panelW, H, t)

		fmt.Print(canvas.render())

		time.Sleep(33 * time.Millisecond)
	}
}
