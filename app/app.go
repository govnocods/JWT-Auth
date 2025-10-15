package app

import (
	"net/http"

	"github.com/govnocods/JWT-Authorization/internal/db"
	"github.com/govnocods/JWT-Authorization/internal/handlers"
	"github.com/govnocods/JWT-Authorization/internal/middlewares"
)

type App struct {
	DB         *db.DataBase
	Router     *http.ServeMux
	Handler    *handlers.Handler
	Middleware *middlewares.Middleware
}

func NewApp(database *db.DataBase) *App {
	app := &App{
		DB:         database,
		Router:     http.NewServeMux(),
		Handler:    handlers.NewHandler(database),
		Middleware: middlewares.NewMiddleware(database),
	}

	app.routes()
	return app
}

func (a *App) routes() {
    a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./web/static/login.html")
    })

    a.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            http.ServeFile(w, r, "./web/static/login.html")
            return
        }

        a.Handler.AuthHandler(w, r) 
    })

    a.Router.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            http.ServeFile(w, r, "./web/static/profile.html")
            return
        }
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    })

    a.Router.Handle("/api/profile", 
        a.Middleware.AuthMiddleware(http.HandlerFunc(a.Handler.ProfileHandler)),
    )
}
