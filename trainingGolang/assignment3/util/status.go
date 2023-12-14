package util

import (
	"assignment3/model"
	"fmt"
)

func GetStatus(wotah []model.WaterStatus) {
	var statusWater string = ""
	var statusWind string = ""
	if len(wotah) > 0 {
		lastStatus := wotah[len(wotah)-1]
		switch {
		case lastStatus.Water < 5:
			statusWater = "Aman"
		case lastStatus.Water >= 6 && lastStatus.Water <= 8:
			statusWater = "Siaga"
		case lastStatus.Water > 8:
			statusWater = "Bahaya"
		}

		switch {
		case lastStatus.Wind < 6:
			statusWind = "Aman"
		case lastStatus.Wind >= 7 && lastStatus.Wind <= 15:
			statusWind = "Siaga"
		case lastStatus.Wind > 15:
			statusWind = "Bahaya"
		}

		fmt.Printf("Water: %#v", lastStatus.Water)
		fmt.Printf(" - Status: %s\n", statusWater)
		fmt.Printf("Wind: %#v", lastStatus.Wind)
		fmt.Printf(" - Status: %s\n", statusWind)
	} else {
		fmt.Println("No data in the JSON file.")
	}
}
