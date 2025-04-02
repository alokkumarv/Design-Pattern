package main

import (
	"fmt"
	"time"
)

type TrafficLightState interface {
	HandleRequest(tls *TrafficLightSystem)
}

type RedLightState struct{}

func (rls *RedLightState) HandleRequest(tls *TrafficLightSystem) {
	fmt.Println("🔴 Red Light - Stop! 🚗🚦")
	time.Sleep(2 * time.Second) // Simulating delay
	tls.SetState(&GreenLightState{})

}

type GreenLightState struct{}

func (gls *GreenLightState) HandleRequest(tls *TrafficLightSystem) {
	fmt.Println("🟢 Green Light - Go! 🚗💨")
	time.Sleep(4 * time.Second)
	tls.SetState(&YellowLightState{})
}

type YellowLightState struct{}

func (yls *YellowLightState) HandleRequest(tls *TrafficLightSystem) {
	fmt.Println("🟡 Yellow Light - Slow down! ⚠️")
	time.Sleep(1 * time.Second)
	tls.SetState(&RedLightState{})
}

// context
type TrafficLightSystem struct {
	state TrafficLightState
}

func (tls *TrafficLightSystem) SetState(state TrafficLightState) {
	tls.state = state
}

func (tls *TrafficLightSystem) ChangeState() {
	tls.state.HandleRequest(tls)
}

func main() {
	tls := TrafficLightSystem{}
	tls.SetState(&RedLightState{})
	for i := 0; i < 100; i++ { // Run for a few cycles
		tls.ChangeState()
	}

}
