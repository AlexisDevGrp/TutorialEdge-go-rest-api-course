package main

import (
	"fmt"
	"github.com/TutorialEdge/go-rest-api-course/database"
	"github.com/TutorialEdge/go-rest-api-course/internal/comment"
	transportHTTP "github.com/TutorialEdge/go-rest-api-course/internal/transport/http"
	log "github.com/sirupsen/logrus"
	"net/http"
)



// App - contains application information
type App struct {
	Name string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName": app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")
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
		log.Error("failed to set up server")
		return err
	}
	return nil
}
func main() {
	fmt.Println("Go Rest API Course")
	app := App{
		Name: "Commenting Service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up our REST API")
		log.Fatal(err)
	}
}
