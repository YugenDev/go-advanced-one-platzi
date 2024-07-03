package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvaliability() {
	fmt.Printf("The item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) updateValue(v string) {
	fmt.Printf("Sending email... item %s avalible from client %s\n", v, eC.id)
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func main() {
	nvidiaItem := NewItem("RTX 4090ti")
	firstOberserver := &EmailClient{
		id: "1",
	}
	secondObserver := &EmailClient{
		id: "2",
	}
	nvidiaItem.register(firstOberserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvaliability()
}
