package main

import "fmt"

type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
}

func showFields(m Monitor) {
	fmt.Println("Resolution: ", m.Resolution, "Connector: ", m.Connector, "Value: ", m.Value)
}

// Encapsulation - only methods within package can access the value
type Minutes struct {
	value int
}

// data can only be accessed via getter and setter
func (m *Minutes) Set(newValue int) {
	// do not have to use * because accessing struct field through pointer
	m.value = newValue % 60
}

func (m *Minutes) Increment() {
	m.value = (m.value + 1) % 60
}

func (m *Minutes) Display() {
	fmt.Println(m.value)
}

func main() {
	monitor := Monitor{}
	monitor.Resolution = "1080p"
	monitor.Connector = "HDMI"
	monitor.Value = 249.99
	fmt.Println(monitor.Resolution, monitor.Connector, monitor.Value)

	monitor2 := Monitor{"1080p", "HDMI", 249.99} // must be in order
	showFields(monitor2)

	monitor3 := Monitor{Value: 249.99, Resolution: "1080p"} // if naming fields, can be in any order
	showFields(monitor3)

	monitor4 := Monitor{}
	showFields(monitor4) // will all be default 0 values

	minutes := Minutes{}
	minutes.Set(63)
	for i := 1; i <= 3; i++ {
		minutes.Increment()
		minutes.Display()
	}
}

type Clock struct {
	Hours   int
	Minutes int
}

func Noon() Clock {
	return Clock{Hours: 12}
}
