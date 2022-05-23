package routes

import (
	"data-siswa/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This is Echo Golang Web Server")
	})

	e.GET("/siswa", controllers.FetchAllSiswa)

	e.POST("/siswa", controllers.StoreSiswa)

	e.PUT("/siswa", controllers.UpdateSiswa)

	e.DELETE("/siswa", controllers.DeleteSiswa)

	return e
}
