package routes

import (
	"test/handlers"
	"test/pkg/mysql"
	"test/repositories"

	"github.com/labstack/echo/v4"
)

func SiswaRoutes(e *echo.Group) {
	siswa := repositories.RepoSiswa(mysql.DB)
	h := handlers.HandlerSiswa(siswa)

	e.GET("/siswa", h.FindSiswa)
	e.GET("/siswa/:id" , h.GetSiswa)
	e.POST("/siswa", h.CreateSiswa)
	e.PATCH("/siswa" ,h.FindSiswa)
}