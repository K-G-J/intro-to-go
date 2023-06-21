package main

import "fmt"

type Monitor struct {
	Resolution string
	Connector  string
	Value      float64
}

func main() {
	monitor := Monitor{}
	monitor.Resolution = "1080p"
	monitor.Connector = "HDMI"
	monitor.Value = 249.99
	fmt.Println(monitor.Resolution, monitor.Connector, monitor.Value)
}
