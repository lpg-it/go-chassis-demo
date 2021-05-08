package main

import (
	"github.com/go-chassis/go-chassis/v2"
	hello "go-chassis-demo/handlers"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/server/
func main() {
	chassis.RegisterSchema("rest", &hello.Hello{})
	if err := chassis.Init(); err != nil {
		return
	}
	chassis.Run()
}
