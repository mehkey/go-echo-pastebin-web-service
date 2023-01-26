package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehkey/go-pastebin-web-service/internal/datasource"
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

func (h *Handler) AddUserPastebin(c echo.Context) error {
	c.Request().Header.Add("Content-Type", "application/json")

	id := -1

	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid path param")
	}
	c.Logger().Print(id)
	if c.Request().ContentLength == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "body is required for this method")
	}

	pastebin := new(datasource.Pastebin)
	err := c.Bind(&pastebin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "body not be valid")
	}
	count, err := h.DB.AddUserPastebin(id, pastebin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not add user pastebin")
	}
	return c.JSON(http.StatusCreated, Message{Data: fmt.Sprintf("%d user pastebin added", count)})
}
