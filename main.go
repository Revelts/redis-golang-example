package main

import (
	"fmt"
	"redis-api/Connections"
	"redis-api/Routes"
)

func main() {
	Connections.InitiateConnection()
	err := Routes.HandleRequests()
	if err != nil {
		fmt.Println(err)
	}
}
