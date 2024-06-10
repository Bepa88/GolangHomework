package route

import (
	"fmt"
	publictransport "homework_6/publicTransport"
)

type Route struct {
	Transports []publictransport.PublicTransport
}

func (r *Route) AddTransport(transport publictransport.PublicTransport) {
	r.Transports = append(r.Transports, transport)
}

func (r Route) GetTransportOnTheRoute() {
	fmt.Println("Маршрут пасажира:")
	for _, v := range r.Transports {
		fmt.Println(v.GetName())
	}
}
