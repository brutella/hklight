package main

import (
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
	"log"
)

func turnLightOn() {
	log.Println("Turn Light On")
}

func turnLightOff() {
	log.Println("Turn Light Off")
}

func main() {
	info := model.Info{
		Name:         "Personal Light Bulb",
		Manufacturer: "Matthias",
	}

	light := accessory.NewLightBulb(info)
	light.OnStateChanged(func(on bool) {
		if on == true {
			turnLightOn()
		} else {
			turnLightOff()
		}
	})

	t, err := hap.NewIPTransport(hap.Config{Pin: "32191123"}, light.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	t.Start()
}
