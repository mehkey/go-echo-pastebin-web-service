package handler

import (
	"bytes"
	"encoding/json"
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
	user = datasource.User{
		ID:        1,
		Email:     "test@test.com",
		Name:      "Joe Boe",
		Pastebins: []datasource.Pastebin{pastebin},
	}
	userAdd = datasource.User{
		ID:        2,
		Email:     "test2@test.com",
		Name:      "Joe Balo",
		Pastebins: nil,
	}
)

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

func TestHandler_GetPastebinsByID_success(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/pastebin/1", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/pastebin/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, h.GetPastebinByID(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		var pastebin *datasource.Pastebin
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &pastebin))
		assert.Equal(t, 1, pastebin.ID)
	}
}

func TestHandler_GetPastebinsByID_failure(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/pastebin/5", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/pastebin/:id")
	c.SetParamNames("id")
	c.SetParamValues("5")

	assert.Error(t, h.GetUserByID(c), "should return error")

}

func TestHandler_AddUserPastebin_Success(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()

	s, _ := json.Marshal(pastebin)
	b := bytes.NewBuffer(s)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/1/pastebin", b)

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/user/1/pastebin")
	//c.SetParamNames("id")
	//c.SetParamValues("1")

	if assert.NoError(t, h.AddUserPastebin(c)) {
		assert.Equal(t, http.StatusCreated, w.Code)
		//var pastebin *datasource.Pastebin
		//assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &pastebin))
		//assert.Equal(t, 1, pastebin.ID)
	}

}

func TestHandler_AddUserPastebin_Failure(t *testing.T) {
	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()

	//s, _ := json.Marshal(pastebin)
	//b := nil //bytes.NewBuffer(s)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/1/pastebin", nil)

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c := e.NewContext(r, w)
	c.SetPath("/api/v1/user/1/pastebin")

	assert.Error(t, h.AddUserPastebin(c), "should return error")

}

func TestHandler_GetPastebinsForUser(t *testing.T) {

	m := &mock{pastebins: []datasource.Pastebin{pastebin}}
	h := NewHandler(m)
	e := echo.New()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/pastebins/user/1", nil)
	w := httptest.NewRecorder()
	c := e.NewContext(r, w)

	c.SetPath("/api/v1/pastebins/user/1")
	c.SetParamNames("userID")
	c.SetParamValues("1")

	if assert.NoError(t, h.GetPastebinsForUser(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		var pastebins []datasource.Pastebin
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &pastebins))
		assert.Equal(t, pastebin, pastebins[0])
	}

}

func (m *mock) GetAllPastebins() ([]datasource.Pastebin, error) {
	return m.pastebins, nil
}

func (m *mock) GetPastebinByID(id int) (*datasource.Pastebin, error) {
	for _, pastebin := range m.pastebins {
		if pastebin.ID == id {
			return &pastebin, nil
		}
	}
	return nil, errors.New("bad stuff happened")
}

func (m *mock) GetPastebinsForUser(id int) ([]datasource.Pastebin, error) {
	if id == 1 {
		return m.pastebins, nil
	}
	return nil, errors.New("bad stuff happened")
}

func (m *mock) AddUserPastebin(id int, pastebin *datasource.Pastebin) (int, error) {
	return -1, nil
}
