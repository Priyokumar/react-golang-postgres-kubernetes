package main

import (
	"api/db"
	"api/routes"
	"fmt"
)

func init() {
	db.ConnectDb()
}

func main() {
	const port int = 8001
	fmt.Println("Hello world.")
	r := routes.GetRouter()
	r.Run(fmt.Sprint(":", port))
}
