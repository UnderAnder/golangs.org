package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

// String форматирует координаты DMS
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// location с широтой и долготой в десятичных градусах
type location struct {
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

// String форматирует location с широтой и долготой
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.Lat, l.Long)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

// decimal конвертирует координаты d/m/s в десятичные градусы.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	elysium := location{
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}
	bytes, err := json.MarshalIndent(elysium, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}
