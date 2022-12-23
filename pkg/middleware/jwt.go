package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Claims jwt.MapClaims
}

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		// auth token have the structure `bearer <token>`
		// so we split it on the space
		splitToken := strings.Split(authorization, " ")

		// if we don't have 2 elements, there is an error
		if len(splitToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "no valid token found")
		}

		// parse the JWT token
		jwtToken := splitToken[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("very-secret"), nil
		})

		if err != nil {
			// we got something different
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			cc := &CustomContext{c, claims}
			return next(cc)

		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}
	}
}
