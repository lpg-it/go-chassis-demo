package main

import (
	"github.com/go-chassis/go-chassis"
	"go-chassis-demo/handlers"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/path/to/conf/folder
func main() {
	chassis.RegisterSchema("rest", &hello.Hello{})

	if err := chassis.Init(); err != nil {
		return
	}
	chassis.Run()
}