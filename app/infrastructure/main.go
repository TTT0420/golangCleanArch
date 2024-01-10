package main

import infrastructure "goAPI/infrastructure/database"

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
