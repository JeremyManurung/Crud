package routes

import (
	"net/http"
	"rest/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil menggunakan echo")
	})

	e.GET("/pegawai", controllers.SelectPegawaiC)
	e.POST("/pegawai", controllers.CreatePegawaiC)
	e.PUT("/pegawai", controllers.UpdatePegawaiC)
	return e
}
