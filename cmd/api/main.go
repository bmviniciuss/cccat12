package main

func main() {
	app := Build()
	app.Listen(":3000")
}
