package controllers

import (
	"data-siswa/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllSiswa(c echo.Context) error {
	result, err := models.FetchAllSiswa()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return c.JSON(http.StatusOK, result)
}

func StoreSiswa(c echo.Context) error {
	nama := c.FormValue("nama")
	kelas := c.FormValue("kelas")
	jenjang := c.FormValue("jenjang")
	smt := c.FormValue("smt")

	result, err := models.StoreSiswa(nama, kelas, jenjang, smt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateSiswa(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	kelas := c.FormValue("kelas")
	jenjang := c.FormValue("jenjang")
	smt := c.FormValue("smt")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSiswa(conv_id, nama, kelas, jenjang, smt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSiswa(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteSiswa(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
