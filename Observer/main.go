package main

import "fmt"

type Observer interface {
	Update(temprature float64)
}

type Display struct {
	name string
}

func (d *Display) Update(temp float64) {
	fmt.Printf("%s Display: Temperature updated to %.2f°C\n", d.name, temp)
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type WeatherStation struct {
	observers  []Observer
	temprature float64
}

func (w *WeatherStation) RegisterObserver(obs Observer) {
	w.observers = append(w.observers, obs)
}

func (w *WeatherStation) RemoveObserver(obs Observer) {
	for i, ob := range w.observers {
		if ob == obs {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}

	}
}

func (w *WeatherStation) NotifyObservers() {
	for _, obs := range w.observers {
		obs.Update(w.temprature)
	}
}

func (w *WeatherStation) SetTemprature(temp float64) {
	w.temprature = temp
	w.NotifyObservers()
}

func main() {

	ws := &WeatherStation{}

	diaplay1 := &Display{name: "Phone"}
	display2 := &Display{name: "Laptop"}

	ws.RegisterObserver(diaplay1)
	ws.RegisterObserver(display2)

	ws.SetTemprature(1.23)
	ws.SetTemprature(1.34)

	ws.RemoveObserver(diaplay1)
	ws.SetTemprature(2.34)

}
