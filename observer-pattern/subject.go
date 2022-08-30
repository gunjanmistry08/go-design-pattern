package main

type Subject interface {
	register(ob Observer)
	deregister(ob Observer)
	notifyAll()
}
