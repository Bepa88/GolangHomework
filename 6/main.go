package main

import (
	"homework_6/passenger"
	publictransport "homework_6/publicTransport"
	"homework_6/route"
)

func main() {
	r := &route.Route{}

	bus := &publictransport.Bus{Name: "Автобус №50"}
	train := &publictransport.Train{Name: "Поїзд Київ-Варшава"}
	airplane := &publictransport.Airplane{Name: "Літак Варшава - Барселона"}

	p := passenger.Passenger{Id: 1, Name: "Оля"}

	bus.AcceptPassengers(p)
	train.AcceptPassengers(p)
	airplane.AcceptPassengers(p)

	r.AddTransport(bus)
	r.AddTransport(train)
	r.AddTransport(airplane)

	r.GetTransportOnTheRoute()
}
