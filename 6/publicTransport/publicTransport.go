package publictransport

import (
	"fmt"
	"homework_6/passenger"
)

type PublicTransport interface {
	AcceptPassengers(p passenger.Passenger)
	DropOffPassengers(p passenger.Passenger)
	GetName() string
}

type Bus struct {
	Name       string
	Passengers []passenger.Passenger
}

func (b *Bus) AcceptPassengers(p passenger.Passenger) {
	b.Passengers = append(b.Passengers, p)
	fmt.Printf("Пасажир %s зайшов в автобус \n", p.Name)
}

func (b *Bus) DropOffPassengers(p passenger.Passenger) {
	for i, passenger := range b.Passengers {
		if passenger.Id == p.Id {
			b.Passengers = append(b.Passengers[:i], b.Passengers[i+1:]...)
			fmt.Printf("Пасажир %s вийшов з автобуса \n", p.Name)
			return
		}
	}
	fmt.Printf("Такого пасажира не знайдено")
}

func (b Bus) GetName() string {
	return b.Name
}

type Train struct {
	Name       string
	Passengers []passenger.Passenger
}

func (t *Train) AcceptPassengers(p passenger.Passenger) {
	t.Passengers = append(t.Passengers, p)
	fmt.Printf("Пасажир %s зайшов в поїзд \n", p.Name)
}

func (t *Train) DropOffPassengers(p passenger.Passenger) {
	for i, passenger := range t.Passengers {
		if passenger.Id == p.Id {
			t.Passengers = append(t.Passengers[:i], t.Passengers[i+1:]...)
			fmt.Printf("Пасажир %s вийшов з поїзда \n", p.Name)
			return
		}
	}
	fmt.Printf("Такого пасажира не знайдено")
}

func (t Train) GetName() string {
	return t.Name
}

type Airplane struct {
	Name       string
	Passengers []passenger.Passenger
}

func (a *Airplane) AcceptPassengers(p passenger.Passenger) {
	a.Passengers = append(a.Passengers, p)
	fmt.Printf("Пасажир %s зайшов в літак \n", p.Name)
}

func (a *Airplane) DropOffPassengers(p passenger.Passenger) {
	for i, passenger := range a.Passengers {
		if passenger.Id == p.Id {
			a.Passengers = append(a.Passengers[:i], a.Passengers[i+1:]...)
			fmt.Printf("Пасажир %s вийшов з літака \n", p.Name)
			return
		}
	}
	fmt.Printf("Такого пасажира не знайдено")
}

func (a Airplane) GetName() string {
	return a.Name
}
