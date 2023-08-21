package handlers

import (
	"net/http"
	"strconv"
	dto "test/dto/result"
	siswadto "test/dto/siswa"
	"test/models"
	"test/repositories"

	"github.com/labstack/echo/v4"
)

type siswaHandler struct {
	SiswaRepo repositories.SiswaRepo
}

func HandlerSiswa(SiswaRepo repositories.SiswaRepo) *siswaHandler {
	return &siswaHandler{SiswaRepo}
}

func (h *siswaHandler) FindSiswa(c echo.Context) error {
	users, err := h.SiswaRepo.FindSiswa()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: users})
}

func (h *siswaHandler) CreateSiswa(c echo.Context) error {
	request := models.Siswa{
		Name:     c.FormValue("name"),
		Nis: c.FormValue("nis"),
		Alamat: c.FormValue("alamat"),
	}

	data, err := h.SiswaRepo.CreateSiswa(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest,
			Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseUser(data),
	})
}

func (h *siswaHandler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	request := models.Siswa{
		Nis: c.FormValue("nis"),
		Name: c.FormValue("name"),
		Alamat: c.FormValue("alamat"),
	}
	user, err := h.SiswaRepo.GetSiswa(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest,
			Message: err.Error()})
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Alamat != "" {
		user.Alamat = request.Alamat
	} 
	if request.Nis != "" {
		user.Nis = request.Nis
	}
	
	data, err := h.SiswaRepo.UpdateSiswa(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest,
			Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseUser(data),
	})

}

func (h *siswaHandler) GetSiswa(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.SiswaRepo.GetSiswa(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest,
			Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseUser(data),
	})
}

func (h *siswaHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.SiswaRepo.GetSiswa(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest,
			Message: err.Error()})
	}
	data, err := h.SiswaRepo.DeleteSiswa(user, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest,
			Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseUser(data),
	})
}

func convertResponseUser(user models.Siswa) siswadto.SiswaResponse {
	return siswadto.SiswaResponse{
		Id: user.ID,
		Nis: user.Nis,
		Name: user.Name,
		Alamat: user.Alamat,
	}
}
