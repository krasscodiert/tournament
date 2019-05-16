package tournamentserver

import (
	"net/http"

	"github.com/labstack/echo"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type register struct {
	DisplayName string `json:"displayName"`
	Username    string `json:"username"`
	Password    string `json:"password"`
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
	a.router.POST("/register", a.register)
	a.router.POST("/createTournament", a.test)
	a.router.GET("/getTournaments/:id", a.test)
	a.router.POST("/addDeck", a.test)
	a.router.GET("/User/:name", a.test)
	a.router.GET("/joinTournament/:id", a.test)
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

	if databaseUser, ok := a.db.getUser(requestUser.Username); ok {
		if databaseUser.Password == requestUser.Password {
			return c.String(http.StatusOK, "logged in")
		}
	}

	return c.String(http.StatusOK, "Password wrong")
}

func (a *API) register(c echo.Context) error {
	requestUser := new(register)

	if err := c.Bind(requestUser); err != nil {
		return err
	}

	if err := a.db.saveUser(requestUser); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "OK")
}
