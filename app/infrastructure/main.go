package main

func main() {
	db := NewDB()
	r := NewRouting(db)
	r.Run()
}
