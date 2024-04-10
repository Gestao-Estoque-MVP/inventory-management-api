package middlewares

import (
	"net/http"
	"strings"

	"github.com/diogoX451/inventory-management-api/internal/service/auth"
	"github.com/gin-gonic/gin"
)

type Unauthorized struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var unauthorized Unauthorized

		unauthorized.Status = "error"
		unauthorized.Code = http.StatusUnauthorized
		unauthorized.Method = ctx.Request.Method
		unauthorized.Message = "Unauthorized"

		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, unauthorized)
			ctx.Abort()
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusUnauthorized, unauthorized)
			ctx.Abort()
			return
		}
		requestToken := splitToken[1]

		tk := auth.NewJWT()
		id, role, err := tk.ValidateToken(requestToken)
		if err != nil {
			unauthorized.Message = err.Error()
			ctx.JSON(http.StatusUnauthorized, unauthorized)
			ctx.Abort()
			return
		}

		ctx.Set("userId", id)
		ctx.Set("role", role)

		ctx.Next()
	}
}
