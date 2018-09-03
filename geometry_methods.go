//pg 155-6 in Go Programming Language Book (Pike, et.al)

package main

import (
	"math"
	"fmt"
)

type Point struct{ X, Y float64 }

// traditional function--not used below at all!
func HypotDistanceTraditional(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) HypotDistance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point //a type slice Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].HypotDistance(path[i]) //calls method, not function
		}
	}
	return sum
}
func main(){
	perim := Path{{1,1},{5,1},{5,4},{1,1}}
	fmt.Println(perim.Distance())//12
}
