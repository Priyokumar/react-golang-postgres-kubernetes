package main

import (
	"api/db"
	"api/routes"
	"fmt"
)

func main() {
	db.ConnectDb()
	const port int = 8000
	fmt.Println("Hello world.")
	r := routes.GetRouter()
	r.Run(fmt.Sprint(":", port))
}
