package router

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/savvy-bit/gin-react-postgres/controller"
	"github.com/savvy-bit/gin-react-postgres/middleware"
)

func restrictToRoles(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := ginjwt.ExtractClaims(c)

		if claims["role"] == nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "User not found"})
			return
		}
		userRole := claims["role"]
		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
	}
}

func Route(app *gin.Engine) {
	indexController := new(controller.IndexController)
	userController := new(controller.UserController)
	authMiddleware := middleware.Auth()

	// Public endpoints
	app.POST("/login", authMiddleware.LoginHandler)

	// Admin endpoints
	admin := app.Group("/admin")
	admin.Use(authMiddleware.MiddlewareFunc())
	{
		admin.GET("/users", restrictToRoles([]string{"admin"}), userController.GetAllUsers)
	}

	// Auth endpoints
	auth := app.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/me", userController.GetMe)
	}

	// Api
	api := app.Group("/api")
	api.Use()
	{
		api.GET("/version", indexController.GetVersion)
	}
}
