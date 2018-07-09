package handlers

import (
    "fmt"

    "github.com/labstack/echo"
    "golang.org/x/net/websocket"
)

func Hello(c echo.Context) error {
    websocket.Handler(func(ws *websocket.Conn) {
        defer ws.Close()
        for {
                        // Write
            if err := websocket.Message.Send(ws, "Hello, Client!"); err != nil {
                c.Logger().Error(err)
            }

                        // Read
            msg := ""
            if err := websocket.Message.Receive(ws, &msg); err != nil {
                c.Logger().Error(err)
            }
            fmt.Printf("%s\n", msg)
        }
    }).ServeHTTP(c.Response(), c.Request())
    return nil
}