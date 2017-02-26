package main

import "fmt"

type Vertex struct {
	Lat, long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	var m1 = map[string]Vertex{
		"Bell Labs": Vertex{
			40, 30,
		},
		"Google": Vertex{
			1, 50,
		},
	}
	fmt.Println(m1)
}
