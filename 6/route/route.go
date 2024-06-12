package route

import (
	"fmt"
	"homework_6/passenger"
)

type PublicTransport interface {
	AcceptPassengers(p passenger.Passenger)
	DropOffPassengers(p passenger.Passenger)
	GetName() string
}

type Route struct {
	Transports []PublicTransport
}

func (r *Route) AddTransport(transport PublicTransport) {
	r.Transports = append(r.Transports, transport)
}

func (r Route) GetTransportOnTheRoute() {
	fmt.Println("Маршрут пасажира:")
	for _, v := range r.Transports {
		fmt.Println(v.GetName())
	}
}
