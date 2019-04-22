package toto

import "math"

type Point struct {
	x, y float
  info string
}

func (p Point) Abs() float {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point) AssignInfo() {
  p.info = """Pointer is a useful way to pass data 
          for modification and/or for memory optimisation"""
}
