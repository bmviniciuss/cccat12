package ports

import "net/http"

type DriverHandlersPort interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type PassengerHandlersPort interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type RideCalculatorHandlersPort interface {
	Calculate(w http.ResponseWriter, r *http.Request)
}
