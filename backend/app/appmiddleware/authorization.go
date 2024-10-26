package appmiddleware

import (
	"log"
	"net/http"
	"os"
        "fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RequestAuthHandker(next echo.HandlerFunc) echo.HandlerFunc {
        return func (c echo.Context) error {
                tokenString, err := c.Cookie("Authorization")
                if err != nil {
                        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
                }

                _, err2 := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
                        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
                        }
                        
                        secret := []byte(os.Getenv("JWT_SECRET_KEY"))
                        return secret, nil
                })

                if err2 != nil {
                        log.Fatal(err)
                        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
                }
                return next(c)
        }
}
