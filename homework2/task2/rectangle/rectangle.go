// Package Rectangle implement functionst to manipulate rectangle.
// The Rectangle struct takes X, Y values in float64 format.

package rectangle

import "errors"

type Rectangle struct {
	X float64
	Y float64
}

// SideChecker return false any of recctangle sides is 0, otherwise true.
func (r *Rectangle) SideChecker() error {
	if r.X > 0 && r.Y > 0 {
		return nil
	}
	return errors.New("triangle sides cannot be lesss or equal 0")
}

// Primeter method return Perimeter of the Rectangle are greater then 0.
func (r *Rectangle) Perimeter() (float64, error) {
	err := r.SideChecker()
	if err != nil {
		return 0, err
	}
	p := 2 * (r.X + r.Y)
	return p, nil
}

// Area method return Area of the Rectangle if both sides are greater then 0.
func (r *Rectangle) Area() (float64, error) {
	err := r.SideChecker()
	if err != nil {
		return 0, err
	}
	a := r.X * r.Y
	return a, nil
}
