package clockface

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

var (
	clockCentre    = 150.0
	secondHandSize = 90.0
	minuteHandSize = 80.0
	hourHandSize   = 50.0
)

func secondHand(w io.Writer, t time.Time) {
	makeHand(w, secondHandPoint(t), "f00", secondHandSize)
}

func minuteHand(w io.Writer, t time.Time) {
	makeHand(w, minuteHandPoint(t), "000", minuteHandSize)
}

func hourHand(w io.Writer, t time.Time) {
	makeHand(w, hourHandPoint(t), "000", hourHandSize)
}

func makeHand(w io.Writer, p Point, color string, size float64) {
	p = Point{p.X * size, p.Y * size}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentre, p.Y + clockCentre}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#%s;stroke-width:3px;"/>`, p.X, p.Y, color)
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}

func secondHandPoint(t time.Time) Point {
	return anglePoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return anglePoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return anglePoint(hoursInRadians(t))
}

func anglePoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SVGWriter(w io.Writer, t time.Time) {
	_, _ = io.WriteString(w, svgStart)
	_, _ = io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	_, _ = io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
