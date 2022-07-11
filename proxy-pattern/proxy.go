package proxy

import "fmt"

// Proxy Pattern

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("Station: %s bought a ticket, remaining %d \n", name, station.stock)
	} else {
		fmt.Println("Tickets has been sold out.")
	}
}

type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("ProxyStation: %v bought a ticket, remaining %d \n", name, proxy.station.stock)
	} else {
		fmt.Println("Tickets has been sold out.")
	}
}
