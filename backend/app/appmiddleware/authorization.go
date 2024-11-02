package appmiddleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RequestAuthHandker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			log.Printf("Error: Not include Authorization")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized", "message": "Not include Authorization"})
		}

		var secretKey string
                var err error

		jwtConf := os.Getenv("JWT_CONFIG")
		if jwtConf == "" {
			secretKey = os.Getenv("JWT_SECRET_KEY")
		} else {
			secretKey, err = getEnvVariable("JWT_CONFIG", "JWT_SECRET_KEY")
		}

		log.Printf("%v", tokenString)
		log.Printf("%v", secretKey)

		if err != nil {
			log.Printf("Error: %v", err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}

		_, err2 := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			secret := []byte(secretKey)
			return secret, nil
		})

		if err2 != nil {
			log.Printf("Error: %v", err2)
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
		return next(c)
	}
}

func getEnvVariable(jsonEnvKey string, key string) (string, error) {
	// Get the environment variable that contains JSON data
	jsonEnv := os.Getenv(jsonEnvKey)
	if jsonEnv == "" {
		return "", errors.New(fmt.Sprintf("Environment variable %s is not configured", jsonEnvKey))
	}

	// Parse the JSON data
	var jsonData map[string]string
	err := json.Unmarshal([]byte(jsonEnv), &jsonData)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to parse JSON from %s: %s", jsonEnvKey, err))
	}

	// Extract the value for the given key
	if val, exists := jsonData[key]; exists {
		return val, nil
	}
	return "", errors.New(fmt.Sprintf("Key %s not found in JSON from %s", key, jsonEnvKey))
}
