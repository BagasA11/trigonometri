package trigonometri

import (
	"math"
)

type Position struct {
	x float64
	y float64
}

type Sender struct {
	post0 Position
	post1 Position
}

/*
get length from p1 to p3
@params: p0 {x, y} p1 {(x0, y0), (x1, y1)}
@return float64
*/
func getLength(p0 Position, p1 Sender) float64 {
	var x float64 = (p1.post0.x - p0.x) - (p1.post1.x - p1.post0.x)
	var y float64 = (p1.post0.y - p0.y) - (p1.post1.y - p1.post0.y)
	return math.Sqrt(math.Pow(x, 2.0) + math.Pow(y, 2.0))
}

func newPosition(xx float64, yy float64) *Position {
	return &Position{
		x: xx,
		y: yy,
	}
}

func newSender(pp0 Position, pp1 Position) *Sender {
	return &Sender{
		post0: pp0,
		post1: pp1,
	}
}
