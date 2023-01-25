package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllPastebins(c echo.Context) error {
	courses, err := h.DB.GetAllPastebins()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}
	return c.JSON(http.StatusOK, courses)
}

func (h *Handler) GetPastebinByID(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}

	course, err := h.DB.GetPastebinByID(id)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}

	if course == nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("course with id : %d not found", id))
	}

	return c.JSON(http.StatusOK, course)
}

func (h *Handler) GetPastebinsForUser(c echo.Context) error {
	id := -1
	if err := echo.PathParamsBinder(c).Int("userID", &id).BindError(); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}
	pastebins, err := h.DB.GetPastebinsForUser(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching data")
	}
	return c.JSON(http.StatusOK, pastebins)
}
