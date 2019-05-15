package tournamentserver

import (
	"net/http"

	"github.com/labstack/echo"
)

type login struct {
	displayName string
	username    string
	password    string
}

// API - Api struct
type API struct {
	router *echo.Echo
	db     *DB
}

// New - Default constructor
func (a *API) New(db *DB) *API {
	a.router = echo.New()
	a.db = db
	return a
}

// Start - start an API instance
func (a *API) Start() {
	a.router.GET("/", a.test)
	a.router.POST("/login", a.login)
	a.router.POST("/Register", a.login)
	a.router.POST("/createTournament", a.test)
	a.router.GET("/getTournaments/:id", a.test)
	a.router.POST("/addDeck", a.test)
	a.router.GET("/User/:name", a.test)
	a.router.GET("/joinTournament/:id", a.test)
	defer a.db.Close()
	a.router.Logger.Fatal(a.router.Start(":1337"))
}

func (a *API) test(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}

func (a *API) login(c echo.Context) error {
	requestUser := new(login)
	if err := c.Bind(requestUser); err != nil {
		return err
	}

	if databaseUser, ok := a.db.getUser(requestUser.username); ok {
		if databaseUser.password == requestUser.password {
			return c.String(http.StatusOK, "")
		}
	}

	return c.String(http.StatusOK, "Password wrong")
}
