package main

// Vector2D is a simple X, Y tuple
type Vector2D struct {
	X int
	Y int
}

func (v Vector2D) AddVector(v2 Vector2D) Vector2D {
	return Vector2D{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
