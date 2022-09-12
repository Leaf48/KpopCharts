package main

import (
	"log"
	"net/http"
	"time"

	"main/internal/pkg"

	"github.com/labstack/echo/v4"
)

func main() {
	r := pkg.Billboard()
	go checkBillboard(&r)

	e := echo.New()

	e.GET("/billboard", func(c echo.Context) error {
		return c.JSON(http.StatusOK, r)
	})
	e.Start(":2000")
}

func checkBillboard(r *pkg.GroupsList) {
	for {
		// refresh every 10 hours
		time.Sleep(time.Duration(10) * time.Hour)
		log.Println("Billboard ranking has been updated!")
		*r = pkg.Billboard()
	}
}
