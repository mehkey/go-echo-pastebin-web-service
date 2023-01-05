package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/mehkey/go-pastebin-web-service/internal/datasource"
)

type mock struct {
	users     []datasource.User
	pastebins []datasource.Pastebin
}

var (
	pastebin = datasource.Pastebin{
		ID:      1,
		UserID:  1,
		Content: "Test Content",
	}
)

/*
func TestHandler_GetAllPastebins(t *testing.T) {
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
*/
/*
func TestHandler_GetPastebinsByID_success(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/pastebins/1", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/pastebins/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, h.GetPastebinsByID(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		var pastebin *datasource.Pastebin
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &pastebin))
		assert.Equal(t, 1, pastebin.ID)
	}
}
*/

func TestHandler_GetPastebinsByID_failure(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/pastebins/5", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/pastebins/:id")
	c.SetParamNames("id")
	c.SetParamValues("5")

	assert.Error(t, h.GetUserByID(c), "should return error")
}

func (m *mock) GetAllPastebins() ([]datasource.Pastebin, error) {
	return m.pastebins, nil
}

func (m *mock) GetAllUsers() ([]datasource.User, error) {
	return nil, nil
}

func (m *mock) GetPastebinByID(id int) (*datasource.Pastebin, error) {
	for _, pastebin := range m.pastebins {
		if pastebin.ID == id {
			return &pastebin, nil
		}
	}
	return nil, errors.New("bad stuff happened")
}

func (m *mock) GetUserByID(id int) (*datasource.User, error) {
	return nil, nil
}

func (m *mock) GetPastebinsForInstructor(id int) ([]datasource.Pastebin, error) {
	return nil, nil
}
func (m *mock) GetPastebinsForUser(id int) ([]datasource.Pastebin, error) {
	return nil, nil
}

func (m *mock) CreateNewUser(*datasource.User) (int, error) {
	return -1, nil
}
func (m *mock) AddUserInterest(id int, interests []string) (int, error) {
	return -1, nil
}
