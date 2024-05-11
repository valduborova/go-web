package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"

  "base64"
  "unixtime"
  "httpmiddleware"
)

func main() {
    router := gin.Default()

    router.Use(httpmiddleware.SetRequestIdHeader())
    router.Use(httpmiddleware.LogRequest())

    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    base64.RegisterRoutes(router)
    unixtime.RegisterRoutes(router)

    srv := http.Server{
        Addr:    fmt.Sprintf(":%d", 3000),
        Handler: router,
    }
    srv.ListenAndServe()
}