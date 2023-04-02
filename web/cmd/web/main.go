package main

import (
	"WEB/pkg/config"
	"WEB/pkg/handlers"
	"WEB/pkg/render"
	"WEB/routers"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	app.IsProduction = false

	ss := scs.New()
	ss.Lifetime = 60 * time.Minute
	//關閉瀏覽器後，Session是否保留
	ss.Cookie.Persist = true
	ss.Cookie.SameSite = http.SameSiteLaxMode
	//是否啟用https(正是區需要啟用)
	ss.Cookie.Secure = app.IsProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routers.Routes(&app),
	}

	err = srv.ListenAndServe()
}
