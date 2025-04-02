package main

import "log"

type Mediator interface {
	Send(msg string, user *Component)
}

type MediatorConcreate struct {
	users []*Component
}

func (mc *MediatorConcreate) AddComponent(user *Component) {
	mc.users = append(mc.users, user)
	log.Printf("Join mediator by user : %v", user.Name)
}

func (mc *MediatorConcreate) Send(msg string, user *Component) {
	for _, c := range mc.users {

		c.Receive(msg)
	}

}

type Component struct {
	Name     string
	Mediator Mediator
}

func (c *Component) SendMessage(msg string, pc *Component) {
	c.Mediator.Send(msg, pc)

}
func (c *Component) Receive(msg string) {
	log.Printf("Message received by %v : %v", c.Name, msg)
}

func main() {

	mediator := MediatorConcreate{}
	alis := &Component{Name: "alis", Mediator: &mediator}
	bob := &Component{Name: "bob", Mediator: &mediator}
	marry := &Component{Name: "marry", Mediator: &mediator}

	mediator.AddComponent(alis)
	mediator.AddComponent(bob)
	mediator.AddComponent(marry)

	alis.SendMessage("Hello bob from alice", bob)
	bob.SendMessage("Hello merry from bob", marry)

}
