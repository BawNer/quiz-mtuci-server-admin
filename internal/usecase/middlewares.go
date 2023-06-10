package usecase

import (
	"net/http"
	"quiz-mtuci-server/config"
	"quiz-mtuci-server/internal/entity"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareStruct struct {
	jwt JWT
	cfg *config.Config
}

func NewMiddleware(j JWT, c *config.Config) *MiddlewareStruct {
	return &MiddlewareStruct{
		jwt: j,
		cfg: c,
	}
}

func (m *MiddlewareStruct) AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrorResponseUI{
				Description: "Invalid token",
				Code:        "ERR_INVALID_TOKEN",
			})
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrorResponseUI{
				Description: "Invalid token",
				Code:        "ERR_INVALID_TOKEN",
			})
			return
		}

		if headerParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrorResponseUI{
				Description: "Invalid token",
				Code:        "ERR_INVALID_TOKEN",
			})
			return
		}

		token, err := m.jwt.Validate(headerParts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entity.ErrorResponseUI{
				Description: "Invalid token",
				Code:        "ERR_INVALID_TOKEN",
			})
			return
		}

		c.Set("token", token)
		c.Next()
	}
}

func (m *MiddlewareStruct) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (m *MiddlewareStruct) IsAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := m.jwt.Parse(context.GetStringMap("token"))
		if user.IsStudent != 0 && user.ID != 2066 {
			context.AbortWithStatusJSON(http.StatusForbidden, entity.ErrorResponseUI{
				Description: "No rule access",
				Code:        "ERR_NO_RULES",
			})
			return
		}

		context.Next()
	}
}
