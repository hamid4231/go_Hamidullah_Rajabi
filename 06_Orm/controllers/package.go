package controllers

import (
	"net/http"
	"orm/models"
	"orm/services"

	"github.com/labstack/echo/v4"
)

type PackageController struct {
	service services.PackageServices
}

func InitPackageController() PackageController {
	return PackageController{
		service: services.PackageServices{},
	}
}

func (pc *PackageController) GetAll(c echo.Context) error {
	pack, err := pc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error occured when fetching data",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": pack,
	})
}

func (pc *PackageController) GetByID(c echo.Context) error {
	id := c.Param("id")
	pack, err := pc.service.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error occured when fetching data",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": pack,
	})
}

func (pc *PackageController) Create(c echo.Context) error {
	var packageReq models.PackageRequest
	if err := c.Bind(&packageReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to add package",
		})
	}

	pack, err := pc.service.Create(packageReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error occured when creating package data",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "created successfully",
		"package": pack,
	})
}

func (pc *PackageController) Update(c echo.Context) error {
	id := c.Param("id")
	var packagereq models.PackageRequest
	if err := c.Bind(&packagereq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "update package failed",
		})
	}
	pack, err := pc.service.Update(id, packagereq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "update failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "package updated",
		"package": pack,
	})
}

func (pc *PackageController) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := pc.service.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "delete failed",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "delete successful",
	})
}
