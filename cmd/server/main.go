package main

import "fmt"

// App - the struct which contanis things like points
// to database connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up our App")
	return nil
}
func main() {
	fmt.Println("Go Rest API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
