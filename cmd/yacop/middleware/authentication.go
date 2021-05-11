package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var userContextKey = &contextKey{"user_context"}

type contextKey struct {
	name string
}

// Strips 'BEARER ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 7 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

type UserContext struct {
	ID string
}

type CustomClaims struct {
	jwt.StandardClaims
}

// A helper to write user_id and user_model to the context
func updateUserContext(c *gin.Context, userId string) {
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), userContextKey, UserContext{
		ID: userId,
	}))
}

func GetUserContext(c *gin.Context) *UserContext {
	userContext, _ := c.Request.Context().Value(userContextKey).(UserContext)
	return &userContext
}

func AuthenticationMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		updateUserContext(c, "")

		token, err := request.ParseFromRequest(
			c.Request,
			AuthorizationHeaderExtractor,
			config.Config.Keycloak.JWKS.KeySet.KeyFunc,
			request.WithClaims(&CustomClaims{}),
		)
		if err != nil {
			if auto401 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			}
			log.Warn(err)
			return
		}

		claims := token.Claims.(*CustomClaims)
		updateUserContext(c, claims.Subject)
	}
}
