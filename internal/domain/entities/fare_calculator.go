package entities

import "errors"

var (
	ErrInvalidFare = errors.New("invalid fare")
)

type FareCalculator interface {
	Calculate(segment Segment) float64
}

type normalFareCalculator struct {
	fare float64
}

func newNormalFareCalculator() *normalFareCalculator {
	return &normalFareCalculator{
		fare: 2.1,
	}
}

func (n *normalFareCalculator) Calculate(segment Segment) float64 {
	return segment.Distance * n.fare
}

type overnightFareCalculator struct {
	fare float64
}

func newOvernightFareCalculator() *overnightFareCalculator {
	return &overnightFareCalculator{
		fare: 3.9,
	}
}

func (o *overnightFareCalculator) Calculate(segment Segment) float64 {
	return segment.Distance * o.fare
}

type overnightSundayFareCalculator struct {
	fare float64
}

func newOvernightSundayFareCalculator() *overnightSundayFareCalculator {
	return &overnightSundayFareCalculator{
		fare: 5.0,
	}
}

func (o *overnightSundayFareCalculator) Calculate(segment Segment) float64 {
	return segment.Distance * o.fare
}

type sundayFareCalculator struct {
	fare float64
}

func newSundayFareCalculator() *sundayFareCalculator {
	return &sundayFareCalculator{
		fare: 2.9,
	}
}

func (s *sundayFareCalculator) Calculate(segment Segment) float64 {
	return segment.Distance * s.fare
}

func createFareCalculator(segment Segment) (FareCalculator, error) {
	if segment.IsOvernight() && !segment.IsSunday() {
		return newOvernightFareCalculator(), nil
	}
	if segment.IsOvernight() && segment.IsSunday() {
		return newOvernightSundayFareCalculator(), nil
	}
	if !segment.IsOvernight() && segment.IsSunday() {
		return newSundayFareCalculator(), nil
	}
	if !segment.IsOvernight() && !segment.IsSunday() {
		return newNormalFareCalculator(), nil
	}
	return nil, ErrInvalidFare
}
