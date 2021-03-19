package main

import (
	"fmt"
	"net/http"
	transportHTTP "github.com/TutorialEdge/go-rest-api-course/internal/transport/http"

)



// App - the struct which contanis things like points
// to database connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up our App")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":9090", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
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
