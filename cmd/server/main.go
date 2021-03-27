package main

import (
	"fmt"
	"github.com/TutorialEdge/go-rest-api-course/database"
	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	"net/http"
	transportHTTP "github.com/TutorialEdge/go-rest-api-course/internal/transport/http"

)



// App - the struct which contanis things like points
// to database connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up our App")
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	commentService := comment.NewService(db)
	handler := transportHTTP.NewHandler(commentService)
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
