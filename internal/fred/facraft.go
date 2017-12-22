package fred

import (
	"math"

	"github.com/daved/halitego/ops"
)

type faCraft struct {
	ops.Ship
}

func makeFACraft(s ops.Ship) faCraft {
	return faCraft{s}
}

// Navigate demonstrates how the player might negotiate obsticles between
// a ship and its target
func (c faCraft) Navigate(target ops.Entity, f field) ops.Messenger {
	ob := f.ObstaclesBetween(c.Entity, target)

	if !ob {
		return c.NavigateTo(target, f.Board)
	}

	x0 := math.Min(c.X, target.X)
	x2 := math.Max(c.X, target.X)
	y0 := math.Min(c.Y, target.Y)
	y2 := math.Max(c.Y, target.Y)

	dx := (x2 - x0) / 5
	dy := (y2 - y0) / 5
	bestdist := 1000.0
	bestTarget := target

	for x1 := x0; x1 <= x2; x1 += dx {
		for y1 := y0; y1 <= y2; y1 += dy {
			intermediateTarget := ops.Entity{
				X:      x1,
				Y:      y1,
				Radius: 0,
				Health: 0,
				Owner:  0,
				ID:     -1,
			}
			ob1 := f.ObstaclesBetween(c.Entity, intermediateTarget)
			if !ob1 {
				ob2 := f.ObstaclesBetween(intermediateTarget, target)
				if !ob2 {
					totdist := math.Sqrt(math.Pow(x1-x0, 2)+math.Pow(y1-y0, 2)) + math.Sqrt(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2))
					if totdist < bestdist {
						bestdist = totdist
						bestTarget = intermediateTarget

					}
				}
			}
		}
	}

	return c.NavigateTo(bestTarget, f.Board)
}

// messenger demonstrates how the player might direct their ships
// in achieving victory
func (bot *Fred) messenger(f field, c faCraft) ops.Messenger {
	if c.Status != ops.Undocked {
		return c.NoOp()
	}

	ps := f.PlanetsByProximity(c)

	for _, p := range ps {
		if (p.Owned == 0 || p.Owner == bot.id) && p.DockedCt < p.PortCt && p.ID%2 == c.ID%2 {
			if msg, err := c.Dock(p); err == nil {
				return msg
			}

			return c.Navigate(ops.Nearest(3, p, c.Ship), f)
		}
	}

	return c.NoOp()
}
