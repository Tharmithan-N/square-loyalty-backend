// package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/tharmi/square-loyalty-backend/controllers"
// 	"github.com/tharmi/square-loyalty-backend/utils"
// )

// func RegisterRoutes(r *gin.Engine) {
// 	api := r.Group("/api")
// 	{
// 		api.POST("/login", utils.FakeLogin)

// 		secured := api.Group("/")
// 		secured.Use(utils.AuthMiddleware())
// 		{
// 			secured.POST("/earn", controllers.EarnPoints)
// 			secured.POST("/redeem", controllers.RedeemPoints)
// 			secured.GET("/balance", controllers.GetBalance)
// 			secured.GET("/history", controllers.GetHistory)
// 		}
// 	}
// }

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tharmi/square-loyalty-backend/controllers"
	"github.com/tharmi/square-loyalty-backend/utils"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Public route - no authentication
	api.POST("/login", utils.FakeLogin)

	// Protected routes
	secured := api.Group("/")
	secured.Use(utils.AuthMiddleware())
	{
		secured.POST("/earn", controllers.EarnPoints)
		secured.POST("/redeem", controllers.RedeemPoints)
		secured.GET("/balance", controllers.GetBalance)
		secured.GET("/history", controllers.GetHistory)
	}
}
