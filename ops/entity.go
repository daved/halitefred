package ops

import "github.com/daved/halitego/geom"

// Entity represents common attributes shared by items in a game map.
type Entity struct {
	id     int
	owner  int
	health float64

	geom.Location
}

// ID ...
func (e Entity) ID() int {
	return e.id
}

// Owner ...
func (e Entity) Owner() int {
	return e.owner
}

// Health ...
func (e Entity) Health() float64 {
	return e.health
}

// Diameter returns the current diameter.
func (e Entity) Diameter() float64 {
	return e.Radius() * 2
}
