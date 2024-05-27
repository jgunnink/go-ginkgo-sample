package main

import "gobdd/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
