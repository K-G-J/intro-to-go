package main

import (
	"fmt"
	"reflect"
	"net"
	"time"
)

func main() {
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.5))
	fmt.Println(reflect.TypeOf("hello"))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf(net.IPv4(127, 0, 0, 1)))
	fmt.Println(reflect.TypeOf(time.Now()))
}