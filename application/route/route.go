package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

type Position struct {
	Lat   float64
	Longi float64
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id cant be empty")
	}

	f, err := os.Open("destinations/" + r.ID + ".txt")

	defer f.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		longi, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}

		r.Positions = append(r.Positions, Position{Lat: lat, Longi: longi})
	}
	return nil
}

func (r *Route) ExportToJson() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for i, v := range r.Positions {
		route.ClientID = r.ClientID
		route.ID = r.ID
		route.Position = []float64{v.Lat, v.Longi}

		route.Finished = false
		if total-1 == i {
			route.Finished = true
		}

		json, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}

		result = append(result, string(json))
	}

	return result, nil
}
