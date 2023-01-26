package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/mehkey/go-pastebin-web-service/internal/datasource"
)

func TestHandler_GetAllUsers(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/pastebins", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	if assert.NoError(t, h.GetAllPastebins(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		var pastebins []datasource.Pastebin
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &pastebins))
		assert.Equal(t, pastebin, pastebins[0])
	}
}

func TestHandler_GetUsersByID_success(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/user/1", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/user/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, h.GetPastebinByID(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		var user *datasource.User
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &pastebin))
		assert.Equal(t, 1, user.ID)
	}
}

func TestHandler_GetUsersByID_failure(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/user/5", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/user/:id")
	c.SetParamNames("id")
	c.SetParamValues("5")

	assert.Error(t, h.GetUserByID(c), "should return error")
}

func (m *mock) GetAllUsers() ([]datasource.User, error) {

	return []datasource.User{user}, nil
}

func (m *mock) GetUserByID(id int) (*datasource.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (m *mock) CreateNewUser(*datasource.User) (int, error) {
	return 1, nil
}
