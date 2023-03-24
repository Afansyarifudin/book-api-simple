package main 

import "book-api/routers"

func main() {
	var PORT = ":3000"

	routers.StartServer().Run(PORT)
}