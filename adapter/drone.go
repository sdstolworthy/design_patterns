package adapter

import "fmt"

type Drone struct {
}

func (d *Drone) Beep() {
	fmt.Println("beep bop boop")
}

func (d *Drone) TakeOff() {
	fmt.Println("taking off")
}

func (d *Drone) SpinRotors() {
	fmt.Println("Spinning Rotors")
}

type DroneAdapter struct {
	Drone *Drone
}

func (da *DroneAdapter) Quack() {
	da.Drone.Beep()
}

func (da *DroneAdapter) Fly() {
	da.Drone.SpinRotors()
	da.Drone.TakeOff()
}
func RunAdapterDrone() {
	droneAdapter := DroneAdapter{Drone: new(Drone)}
	droneAdapter.Quack()
	droneAdapter.Fly()
}
