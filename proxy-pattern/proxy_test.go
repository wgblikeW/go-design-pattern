package proxy

import "testing"

func TestProxy(t *testing.T) {
	station := &Station{stock: 10}
	proxyStation := &StationProxy{station: station}

	proxyStation.sell("Peter")
	station.sell("Jack")
}
