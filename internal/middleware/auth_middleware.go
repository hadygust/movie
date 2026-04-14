package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadygust/movie/internal/env"
	"github.com/hadygust/movie/internal/service"
)

func (m AuthMiddleware) RequireAuth(c *gin.Context) {

	// get cookie
	tokenString, err := c.Cookie("Authentication")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		secret, _ := env.GetSecret()
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": errors.New("token expired"),
			})
		}

		// Get user
		id := claims["id"].(string)
		user, err := m.svc.FindByID(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Attach user to req
		c.Set("user", user)

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errors.New("Claims not found"),
		})
		return
	}

	c.Next()
}

func (m *AuthMiddleware) RequireNonUser(c *gin.Context) {
	_, err := c.Cookie("Authentication")
	// if cookie exists
	if err == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "already logged in",
		})
		return
	}

	c.Next()
}

type AuthMiddleware struct {
	svc service.UserService
}

func NewAuthMiddleware(svc service.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		svc: svc,
	}
}
