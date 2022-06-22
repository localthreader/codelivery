package main

import (
	"github.com/localthreader/codelivery/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
}
