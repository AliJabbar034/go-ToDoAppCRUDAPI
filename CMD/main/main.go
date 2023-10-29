package main

import (
	"fmt"

	"github.com/alijabbar034/pkg/model"
	"github.com/alijabbar034/pkg/routes"
)

func main() {
	fmt.Println("Server is listening on port 1000")
	model.Init()
	routes.InitializeRoute()
}
