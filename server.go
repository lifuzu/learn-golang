package server

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "github.com/lifuzu/learn-golang/handlers"
)

func main() {

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.File("/", "public/index.html")

    // handle websocket from frontend
    e.GET("/ws", handlers.Hello)

    e.GET("/hello", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
