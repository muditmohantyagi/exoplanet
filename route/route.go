package route

import (
	"github.com/gin-contrib/cors"
	"planet.com/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Max"}
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("POST", "GET", "PUT", "DELETE", "UPDATE", "OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))
	/**Allow origin CORS setting end:**/

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})
	/*-------------Routeing started---------------*/
	user := r.Group("api/exoplanet")
	{
		var ExoplanetController = new(controller.ExoplanetController)
		user.POST("/add_exoplanet", ExoplanetController.AddExoplanet)
		user.GET("/list_all_exoplanet", ExoplanetController.ListAllExoplanet)
		user.GET("/list_exoplanet_byid/:id", ExoplanetController.ListExoplanetById)
		user.PUT("/update_exoplanet", ExoplanetController.UpdateExoplanet)
		user.DELETE("/delete_exoplanet_byid/:id", ExoplanetController.DeleteExoplanetById)
		user.GET("/fuel_estimation", ExoplanetController.FuelEstimation)

		//

	}

	return r

}
